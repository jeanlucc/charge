package repositories

import (
	"github.com/theodo/scalab/src/entities"
)

func FindUser(id int) (user entities.User) {
	return getMappedUser("SELECT * FROM users WHERE id = $1", id)
}

func FindUserByEmail(name string) entities.User {
	return getMappedUser("SELECT * FROM users WHERE email = $1", name)
}

func getMappedUser(query string, args ...interface{}) (user entities.User) {
	getOneResult([]interface{}{&user.Id, &user.Name, &user.Password, &user.Email}, query, args...)
	return
}
