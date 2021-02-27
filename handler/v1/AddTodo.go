package v1

import "github.com/gofiber/fiber/v2"

func AddTodo(ctx *fiber.Ctx) error {
    return ctx.SendString("Hello, World!")
}