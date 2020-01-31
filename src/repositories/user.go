package repositories

import (
	"github.com/go-pg/pg/v9/orm"
	"github.com/theodo/scalab/src/database"
	"github.com/theodo/scalab/src/entities"
)

type userRepository struct{}

type usermap struct {
	user entities.User
	role string
}

func NewUserRepository() userRepository {
	return userRepository{}
}

func (r *userRepository) Find(id int) (e entities.User, err error) {
	err = r.baseQuery(&e).Where("id = ?", id).Select()
	return
}

func (r *userRepository) FindByEmail(email string) (e entities.User, err error) {
	err = r.baseQuery(&e).Where("email = ?", email).Select()
	return
}

func (r *userRepository) FindForLogin(email string) (e entities.User, err error) {
	err = database.Db().Model(&e).Relation("Roles").Limit(1).Where("email = ?", email).Select()
	return
}

func (r *userRepository) Create(pe *entities.User) (err error) {
	err = database.Db().Insert(pe)
	return
}

func (r *userRepository) baseQuery(pe *entities.User) *orm.Query {
	return database.Db().Model(pe).Relation("Roles").Limit(1).ExcludeColumn("password")
}
