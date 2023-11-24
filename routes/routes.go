// routes/routes.go
package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber/controllers"
	"go-fiber/repositories"
	"go-fiber/services"
)

type StatusRoutes struct {
	NameUrlRoute string
	Status       bool
}

type RoutesData struct {
	RouteTitle  string
	NamesRoutes []StatusRoutes
}

func SetupRoutes(app *fiber.App) {

	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(&userRepo)
	userController := controllers.NewUserController(&userService)

	api := app.Group("/api/v1")
	api.Get("/users", userController.GetAll)
	api.Post("/users", userController.Create)
	api.Get("/users/:id", userController.Show)

	// Route untuk endpoint ROUTE HOME (/)
	app.Get("/", func(ctx *fiber.Ctx) error {
		routeName := RoutesData{
			RouteTitle: "http://localhost:3000/",
			NamesRoutes: []StatusRoutes{
				{NameUrlRoute: "users/", Status: false},
				{NameUrlRoute: "users?{params}", Status: true},
			},
		}
		// Render template dengan menggunakan data
		return ctx.Render("index", fiber.Map{
			"Data": routeName,
		}, "index")
	})

}
