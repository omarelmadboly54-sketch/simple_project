package presenters

import "simple_project/internal/models"

type CreateMovieReq struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Image       string `json:"image" binding:"required"`
	Genre       string `json:"genre" binding:"required"`
}

func (m *CreateMovieReq) ToMovie() models.Movie{
	return models.Movie{
		Title: m.Title,
		Description: m.Description,
		Image: m.Image,
		Genre: m.Genre,
	}
}
