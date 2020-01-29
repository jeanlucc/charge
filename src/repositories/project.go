package repositories

import (
	"github.com/theodo/scalab/src/entities"
)

type projectRepository struct{}

func NewProjectRepository() projectRepository {
	return projectRepository{}
}

func (r *projectRepository) FindAll() ([]entities.Project, error) {
	return r.getMappedProjectList("SELECT * FROM projects")
}

func (r *projectRepository) FindByOwner(u entities.User) ([]entities.Project, error) {
	return r.getMappedProjectList("SELECT id, name FROM projects AS p JOIN projects_users_relation AS pu ON p.id = pu.project_id AND pu.user_id = $1", u.Id)
}

func (r *projectRepository) Find(id int) (entities.Project, error) {
	return r.getMappedProject("SELECT * FROM projects WHERE id = $1", id)
}

func (r *projectRepository) Create(project entities.Project) (entities.Project, error) {
	return r.getMappedProject("INSERT INTO projects (name) VALUES ($1) RETURNING *", project.Name)
}

func (r *projectRepository) getProjectMapping(project *entities.Project) []interface{} {
	return []interface{}{&project.Id, &project.Name}
}

func (r *projectRepository) projectToMapRetriever() (interface{}, []interface{}) {
	var project entities.Project
	fields := r.getProjectMapping(&project)
	return &project, fields
}

func (r *projectRepository) resultsToEntities(results []interface{}) (projects []entities.Project, err error) {
	for _, result := range results {
		if project, ok := (result).(*entities.Project); ok {
			projects = append(projects, *project)
		} else {
			err = &GetMappedResultError{}
			return
		}
	}
	return
}

func (r *projectRepository) getMappedProject(query string, args ...interface{}) (project entities.Project, err error) {
	fields := r.getProjectMapping(&project)
	err = getOneResult(fields, query, args...)
	return
}

func (r *projectRepository) getMappedProjectList(query string, args ...interface{}) (projects []entities.Project, err error) {
	results, err := getResults(r.projectToMapRetriever, query, args...)
	if err != nil {
		return
	}
	return r.resultsToEntities(results)
}
