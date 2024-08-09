package webhooks

import "fmt"

type ValidationError struct {
	Field string
	Value string
	Err   error
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("invalid `%s` = `%s`: %s", e.Field, e.Value, e.Err)
}

func (e ValidationError) Unwrap() error {
	return e.Err
}

func newValidationError(field, value string, err error) ValidationError {
	return ValidationError{
		Field: field,
		Value: value,
		Err:   err,
	}
}

func IsValidationError(err error) bool {
	_, ok := err.(ValidationError)
	return ok
}
