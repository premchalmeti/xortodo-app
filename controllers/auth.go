package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"todo_app/models"
	"todo_app/utils"
)

// RegisterationAPI adds new user into db
func RegisterationAPI(w http.ResponseWriter, r *http.Request) {
	var userCredentials utils.Credentials
	json.NewDecoder(r.Body).Decode(&userCredentials)
	fmt.Println(userCredentials)

	db := models.GetDB()
	var newUser models.User
	newUser.Email = userCredentials.Email
	newUser.PassWord = utils.GetHashedPassword(userCredentials.PassWord)
	db.Create(&newUser)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp := utils.OKResponse{Ok: true}
	json.NewEncoder(w).Encode(resp)
}

// LoginAPI validates user credentials and responsible for session creation
func LoginAPI(w http.ResponseWriter, r *http.Request) {
	var credentials utils.Credentials
	json.NewDecoder(r.Body).Decode(&credentials)
	fmt.Println(credentials)

	db := models.GetDB()
	var user models.User
	result := db.Where("email = ?", credentials.Email).First(&user)

	if result.Error != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	} else if result.RowsAffected == 0 || !utils.CompareHashedPasswords(
		credentials.PassWord, user.PassWord,
	) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		resp := utils.OKLoginResponse{Ok: true, User: user}
		json.NewEncoder(w).Encode(&resp)
	}

}

// LogoutAPI wipes out the user session
func LogoutAPI(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp := utils.OKResponse{Ok: true}
	json.NewEncoder(w).Encode(resp)
}
