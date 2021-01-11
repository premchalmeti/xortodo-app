package utils

type DoneRequest struct {
	Id   uint ` json:"id"`
	Done bool `json:"done"`
}

type Credentials struct {
	Email    string `json:"email"`
	PassWord string `json:"password"`
}
