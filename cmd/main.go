package main

import (
	"log"
	"os"

	"github.com/asakshat/go_blog/db"
	"github.com/asakshat/go_blog/internal/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://localhost:3000, http://localhost:5173 ,https://comforting-speculoos-51d893.netlify.app,https://iot-blog-project.netlify.app,https://blogpost-go.netlify.app/",
	}))
	db.Connect()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("HELLO, TEST GOLANG BLOG. IF THIS WORKS BACKEND WORKS!")
	})
	routes.Auth(app)
	routes.Blog(app)
	routes.Likes(app)
	routes.Comments(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(app.Listen(":" + port))

}
