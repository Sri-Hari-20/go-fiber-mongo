package v1

import (
	"ToDoList/database"
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteOne(ctx *fiber.Ctx) error {
    id := ctx.Params("id")
    idHex, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        log.Fatalln("Hex conversion:", err)
        ctx.Status(500).JSON(&fiber.Map{
            "success" : false,
            "message" : err,
        })
    }

    result, err := database.Collection.DeleteOne(context.TODO(), bson.M{"_id": idHex})
    if err != nil {
        log.Fatalln("DeleteOne:", err)
        ctx.Status(500).JSON(&fiber.Map{
            "success" : false,
            "message" : err,
        })
    }

    return ctx.Status(200).JSON(&fiber.Map{
        "success" : true,
        "data" : result.DeletedCount,
    })
}