package routes

import (
	"github.com/gleo08/fresherOCG/controllers"
	"github.com/gleo08/fresherOCG/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetUp(app *fiber.App) {

	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	app.Use(middlewares.IsAuthenticated)

	app.Put("/api/users/info", controllers.UpdateInfo)
	app.Put("api/users/password", controllers.UpdatePassword)

	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)

	app.Get("/api/users", controllers.AllUsers)
	app.Get("/api/user/:id", controllers.GetUserById)
	app.Delete("/api/user/:id", controllers.DeleteUserById)

	app.Get("/api/products", controllers.AllProducts)
	app.Post("/api/products", controllers.CreateProduct)
	app.Get("/api/products/:id", controllers.GetProductById)
	app.Put("/api/products/:id", controllers.UpdateProductById)
	app.Delete("/api/product/:id", controllers.DeleteProductById)

	app.Post("/api/upload", controllers.Upload)
	app.Static("/api/uploads", "./uploads")
}
