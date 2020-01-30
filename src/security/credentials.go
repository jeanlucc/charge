package security

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ConfirmedCredentials struct {
	Username          string `json:"username"`
	Password          string `json:"password"`
	ConfirmedPassword string `json:"confirmedPassword"`
}
