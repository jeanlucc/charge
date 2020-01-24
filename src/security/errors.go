package security

type PasswordMismatchError struct {}

func (e *PasswordMismatchError) Error() string {
	return "The provided password does not match"
}

type ConfirmedPasswordMismatchError struct {}

func (e *ConfirmedPasswordMismatchError) Error() string {
	return "The provided passwords does not match"
}
