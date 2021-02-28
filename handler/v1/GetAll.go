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
    toDos := model.ToDos{}

    if err != nil {
        log.Fatalln("URI Decode:", err)
        return ctx.Status(500).JSON(&fiber.Map{
            "success" : false,
            "message" : err,
        })
    }

    cursor, err := database.Collection.Find(context.TODO(), bson.M{"by" : user})
    if err != nil {
        log.Fatalln("DB Read:", err)
        return ctx.Status(500).JSON(&fiber.Map{
            "success" : false,
            "message" : err,
        })
    }

    // Iterate through cursor
    for cursor.Next(context.TODO()) {
        toDo := model.ToDo{}
        err = cursor.Decode(&toDo)

        if err != nil {
            log.Fatalln("Cursor access:", err)
            return ctx.Status(500).JSON(&fiber.Map{
                "success" : false,
                "message" : err,
            })
        }

        toDos.ToDos = append(toDos.ToDos, toDo)
    }

    if err := cursor.Err(); err != nil {
        log.Fatal("Cursor failure:", err)
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