package entities

type User struct {
	Id       int    `json:"Id"`
	Username string `json:"name"`
	Email    string `json:"email" pg:",unique"`
	Password string `json:"password"`
	Roles    []Role `json:"roles" pg:"many2many:roles_users_relation"`
}

type RoleUser struct {
	tableName struct{} `pg:"roles_users_relation"`
	UserId    int
	RoleId    int
}

type Role struct {
	Role string
}
