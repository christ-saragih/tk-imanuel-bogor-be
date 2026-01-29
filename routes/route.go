package routes

import (
	"log"

	"github.com/christ-saragih/tk-imanuel-bogor-be/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func Setup(app *fiber.App, uc *controllers.UserController) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error load .env file")
	}

	app.Post("/v1/auth/register", uc.Register)
}