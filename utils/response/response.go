package response

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	utl "mindmap-go/utils"
)

// Messages is an alias for any slice.
type Messages = []any

// RespBody is used to return standardized responses.
type RespBody struct {
	Code     int            `json:"code"`
	Messages Messages       `json:"messages,omitempty"`
	Errors   *utl.ErrorData `json:"errors,omitempty"`
	Data     any            `json:"data,omitempty"`
}

// Send is a function to return response to Fiber handler.
func (resp *RespBody) Send(c *fiber.Ctx) error {
	// Set default status
	if resp.Code == 0 {
		resp.Code = fiber.StatusOK
	}

	c.Status(resp.Code)

	// Return JSON
	return c.JSON(resp)
}

// IsProduction description is unnecessary.
var IsProduction bool

// ErrorHandler is a default error handler
var ErrorHandler = func(c *fiber.Ctx, err error) error {
	builder := NewResponseBuilder().WithCode(fiber.StatusInternalServerError)
	// Handle errors
	switch e := err.(type) {
	case validation.Errors:
		builder.WithCode(fiber.StatusForbidden).WithErrors(&utl.ErrorData{ErrorType: "validation", Data: e})
	case *utl.DuplicateEntryError:
		builder.WithCode(fiber.StatusBadRequest).WithErrors(&utl.ErrorData{Data: e.Message})
	case *utl.NonExistentEntryError:
		builder.WithCode(fiber.StatusNotFound).WithErrors(&utl.ErrorData{Data: e.Message})
	case *utl.UnauthorizedEntryError:
		builder.WithCode(fiber.StatusUnauthorized).WithErrors(&utl.ErrorData{Data: e.Message})
	case *fiber.Error:
		builder.WithCode(e.Code).WithErrors(&utl.ErrorData{Data: e.Message})
	case *utl.CustomError:
		builder.WithCode(e.Code).WithErrors(&utl.ErrorData{Data: e.Message})
	}
	resp := builder.Build()

	if IsProduction {
		return resp.Send(c)
	}

	errText := fmt.Sprintf("%+v", err.Error())
	if resp.Code != fiber.StatusInternalServerError {
		zap.L().Debug(errText)
	} else {
		zap.L().Error(errText)
	}

	return resp.Send(c)
}
