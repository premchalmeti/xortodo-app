package utils

type OKResponse struct {
	Ok bool `json:"ok"`
}

type CreateTodoResponse struct {
	Ok     bool `json:"ok"`
	TodoId uint `json:"todo_id"`
}
