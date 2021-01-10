package routers

import (
	"net/http"

	"github.com/gorilla/mux"

	"todo_app/controllers"
)

// SetupRouter function will create a new router and register
// all the handlers or controllers
func SetupRouter() {
	r := mux.NewRouter()

	todoRouter := r.PathPrefix("/todos/").Subrouter()
	todoRouter.HandleFunc("/", controllers.TodoListAPI).Methods(http.MethodGet)
	todoRouter.HandleFunc("/", controllers.TodoCreateAPI).Methods(http.MethodPost)
	todoRouter.HandleFunc("/done/", controllers.TodoDoneAPI).Methods(http.MethodPost)
	todoRouter.HandleFunc("/{id:[0-9]+}/", controllers.TodoDetailAPI).Methods(http.MethodGet)
	todoRouter.HandleFunc("/{id:[0-9]+}/", controllers.TodoUpdateAPI).Methods(http.MethodPut)
	todoRouter.HandleFunc("/{id:[0-9]+}/", controllers.TodoDeleteAPI).Methods(http.MethodDelete)

	// register router to path
	http.Handle("/", r)
}
