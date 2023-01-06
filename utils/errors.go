package utils

import "fmt"

type DuplicateEntryError struct {
	Message string
}

func (m *DuplicateEntryError) Error() string {
	return "cannot insert duplicate entry"
}

type UnauthorizedEntryError struct {
	Message string
}

func (m *UnauthorizedEntryError) Error() string {
	return "unauthorized"
}

type NonExistentEntryError struct {
	Message string
}

func (m *NonExistentEntryError) Error() string {
	return "entry does not exist"
}

// CustomError is used to handle error with custom error handler.
type CustomError struct {
	Code    int `json:"code"`
	Message any `json:"message"`
}

func (e *CustomError) Error() string {
	return fmt.Sprint(e.Message)
}

type ErrorData struct {
	ErrorType string `json:"type"`
	Data      any    `json:"data"`
}
