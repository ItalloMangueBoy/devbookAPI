package routes

import (
	"devbookAPI/src/controllers"
)

var postRoutes = []Route{
	{
		URI:     "/posts",
		Method:  "POST",
		Handler: controllers.CreatePost,
		Auth:    true,
	},

	{
		URI:     "/timeline",
		Method:  "GET",
		Handler: controllers.GetTimeline,
		Auth:    true,
	},

	{
		URI:     "/posts/{id}",
		Method:  "GET",
		Handler: controllers.GetPost,
		Auth:    true,
	},

	{
		URI:     "/posts/{id}",
		Method:  "DELETE",
		Handler: controllers.DeletePost,
		Auth:    true,
	},
}
