package controllers

import (
	"database/sql"
	"devbookAPI/src/auth"
	forms "devbookAPI/src/forms/user_forms"
	"devbookAPI/src/helper/followhelper"
	userhelper "devbookAPI/src/helper/userHelper"
	"devbookAPI/src/model"
	"devbookAPI/src/view"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// PostUser: Creates a new user
func PostUser(w http.ResponseWriter, r *http.Request) {
	var user model.User

	if err := user.Prepare(r, "register"); err != nil {
		view.GenErrorTemplate(err).Send(w, 422)
		return
	}

	if err := userhelper.CreateUser(&user); err != nil {
		view.GenErrorTemplate(err).Send(w, 500)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("/users/%d", user.Id))
	w.WriteHeader(201)
}

// GetUsers: Search all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	search := strings.ToLower(r.URL.Query().Get("user"))
	search = strings.TrimSpace(search)

	users, err := userhelper.SearchUsers(search)
	if err != nil {
		view.GenErrorTemplate(err).Send(w, 500)
		return
	}

	view.JSON(w, 200, users)
}

// GetUser: Search one user
func GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		view.GenErrorTemplate(err).Send(w, 422)
	}

	user, err := userhelper.SearchUserById(id)

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
func PutUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	var err error

	user.Id, err = strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		view.GenErrorTemplate(err).Send(w, 422)
		return
	}

	authId, err := auth.GetAuthenticatedId(r)
	if err != nil {
		view.GenErrorTemplate(err).Send(w, 401)
		return
	}

	if err := auth.CheckUserPermision(user.Id, authId); err != nil {
		view.GenErrorTemplate(err).Send(w, 403)
		return
	}

	if err := user.Prepare(r, "update"); err != nil {
		view.GenErrorTemplate(err).Send(w, 422)
		return
	}

	if err := userhelper.UpdateUser(&user); err != nil {
		view.GenErrorTemplate(err).Send(w, 500)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("/users/%d", user.Id))
	w.WriteHeader(204)
}

// PutUserPassword: Updates an user password
func PutUserPassword(w http.ResponseWriter, r *http.Request) {
	var form forms.PasswordForm
	var err error

	// read request
	id, err := auth.GetAuthenticatedId(r)
	if err != nil {
		view.GenErrorTemplate(err).Send(w, 401)
		return
	}

	if err := form.Prepare(r, id); err != nil {
		view.GenErrorTemplate(err).Send(w, 422)
		return
	}

	// update password operation
	if err := userhelper.UpdateUserPassword(id, form.NewPassword); err != nil {
		view.GenErrorTemplate(err).Send(w, 500)
		return
	}

	// send response
	w.WriteHeader(204)
}

// DeleteUser: Deletes an user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		view.GenErrorTemplate(err).Send(w, 422)
		return
	}

	authId, err := auth.GetAuthenticatedId(r)
	if err != nil {
		view.GenErrorTemplate(err).Send(w, 401)
		return
	}

	if err := auth.CheckUserPermision(id, authId); err != nil {
		view.GenErrorTemplate(err).Send(w, 403)
		return
	}

	_, err = userhelper.SearchUserById(id)
	if err == sql.ErrNoRows {
		view.GenErrorTemplate(err).Send(w, 404)
		return
	}

	if err != nil {
		view.GenErrorTemplate(err).Send(w, 500)
		return
	}

	err = userhelper.DeleteUser(id)
	if err != nil {
		view.GenErrorTemplate(err).Send(w, 500)
		return
	}

	w.WriteHeader(204)
}

// GetUserFollowers: Search all followers of an user
func GetUserFollowers(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		view.GenErrorTemplate(err).Send(w, 422)
	}

	followers, err := followhelper.GetUserFollowers(id)

	if err != nil {
		view.GenErrorTemplate(err).Send(w, 500)
		return
	}

	fmt.Println(followers)

	view.JSON(w, 200, followers)
}

// GetUserFollows: Search all users that an user follows
func GetUserFollows(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		view.GenErrorTemplate(err).Send(w, 422)
	}

	followers, err := followhelper.GetUserFollows(id)

	if err != nil {
		view.GenErrorTemplate(err).Send(w, 500)
		return
	}

	fmt.Println(followers)

	view.JSON(w, 200, followers)
}

// FollowUser: Follows an user
func FollowUser(w http.ResponseWriter, r *http.Request) {
	followerID, err := auth.GetAuthenticatedId(r)
	if err != nil {
		view.GenErrorTemplate(err).Send(w, 401)
		return
	}

	userID, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		view.GenErrorTemplate(err).Send(w, 400)
		return
	}

	if userID == followerID {
		view.GenErrorTemplate(errors.New("you cannot follow yourself")).Send(w, 403)
		return
	}

	if err := followhelper.FollowUser(userID, followerID); err != nil {
		view.GenErrorTemplate(err).Send(w, 500)
		return
	}

	w.WriteHeader(201)
}

// UnfollowUser: Unfollows an user
func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	followerID, err := auth.GetAuthenticatedId(r)
	if err != nil {
		view.GenErrorTemplate(err).Send(w, 401)
		return
	}

	userID, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		view.GenErrorTemplate(err).Send(w, 400)
		return
	}

	if userID == followerID {
		view.GenErrorTemplate(errors.New("you cannot follow yourself")).Send(w, 403)
		return
	}

	if err := followhelper.UnfollowUser(userID, followerID); err != nil {
		view.GenErrorTemplate(err).Send(w, 500)
		return
	}

	w.WriteHeader(204)
}
