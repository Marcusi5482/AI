package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
)

func main() {

	app := fiber.New()

	app.Get("/*", static.New("./img"))

	app.Get("/img*", static.New("./img"))

	app.Get("/", static.New("./index.html"))

	log.Fatal(app.Listen(":3000"))
}
