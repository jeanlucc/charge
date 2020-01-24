package security

type Credentials struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type ConfirmedCredentials struct {
	Username          string `form:"username"`
	Password          string `form:"password"`
	ConfirmedPassword string `form:"password"`
}
