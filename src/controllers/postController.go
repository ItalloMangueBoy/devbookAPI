package controllers

import (
	postsforms "devbookAPI/src/forms/postsForms"
	"devbookAPI/src/helper/posthelper"
	"devbookAPI/src/view"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreatePost: create a new post
func CreatePost(w http.ResponseWriter, r *http.Request) {
	// read post data
	var form postsforms.Create

	if err := form.Prepare(r); err != nil {
		view.GenErrorTemplate(err).Send(w, 422)
		return
	}

	// insert post in database
	if err := posthelper.Create(&form); err != nil {
		view.GenErrorTemplate(err).Send(w, 500)
		return
	}

	// Return success
	w.WriteHeader(201)
}

// GetPost: return a post
func GetPost(w http.ResponseWriter, r *http.Request) {
	// read post ID
	ID, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		view.GenErrorTemplate(errors.New("invalid post ID")).Send(w, 422)
		return
	}

	// search post in database
	post, err := posthelper.GetByID(ID)
	if err != nil {
		view.GenErrorTemplate(err).Send(w, 500)
		return
	}

	// return success
	view.JSON(w, 200, post)
}

// DeletePost: delete a post
func DeletePost(w http.ResponseWriter, r *http.Request) {

}

// GetTimeline: return the user timeline
func GetTimeline(w http.ResponseWriter, r *http.Request) {

}
