package repositories

import (
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

func (r *userRepository) Find(id int) (entities.User, error) {
	return r.getMapped("SELECT * FROM users WHERE id = $1", id)
}

func (r *userRepository) FindByEmail(name string) (entities.User, error) {
	return r.getMapped("SELECT * FROM users WHERE email = $1", name)
}

func (r *userRepository) Create(user entities.User) (entities.User, error) {
	return r.getMapped("INSERT INTO users (username, password, email) VALUES ($1, $2, $3) RETURNING *", user.Name, user.Password, user.Email)
}

func (r *userRepository) FindWithRoles(id int) (user entities.User, err error) {
	entityAndMapping := func() (interface{}, []interface{}) {
		var u usermap
		fields := r.getMapping(&u.user, &u.role)
		return &u, fields
	}
	results, err := getResults(entityAndMapping, "SELECT u.*, r.role FROM users AS u JOIN roles_users_relation AS ur ON u.id = ur.user_id JOIN roles AS r ON ur.role_id = r.id WHERE u.id = $1", id)
	if err != nil {
		return
	}
	users, err := r.resultsToEntities(results)
	if len(users) != 1 || err != nil {
		err = &GetMappedResultError{err}
		return
	}
	return users[0], nil
}

func (r *userRepository) getMapping(user *entities.User, otherFields ...interface{}) []interface{} {
	var password string
	return append([]interface{}{&user.Id, &user.Name, &password, &user.Email}, otherFields...)
}

func (r *userRepository) entityAndMapping() (interface{}, []interface{}) {
	var user entities.User
	fields := r.getMapping(&user)
	return &user, fields
}

func (r *userRepository) resultsToEntities(results []interface{}) (users []entities.User, err error) {
	var usersMap = make(map[int]entities.User)
	for _, result := range results {
		e, ok := (result).(*usermap)
		if ok {
			var tmpU entities.User
			if _, ok := usersMap[e.user.Id]; ok {
				tmpU = usersMap[e.user.Id]
			} else {
				tmpU = e.user
			}
			usersMap[e.user.Id] = r.mergeEntries(tmpU, *e)
		} else {
			err = &GetMappedResultError{}
			return
		}
	}

	users = make([]entities.User, 0, len(usersMap))
	for _, u := range usersMap {
		users = append(users, u)
	}
	return
}

func (r *userRepository) mergeEntries(u entities.User, e usermap) entities.User {
	u.Roles = append(u.Roles, e.role)
	return u
}

func (r *userRepository) getMapped(query string, args ...interface{}) (user entities.User, err error) {
	err = getOneResult(r.getMapping(&user), query, args...)
	return
}
