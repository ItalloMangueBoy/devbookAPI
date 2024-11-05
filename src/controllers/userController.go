package controllers

import (
	"devbookAPI/src/helper"
	"devbookAPI/src/model"
	"devbookAPI/src/view"
	"fmt"
	"net/http"
	"strings"
)

// CreateUser: Creates a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User

	if err := user.Prepare(r); err != nil {
		view.GenErrorTemplate(err).Send(w, 422)
		return
	}

	if err := helper.CreateUser(&user); err != nil {
		view.GenErrorTemplate(err).Send(w, 500)
		return
	}

	w.WriteHeader(201)
	w.Header().Set("Location", fmt.Sprintf("/users/%d", user.Id))
}

// GetUsers: Search all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	search := strings.ToLower(r.URL.Query().Get("user"))
	search = strings.TrimSpace(search)

	users, err := helper.SearchUser(search)
	if err != nil {
		view.GenErrorTemplate(err).Send(w, 500)
		return
	}

	view.JSON(w, 200, users)
}

// GetUser: Search one user
func GetUser(w http.ResponseWriter, r *http.Request) {

}

// UpdateUser: Updates an user
func UpdateUser(w http.ResponseWriter, r *http.Request) {

}

// DeleteUser: Deletes an user
func DeleteUser(w http.ResponseWriter, r *http.Request) {

}
