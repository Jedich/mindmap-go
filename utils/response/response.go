package response

import (
	"encoding/json"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"go.uber.org/zap"
)

// Alias for any slice.
type Messages = []any

// A struct to handle error with custom error handler.
type Error struct {
	Code    int `json:"code"`
	Message any `json:"message"`
}

// Error makes it compatible with the `error` interface.
func (e *Error) Error() string {
	return fmt.Sprint(e.Message)
}

// Resp is used to return standardized responses.
type Resp struct {
	Code     int      `json:"code"`
	Messages Messages `json:"messages,omitempty"`
	Data     any      `json:"data,omitempty"`
}

// Nothing to describe this fucking variable.
var IsProduction bool

// ErrorHandler is a default error handler
var ErrorHandler = func(c *fiber.Ctx, err error) error {
	resp := Resp{
		Code: fiber.StatusInternalServerError,
	}
	// Handle errors
	switch e := err.(type) {
	case validation.Errors:
		resp.Code = fiber.StatusForbidden
		res, _ := json.Marshal(err)
		resp.Data = res
	case *fiber.Error:
		resp.Code = e.Code
		resp.Messages = Messages{e.Message}
	case *Error:
		resp.Code = e.Code
		resp.Messages = Messages{e.Message}
	}

	if !IsProduction {
		zap.Error(err)
	}

	return Response(c, resp)
}

// NewErrors creates multiple new Error messages
func NewErrors(code int, messages ...any) *Error {
	e := &Error{
		Code:    code,
		Message: utils.StatusMessage(code),
	}
	if len(messages) > 0 {
		e.Message = messages
	}
	return e
}

// NewError creates singular new Error message
func NewError(code int, messages ...any) *Error {
	e := &Error{
		Code:    code,
		Message: utils.StatusMessage(code),
	}
	if len(messages) > 0 {
		e.Message = messages[0]
	}
	return e
}

// Response is a function to return beautiful responses.
func Response(c *fiber.Ctx, resp Resp) error {
	// Set status
	if resp.Code == 0 {
		resp.Code = fiber.StatusOK
	}
	c.Status(resp.Code)

	// Return JSON
	return c.JSON(resp)
}
