package entities

type Project struct {
	Id    int    `json:"Id"`
	Name  string `json:"name" pg:",unique"`
	Owner []User `json:"owner" pg:"many2many:projects_users_relation,joinFK:user_id"`
}
