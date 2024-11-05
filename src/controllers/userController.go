package controllers

import (
	"database/sql"
	"devbookAPI/src/helper"
	"devbookAPI/src/model"
	"devbookAPI/src/view"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
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

	users, err := helper.GetUsers(search)
	if err != nil {
		view.GenErrorTemplate(err).Send(w, 500)
		return
	}

	view.JSON(w, 200, users)
}

// GetUser: Search one user
func GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		view.GenErrorTemplate(err).Send(w, 422)
	}

	user, err := helper.GetUserById(id)

	if err == sql.ErrNoRows {
		view.GenErrorTemplate(err).Send(w, 404)
		return
	}

	if err != nil {
		view.GenErrorTemplate(err).Send(w, 500)
		return
	}

	view.JSON(w, 200, user)
}

// UpdateUser: Updates an user
func UpdateUser(w http.ResponseWriter, r *http.Request) {

}

// DeleteUser: Deletes an user
func DeleteUser(w http.ResponseWriter, r *http.Request) {

}
