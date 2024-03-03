package main

import (
	"fmt"
	"github.com/TelitsynNikita/k8s-flower-store/auth_api/services"
	"github.com/gofiber/fiber/v3"
	"log"
)

func main() {
	app := fiber.New()

	app.Use("/api", func(c fiber.Ctx) error {
		fmt.Println("middleware is ok...")
		return c.Next()
	})

	app.Post("/api/sign_up", services.SignUp)
	app.Post("/api/sign_in", services.SignIn)

	log.Fatal(app.Listen(":8080"))
}
