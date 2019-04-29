package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"peak/forms"
	"peak/results"
	"peak/services"
)

type ProjectController struct{}

func (ctrl ProjectController) Create(ctx *gin.Context) {
	var createProjectForm forms.CreateProjectForm
	if ctx.BindJSON(&createProjectForm) != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"message": "Invalid form", "form": createProjectForm})
		ctx.Abort()
		return
	}

	var projectService services.ProjectService
	var result results.ProjectResult
	result = projectService.CreateProject(createProjectForm)
	ctx.JSON(200, gin.H{"ceshi": "ceshi", "URl": result})
}

func (ctrl ProjectController) List(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"ceshi": "ceshi", "URl": "ceshi"})
}
