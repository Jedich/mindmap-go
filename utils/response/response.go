package response

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"go.uber.org/zap"
	utl "mindmap-go/utils"
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

type ErrorData struct {
	ErrorType string `json:"type"`
	Data      any    `json:"data"`
}

// Body is used to return standardized responses.
type Body struct {
	Code     int        `json:"code"`
	Messages Messages   `json:"messages,omitempty"`
	Errors   *ErrorData `json:"errors,omitempty"`
	Data     any        `json:"data,omitempty"`
}

// IsProduction description is unnecessary.
var IsProduction bool

// ErrorHandler is a default error handler
var ErrorHandler = func(c *fiber.Ctx, err error) error {
	resp := Body{
		Code: fiber.StatusInternalServerError,
	}
	// Handle errors
	switch e := err.(type) {
	case validation.Errors:
		resp.Code = fiber.StatusForbidden
		resp.Errors = &ErrorData{ErrorType: "validation", Data: e}
	case *utl.DuplicateEntryError:
		resp.Code = fiber.StatusBadRequest
		resp.Errors = &ErrorData{Data: e.Message}
	case *utl.NonExistentEntryError:
		resp.Code = fiber.StatusNotFound
		resp.Errors = &ErrorData{Data: e.Message}
	case *utl.UnauthorizedEntryError:
		resp.Code = fiber.StatusUnauthorized
		resp.Errors = &ErrorData{Data: e.Message}
	case *fiber.Error:
		resp.Code = e.Code
		resp.Errors = &ErrorData{Data: e.Message}
	case *Error:
		resp.Code = e.Code
		resp.Errors = &ErrorData{Data: e.Message}
	}

	if IsProduction {
		return Send(c, resp)
	}

	errText := fmt.Sprintf("%+v", err.Error())
	if resp.Code != fiber.StatusInternalServerError {
		zap.L().Debug(errText)
	} else {
		zap.L().Error(errText)
	}

	return Send(c, resp)
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

// Send is a function to return beautiful responses.
func Send(c *fiber.Ctx, resp Body) error {
	// Set status
	if resp.Code == 0 {
		resp.Code = fiber.StatusOK
	}
	c.Status(resp.Code)

	// Return JSON
	return c.JSON(resp)
}
