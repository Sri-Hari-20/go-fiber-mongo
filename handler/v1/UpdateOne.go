package v1

import "github.com/gofiber/fiber/v2"

func UpdateOne(ctx *fiber.Ctx) error {
    return ctx.SendString("Hello, World!")
}