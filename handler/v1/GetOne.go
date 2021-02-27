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

func GetOne(ctx *fiber.Ctx) error {
    id := ctx.Params("id")
    toDo := model.ToDo{}

    idHex, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        log.Fatalln("Hex conversion:", err)
    }

    err = database.Collection.FindOne(context.TODO(), bson.M{"_id" :  idHex}).Decode(&toDo)
    if err != nil {
        log.Fatalln("GetOne:", err)
        ctx.Status(500).JSON(&fiber.Map{
            "success" : false,
            "message" : err,
        })
    }
    
    // If the database read is successful
    ctx.Status(200).JSON(&fiber.Map {
        "success" : true,
        "data" : toDo,
    })

    return nil
}