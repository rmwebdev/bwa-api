package routes

import (
	"api/app/controller/user"
	"api/app/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/viper"
)

// Handle all request to route to controller
func Handle(app *fiber.App) {
	app.Use(cors.New())

	api := app.Group(viper.GetString("ENDPOINT"), middleware.Oauth2Authentication)

	api.Post("/users", user.PostUser)
	api.Get("/users/:id", user.GetUserID)
	api.Put("/users/:id", user.PutUser)
	api.Delete("/users/:id", user.DeleteUser)

}
