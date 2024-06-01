package routes

import (
	"github.com/asakshat/go_blog/internal/controllers"
	"github.com/asakshat/go_blog/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

func Auth(app *fiber.App) {

	// Authentcation routes
	app.Post("/api/user/signup", controllers.Signup) // Signup
	app.Post("/api/user/login", controllers.Login)   // Login w Cookies
	app.Post("/api/user/logout", controllers.Logout) // Logout w removing cookies
	app.Get("/api/user", controllers.User)           // get logged in user

}

func Blog(app *fiber.App) {
	app.Get("/api/blog/all", controllers.GetAllPosts)
	app.Get("/api/blog/:post_id", controllers.GetPostWithIdHandler)
	app.Get("/api/blog/all/:user_id", controllers.GetAllPostByUserId)

	app.Use(middlewares.Authenticate)
	// Post blog routes
	app.Post("/api/blog/:user_id", controllers.CreateBlog)
	app.Put("/api/blog/edit/:user_id/:post_id", controllers.EditPost)
	app.Delete("/api/blog/delete/:user_id/:post_id", controllers.DeletePost)

}

func Likes(app *fiber.App) {
	app.Use(middlewares.Authenticate)

	app.Post("/api/blog/like/:user_id/:post_id", controllers.LikePost)
	app.Post("/api/blog/unlike/:user_id/:post_id", controllers.UnlikePost)
}

func Comments(app *fiber.App) {
	app.Use(middlewares.Authenticate)

	app.Post("/api/blog/comment/post/:user_id/:post_id", controllers.PostComment)
	app.Put("/api/blog/comment/edit/:user_id/:comment_id", controllers.EditComment)
	app.Delete("/api/blog/comment/delete/:user_id/:comment_id", controllers.DeleteComment)

}
