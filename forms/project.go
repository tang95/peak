package forms

type CreateProjectForm struct {
	Name string `json:"name" binding:"required,max=100,min=5"`
}
