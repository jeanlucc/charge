package security

type Credentials struct {
	Username string `form:"username"`
	Password string `form:"password"`
}
