package routes

import (
	"main/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// authController.go
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)

	// postController.go
	app.Post("/api/posts", controllers.CreatePost)
	app.Get("/api/posts", controllers.FindPosts)
	app.Get("/api/posts/:id", controllers.FindSinglePost)
	app.Get("/api/posts/:id", controllers.FindSinglePost)
	app.Put("/api/posts/:id", controllers.UpdatePost)
	app.Delete("/api/posts/:id", controllers.DeletePost)

	// commentsController.go
	app.Post("/api/comments", controllers.CreateComment)
	app.Get("/api/comments/:id", controllers.FindComments)
	app.Put("/api/comments/:id", controllers.UpdateComment)
	app.Delete("/api/comments/:id", controllers.DeleteComment)
}
