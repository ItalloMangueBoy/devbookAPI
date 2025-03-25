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

// ListUserPosts: return all posts from a user
func GetUserPosts(w http.ResponseWriter, r *http.Request) {
	// read user ID
	userID, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		view.GenErrorTemplate(errors.New("invalid user ID")).Send(w, 422)
		return
	}

	// search posts in database
	posts, err := posthelper.ListUserPosts(userID)
	if err != nil {
		view.GenErrorTemplate(err).Send(w, 500)
		return
	}

	if len(posts) == 0 {
		view.GenErrorTemplate(errors.New("no posts found")).Send(w, 404)
		return
	}

	// return success
	view.JSON(w, 200, posts)
}

// LikePost: like a post
func LikePost(w http.ResponseWriter, r *http.Request) {
	handlePostReaction(w, r, posthelper.Like)
}

// DislikePost: dislike a post
func DislikePost(w http.ResponseWriter, r *http.Request) {
	handlePostReaction(w, r, posthelper.Dislike)
}

func handlePostReaction(w http.ResponseWriter, r *http.Request, reactionFunc func(int64) error) {
	// get request data
	ID, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		view.GenErrorTemplate(errors.New("invalid post ID")).Send(w, 422)
		return
	}

	// register reaction in database
	if err := reactionFunc(ID); err != nil {
		view.GenErrorTemplate(err).Send(w, 500)
		return
	}

	// return success
	w.WriteHeader(204)
}

// DeletePost: delete a post
func DeletePost(w http.ResponseWriter, r *http.Request) {
	// get request data
	authID, err := auth.GetAuthenticatedId(r)
	if err != nil {
		view.GenErrorTemplate(err).Send(w, 500)
		return
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
		view.GenErrorTemplate(errors.New("you cannot delete a post that is not yours")).Send(w, 403)
		return
	}

	// remove post from database
	if err := posthelper.Delete(postID); err != nil {
		view.GenErrorTemplate(err).Send(w, 500)
		return
	}

	// return success
	w.WriteHeader(204)
}
