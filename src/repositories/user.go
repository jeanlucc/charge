package repositories

import (
	"github.com/theodo/scalab/src/entities"
)

type UserRepository interface {
	Find(id int) (entities.User, error)
	FindByEmail(name string) (entities.User, error)
	Create(user entities.User) (entities.User, error)
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return userRepository{}
}

func (r userRepository) Find(id int) (entities.User, error) {
	return r.getMappedUser("SELECT * FROM users WHERE id = $1", id)
}

func (r userRepository) FindByEmail(name string) (entities.User, error) {
	return r.getMappedUser("SELECT * FROM users WHERE email = $1", name)
}

func (r userRepository) Create(user entities.User) (entities.User, error) {
	return r.getMappedUser("INSERT INTO users (username, password, email) VALUES ($1, $2, $3) RETURNING *", user.Name, user.Password, user.Email)
}

func (r userRepository) getMappedUser(query string, args ...interface{}) (user entities.User, err error) {
	err = getOneResult([]interface{}{&user.Id, &user.Name, &user.Password, &user.Email}, query, args...)
	return
}
