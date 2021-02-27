package main

import (
	"log"

	"ToDoList/config"
	"ToDoList/database"
	"ToDoList/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
    serverPort := config.Config("PORT")
    // Try connecting to the database
    err := database.Connect()

    if err != nil {
        log.Fatal(err)
    }
    
    app := fiber.New()
    router.SetupRoutes(app)

    app.Listen(":" + serverPort)
}