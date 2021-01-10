package main

import (
	"fmt"
	"log"
	"net/http"

	"todo_app/models"
	"todo_app/routers"
)

var Todos []models.Todo

func main() {
	routers.SetupRouter()

	fmt.Println("Started Webserver on Port 8000")
	log.Fatalln(http.ListenAndServe(":8000", nil))
}
