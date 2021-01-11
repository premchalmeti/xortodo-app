package utils

import "todo_app/models"

type OKResponse struct {
	Ok bool `json:"ok"`
}

type OKLoginResponse struct {
	Ok   bool        `json:"ok"`
	User models.User `json:"user"`
}

type CreateTodoResponse struct {
	Ok     bool `json:"ok"`
	TodoId uint `json:"todo_id"`
}
