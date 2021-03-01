package v1

import (
	"ToDoList/database"
	"ToDoList/model"
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateOne(ctx *fiber.Ctx) error {
    id := ctx.Params("id")
    currentToDo := model.ToDo{}

    idHex, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        log.Fatalln("Hex conversion:", err)
    }

    // Get the current ToDo information
    err = database.Collection.FindOne(context.TODO(), bson.M{"_id" :  idHex}).Decode(&currentToDo)
    if err != nil {
        log.Fatalln("UpdateOne:", err)
        ctx.Status(500).JSON(&fiber.Map{
            "success" : false,
            "message" : err,
        })
    }

    newToDo := model.ToDo{}

    // Parse PUT data
    if err := ctx.BodyParser(&newToDo); err != nil {
        log.Fatalln("Wrong data format received from client")
        return ctx.Status(400).JSON(&fiber.Map{
            "success" : false,
            "message": err,
        })
    }

    // Get the bson.M that is required to send
    updateBson := model.PrepareBsonTodo(newToDo)

    log.Println(updateBson)

    result, err := database.Collection.UpdateOne(context.TODO(), bson.M{"_id" : idHex}, updateBson)

    if err != nil {
        log.Fatalln("UpdateOne:", err)
        ctx.Status(500).JSON(&fiber.Map{
            "success" : false,
            "message" : err,
        })
    }

    return ctx.Status(200).JSON(&fiber.Map{
        "success" : true,
        "data" : result.ModifiedCount,
    })
}