package v1

import (
	"ToDoList/database"
	"ToDoList/model"
	"context"
	"log"
	"net/url"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAll(ctx *fiber.Ctx) error {
    user := ctx.Params("user")
    user, err := url.QueryUnescape(user)
    if err != nil {
        log.Fatalln("URI Decode:", err)
        return ctx.Status(400).JSON(&fiber.Map{
            "success" : false,
            "message" : err,
        })
    }

    toDos := model.ToDos{}

    cursor, err := database.Collection.Find(context.TODO(), bson.M{"by" : user})
    if err != nil {
        log.Fatalln("DB Read:", err)
        return ctx.Status(500).JSON(&fiber.Map{
            "success" : false,
            "message" : err,
        })
    }

    if err := cursor.Err(); err != nil {
        log.Fatal("Cursor failure:", err)
        return ctx.Status(500).JSON(&fiber.Map{
            "success" : false,
            "message" : err,
        })
    }

    if err = cursor.All(context.TODO(), &toDos.ToDos); err != nil {
        log.Fatalln("Cursor decode fail:", err)
        return ctx.Status(500).JSON(&fiber.Map{
            "success" : false,
            "message" : err,
        })
    }

    cursor.Close(context.TODO())
    
    return ctx.Status(200).JSON(&fiber.Map{
        "success" : true,
        "data" : toDos.ToDos,
    })
}