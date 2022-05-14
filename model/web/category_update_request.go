package web

type CategoryUpdateRequest struct {
	Id int `json:"id" validate:"required,min=1"`
	Name string `json:"name" validate:"required,max=200,min=1"`
}