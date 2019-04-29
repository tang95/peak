package services

import (
	"github.com/globalsign/mgo/bson"
	"peak/forms"
	"peak/models"
	"peak/results"
)

type ProjectService struct{}

func (s ProjectService) CreateProject(form forms.CreateProjectForm) (result results.ProjectResult) {
	var projectModel models.ProjectModel
	project := models.Project{Name: form.Name}
	project.Id = bson.NewObjectId()
	err := projectModel.Create(project)
	if err != nil {
		return
	}
	result.Id = project.Id
	result.Name = project.Name
	return result
}
