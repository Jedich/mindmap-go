package utils

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
