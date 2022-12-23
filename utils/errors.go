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

type NonExistentEntryError struct {
	Message string
}

func (m *NonExistentEntryError) Error() string {
	return "entry does not exist"
}
