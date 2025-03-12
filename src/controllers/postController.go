package controllers

import (
	"devbookAPI/src/auth"
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

// GetTimeline: return the user timeline
func GetTimeline(w http.ResponseWriter, r *http.Request) {
	// get authenticated user ID
	ID, err := auth.GetAuthenticatedId(r)
	if err != nil {
		view.GenErrorTemplate(err).Send(w, 500)
		return
	}

	// search timeline posts in database
	timeline, err := posthelper.GetTimeline(ID)
	if err != nil {
		view.GenErrorTemplate(err).Send(w, 500)
		return
	}

	//return success
	view.JSON(w, 200, timeline)
}

// DeletePost: delete a post
func DeletePost(w http.ResponseWriter, r *http.Request) {
	// get request data
	authID, err := auth.GetAuthenticatedId(r)
	if err != nil {
		view.GenErrorTemplate(err).Send(w, 500)
	}

	postID, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		view.GenErrorTemplate(errors.New("invalid post ID")).Send(w, 422)
		return
	}

	// search post in database
	post, err := posthelper.GetByID(postID)
	if err != nil {
		view.GenErrorTemplate(err).Send(w, 500)
		return
	}

	// check if auth user is the post author
	if post.AuthorID != authID {
		err := errors.New("you cannot delete a post that is not yours")
		view.GenErrorTemplate(err).Send(w, 403)
	}

	// remove post from database
	if err := posthelper.Delete(postID); err != nil {
		view.GenErrorTemplate(err).Send(w, 500)
	}

	// return success
	w.WriteHeader(204)
}
