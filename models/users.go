package models

// User model for User Records
type User struct {
	BaseModel
	Email    string `json:"email"`
	PassWord string `json:"-"`
}
