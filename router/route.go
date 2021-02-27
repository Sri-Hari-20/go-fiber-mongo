package router

import (
	v1 "ToDoList/handler/v1"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
    // Group
    api := app.Group("/api/v1")

    // Middleware if any

    // Routes
    // Version 1
    // Add a todo
    api.Post("/add", v1.AddTodo)

    // Get a todo
    api.Get("/getOne/:id", v1.GetOne)
    // Get all todos by user
    api.Get("/getAll", v1.GetAll)

    // Update a todo
    api.Put("/updateOne/:id", v1.UpdateOne)

    // Delete a todo
    api.Delete("/deleteOne/:id", v1.DeleteOne)
    api.Delete("/deleteMultiple", v1.DeleteMultiple)
}