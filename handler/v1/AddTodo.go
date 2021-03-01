package v1

import (
	"ToDoList/database"
	"ToDoList/model"
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func AddTodo(ctx *fiber.Ctx) error {
    // Create a new instance
    toDo := model.ToDo{}

    // Parse POST data
    if err := ctx.BodyParser(&toDo); err != nil {
        log.Fatalln("Wrong data format received from client")
        return ctx.Status(400).JSON(&fiber.Map{
            "success" : false,
            "message": err,
        })
    }
    // Add the current timestamp
    toDo.CreatedOn = time.Now().UnixNano() / int64(time.Millisecond)

    // Insert to DB
    insertResult, err := database.Collection.InsertOne(context.Background(), toDo)

    if err != nil {
        log.Fatalln("Insert:", err)
        return ctx.Status(500).JSON(&fiber.Map{
            "success" : false,
            "message" : err,
        })
    }

    return ctx.Status(200).JSON(&fiber.Map{
        "success" : true,
        "data" : insertResult.InsertedID,
    })
}