package v1

import (
	"ToDoList/database"
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteMultiple(ctx *fiber.Ctx) error {
    listToDelete := [] string {}
    if err := ctx.BodyParser(&listToDelete); err != nil {
        log.Fatalln("DeleteMany receive list:", err)
        ctx.Status(400).JSON(&fiber.Map {
            "success" : false,
            "message" : err,
        })
    }

    listToDeleteHex := [] primitive.ObjectID {}
    for i := 0; i < len(listToDelete); i++ {
        idHex, err := primitive.ObjectIDFromHex(listToDelete[i])
        if err != nil {
            log.Fatalln("Hex conversion:", err)
            ctx.Status(500).JSON(&fiber.Map{
                "success" : false,
                "message" : err,
            })
        }
        listToDeleteHex = append(listToDeleteHex, idHex)
    } 

    result, err := database.Collection.DeleteMany(context.TODO(), bson.M{"_id" : bson.M{"$in" : listToDeleteHex}})
    if err != nil {
        log.Fatalln("Database delete many:", err)
        ctx.Status(500).JSON(&fiber.Map {
            "success" : false,
            "message" : err,
        })
    }

    return ctx.Status(200).JSON(&fiber.Map {
        "success" : true,
        "message" : result.DeletedCount,
    })
}