package controllers

import (
	postsforms "devbookAPI/src/forms/postsForms"
	"devbookAPI/src/helper/posthelper"
	"devbookAPI/src/view"
	"net/http"
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

}

// DeletePost: delete a post
func DeletePost(w http.ResponseWriter, r *http.Request) {

}

// GetTimeline: return the user timeline
func GetTimeline(w http.ResponseWriter, r *http.Request) {

}
