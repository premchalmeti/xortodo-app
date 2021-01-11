package configs

import "fmt"

var DB_HOST = "localhost"
var DB_USER = "postgres"
var DB_PWD = "dummypwd"
var DB_NAME = "todo_db"
var DB_PORT = 5432

var DSN = fmt.Sprintf(
	"host=%s user=%s password=%s dbname=%s port=%d",
	DB_HOST,
	DB_USER,
	DB_PWD,
	DB_NAME,
	DB_PORT,
)
