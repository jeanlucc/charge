package repositories

import (
	"github.com/theodo/scalab/src/database"
	"github.com/theodo/scalab/src/entities"
)

type projectRepository struct{}

func NewProjectRepository() projectRepository {
	return projectRepository{}
}

func (r *projectRepository) FindAll() (es []entities.Project, err error) {
	err = database.Db().Model(&es).Select()
	return
}

func (r *projectRepository) FindByOwner(u entities.User) (es []entities.Project, err error) {
	err = database.Db().Model((*entities.Project)(nil)).
		Join("JOIN projects_users_relation AS pu ON pu.project_id = project.id").
		Where("pu.user_id = ?", u.Id).Select(&es)
	return
}

func (r *projectRepository) Find(id int) (e entities.Project, err error) {
	e.Id = id
	err = database.Db().Select(&e)
	return
}

func (r *projectRepository) Create(pe *entities.Project) error {
	return database.Db().Insert(pe)
}
