package utils

import "errors"

var (
	DuplicateEntryError = errors.New("cannot insert duplicate entry")
)
