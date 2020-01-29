package security

import "fmt"

type PasswordMismatchError struct{}

func (e *PasswordMismatchError) Error() string {
	return "The provided password does not match"
}

type ConfirmedPasswordMismatchError struct{}

func (e *ConfirmedPasswordMismatchError) Error() string {
	return "The provided passwords does not match"
}

type InvalidAuthSessionError struct {
	Previous error
}

func (e *InvalidAuthSessionError) Error() string {
	return fmt.Sprint("Could not retrieve user from session: ", e.Previous.Error())
}
