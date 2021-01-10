package models

// Todo model for Todo Records
type Todo struct {
	BaseModel
	Title string `json:"title" gorm:"not null"`
	Desc  string `json:"desc"`
	Done  bool   `json:"done"`
}
