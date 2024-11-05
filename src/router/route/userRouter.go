package route

import (
	"devbookAPI/src/controllers"
)

var userRoutes = []Route{
	{
		URI:     "/users",
		Method:  "POST",
		Handler: controllers.CreateUser,
		Auth:    false,
	},

	{
		URI:     "/users",
		Method:  "GET",
		Handler: controllers.GetUsers,
		Auth:    false,
	},

	{
		URI:     "/users/{id}",
		Method:  "GET",
		Handler: controllers.GetUser,
		Auth:    false,
	},

	{
		URI:     "/users{id}",
		Method:  "PUT",
		Handler: controllers.UpdateUser,
		Auth:    false,
	},

	{
		URI:     "/users{id}",
		Method:  "DELETE",
		Handler: controllers.DeleteUser,
		Auth:    false,
	},
}
