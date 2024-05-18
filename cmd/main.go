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
		AllowOrigins:     "http://localhost:3000",
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
		port = "3000"
	}
	log.Fatal(app.Listen(":" + port))

}
