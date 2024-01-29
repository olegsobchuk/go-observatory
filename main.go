package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/requestid"
)

func main() {
	app := fiber.New()
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format: "${time} | ${status} | ${latency} | ${ip} | ${method} | ${path} ${queryParams} | ${error} \n",
	}))

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello world!")
	})

	app.Listen(":3000")
}
