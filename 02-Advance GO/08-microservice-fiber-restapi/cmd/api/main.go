package main

import (
	"08-microservice-fiber-restapi/internal/handlers"
	"08-microservice-fiber-restapi/internal/repositories"
	"08-microservice-fiber-restapi/internal/services"
	"08-microservice-fiber-restapi/pkg/database"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	var port = 3000

	database.ConnectDB()

	app := fiber.New(fiber.Config{
		AppName: "GO Fiber REST API Microservice v1.0",
	})

	// Middleware Log API
	app.Use(logger.New())

	// Setup Dependency Injection
	userRepository := &repositories.UserRepository{}
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	// Routing
	api := app.Group("/api/v1")
	api.Get("/user/id/:id", userHandler.GetUserByID)
	api.Get("/user/name/:name", userHandler.GetUserByName)
	api.Get("/users", userHandler.GetAllUsers)
	api.Post("/user", userHandler.CreateUser)
	api.Put("/user/:id", userHandler.UpdateUser)
	api.Delete("/user/:id", userHandler.DeleteUser)

	// Error handling Route not found
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{"message": "Route Not Found"})
	})

	fmt.Printf("Server starting on port %s...\n", port)
	app.Listen(fmt.Sprintf(":%d", port))
}
