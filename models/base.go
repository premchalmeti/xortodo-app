package models

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"todo_app/configs"
)

var db *gorm.DB

// BaseModel provides default fields
type BaseModel struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func init() {
	conn, err := gorm.Open(postgres.Open(configs.DSN), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db = conn

	db.AutoMigrate(&Todo{})
}

// GetDB exported interface to access DB instance
func GetDB() *gorm.DB {
	return db
}
