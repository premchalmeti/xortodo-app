package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"todo_app/models"
	"todo_app/utils"
)

// TodoCreateAPI handles Create Todo request
func TodoCreateAPI(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)

	db := models.GetDB()
	db.Create(&todo)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp := utils.CreateTodoResponse{Ok: true, TodoId: todo.ID}
	json.NewEncoder(w).Encode(resp)
}

// TodoListAPI handles fetch Todo request
func TodoListAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Todos []models.Todo
	db := models.GetDB()
	db.Find(&Todos)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Todos)
}

// TodoDetailAPI handles detail view for a Todo request
func TodoDetailAPI(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	db := models.GetDB()
	var todo models.Todo
	result := db.First(&todo, vars["id"])
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			http.Error(w, "Not Found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(todo)
	}
}

// TodoDeleteAPI handles delete Todo request
func TodoDeleteAPI(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	db := models.GetDB()
	result := db.Delete(&models.Todo{}, vars["id"])
	if result.RowsAffected == 0 {
		http.Error(w, "Record not found", http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusNoContent)
		json.NewEncoder(w).Encode("Deleted")
	}
}

// TodoUpdateAPI handles update Todo request
func TodoUpdateAPI(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	db := models.GetDB()
	var todo models.Todo
	result := db.First(&todo, vars["id"])
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			http.Error(w, "Not Found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	} else {
		var updatedTodo models.Todo
		json.NewDecoder(r.Body).Decode(&updatedTodo)

		todo.Title = updatedTodo.Title
		todo.Desc = updatedTodo.Desc
		todo.Done = updatedTodo.Done

		db.Save(&todo)

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		resp := utils.OKResponse{Ok: true}
		json.NewEncoder(w).Encode(resp)
	}
}

// TodoDoneAPI handles request to mark a Todo as `Done`
func TodoDoneAPI(w http.ResponseWriter, r *http.Request) {
	var doneRequest utils.DoneRequest
	json.NewDecoder(r.Body).Decode(&doneRequest)

	db := models.GetDB()
	var todo models.Todo
	result := db.First(&todo, doneRequest.Id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			http.Error(w, "Not Found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	} else {
		todo.Done = doneRequest.Done
		db.Save(&todo)

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		resp := utils.OKResponse{Ok: true}
		json.NewEncoder(w).Encode(resp)
	}
}
