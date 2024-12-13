package routes

import (
	"devbookAPI/src/controllers"
)

var authRoutes = []Route{
	{
		URI:     "/login",
		Method:  "POST",
		Handler: controllers.Login,
		Auth:    false,
	},
}