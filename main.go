package main

import (
	"fmt"
	pingerino "go-observatory/app/services/pingerino"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func main() {
	app := fiber.New()
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format: "${time} | ${status} | ${latency} | ${ip} | ${method} | ${path} ${queryParams} | ${error} \n",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world!")
	})

	// payload := []byte(`[{"destination": "ukr.net", "port": "80"}]`)
	// storage.Place()
	for _, pngr := range pingerino.Pingers {
		fmt.Printf("Start - %s\n", pngr.Alias)
		err := pingerino.Ping(&pngr)
		if err != nil {
			continue
		}
	}

	log.Fatal(app.Listen(":3000"))
}
