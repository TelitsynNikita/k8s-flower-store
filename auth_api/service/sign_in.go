package service

import (
	"encoding/json"
	"fmt"
	"github.com/TelitsynNikita/k8s-flower-store/auth_api/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"net/http"
)

func SignIn(c fiber.Ctx) error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	var body models.LogoPass
	err := json.Unmarshal(c.Request().Body(), &body)
	if err != nil {
		return c.JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "cannot unmarshal body",
		})
	}

	err = validate.Struct(body)
	if err != nil {
		return c.JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": fmt.Sprintf("validation error, %s", err),
		})
	}

	return c.JSON(body)
}
