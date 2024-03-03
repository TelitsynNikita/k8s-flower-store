package main

import (
	"fmt"
	"github.com/TelitsynNikita/k8s-flower-store/auth_api/service"
	"github.com/gofiber/fiber/v3"
	"log"
)

func main() {
	app := fiber.New()

	app.Use("/api", func(c fiber.Ctx) error {
		fmt.Println("middleware is ok...")
		return c.Next()
	})

	app.Post("/api/sign_up", service.SignUp)
	app.Post("/api/sign_in", service.SignIn)

	log.Fatal(app.Listen(":8080"))
}
