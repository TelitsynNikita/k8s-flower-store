package main

import (
	"encoding/json"
	"fmt"
	"github.com/TelitsynNikita/k8s-flower-store/catalog_flowers_api/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"log"
	"net/http"
	"strconv"
)

var storage = []models.Flower{
	{
		FlowerID:    1,
		FlowerName:  "Роза красная",
		FlowerPrice: 250,
	},
	{
		FlowerID:    2,
		FlowerName:  "Роза белая",
		FlowerPrice: 220,
	},
	{
		FlowerID:    3,
		FlowerName:  "Роза розовая",
		FlowerPrice: 200,
	},
	{
		FlowerID:    4,
		FlowerName:  "Тюльпан красный",
		FlowerPrice: 150,
	},
	{
		FlowerID:    5,
		FlowerName:  "Тюльпан белый",
		FlowerPrice: 120,
	},
	{
		FlowerID:    6,
		FlowerName:  "Тюльпан жёлтый",
		FlowerPrice: 100,
	},
}

func main() {
	app := fiber.New()

	app.Use("/api", func(c fiber.Ctx) error {
		fmt.Println("middleware is ok...")
		return c.Next()
	})

	//example
	//request - http://localhost:8080/api/get_all_flowers
	//response -
	//{
	//	FlowerID:    1,
	//	FlowerName:  "Роза красная",
	//	FlowerPrice: 250,
	//},
	//{
	//	FlowerID:    2,
	//	FlowerName:  "Роза белая",
	//	FlowerPrice: 220,
	//},...
	app.Get("/api/get_all_flowers", func(c fiber.Ctx) error {
		return c.JSON(storage)
	})

	//example
	//request - http://localhost:8080/api/get_flower_by_id/1
	//response -
	//{
	//	FlowerID:    1,
	//	FlowerName:  "Роза красная",
	//	FlowerPrice: 250,
	//},
	app.Get("/api/get_flower_by_id/:flower_id", func(c fiber.Ctx) error {
		flowerID, err := strconv.ParseUint(c.Params("flower_id"), 10, 64)
		if err != nil {
			return c.JSON(fiber.Map{
				"status":  http.StatusBadRequest,
				"message": "cannot strconv.ParseUint",
			})
		}

		for i := range storage {
			if storage[i].FlowerID == flowerID {
				return c.JSON(storage[i])
			}
		}

		return c.JSON(fiber.Map{
			"status":  http.StatusNotFound,
			"message": "cannot found the flower",
		})
	})

	//example
	//request - http://localhost:8080/api/add_new_flower
	//response -
	//{
	//	"message": "Success",
	//	"status": 200
	//}
	app.Post("/api/add_new_flower", func(c fiber.Ctx) error {
		validate := validator.New(validator.WithRequiredStructEnabled())
		var flower models.Flower
		err := json.Unmarshal(c.Request().Body(), &flower)
		if err != nil {
			return c.JSON(fiber.Map{
				"status":  http.StatusBadRequest,
				"message": "cannot unmarshal body",
			})
		}

		err = validate.Struct(flower)
		if err != nil {
			return c.JSON(fiber.Map{
				"status":  http.StatusBadRequest,
				"message": fmt.Sprintf("validation error, %s", err),
			})
		}

		flower.FlowerID = uint64(len(storage) + 1)
		storage = append(storage, flower)

		return c.JSON(fiber.Map{
			"status":  http.StatusOK,
			"message": "Success",
		})
	})

	log.Fatal(app.Listen(":8080"))
}
