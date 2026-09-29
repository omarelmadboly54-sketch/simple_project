package presenters

import "simple_project/internal/models"

type CreateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
	
}

func (c *CreateUserRequest) ToUser() models.User{
	
	return models.User{
		Name: c.Name,
		Email: c.Email,
		Password: c.Password,
		Role: "user",
	}
}

