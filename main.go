package main

import (
	"log"

	"github.com/christ-saragih/tk-imanuel-bogor-be/config"
	"github.com/christ-saragih/tk-imanuel-bogor-be/controllers"
	"github.com/christ-saragih/tk-imanuel-bogor-be/database/seed"
	"github.com/christ-saragih/tk-imanuel-bogor-be/repositories"
	"github.com/christ-saragih/tk-imanuel-bogor-be/routes"
	"github.com/christ-saragih/tk-imanuel-bogor-be/services"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	seed.SeedAdmin()

	app := fiber.New()

	// user
	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	// teacher
	teacherRepo := repositories.NewTeacherRepository()
	teacherService := services.NewTeacherService(teacherRepo)
	teacherController := controllers.NewTeacherController(teacherService)

	// jumbotron
	jumbotronRepo := repositories.NewJumbotronRepository()
	jumbotronService := services.NewJumbotronService(jumbotronRepo)
	jumbotronController := controllers.NewJumbotronController(jumbotronService)

	routes.Setup(app, userController, teacherController, jumbotronController)

	port := config.AppConfig.AppPort
	log.Println("Server is running on port:", port)
	log.Fatal(app.Listen(":" + port))
}