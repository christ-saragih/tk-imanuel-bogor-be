package routes

import (
	"log"

	"github.com/christ-saragih/tk-imanuel-bogor-be/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func Setup(
	app *fiber.App, 
	uc *controllers.UserController, 
	tc *controllers.TeacherController,
	jc *controllers.JumbotronController,
	cc *controllers.ContactController,
	bc *controllers.BlogController,
) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error load .env file")
	}

	app.Static("/uploads", "./uploads")

	app.Post("/v1/auth/register", uc.Register)
	app.Post("/v1/auth/login", uc.Login)


	app.Post("/v1/teachers", tc.CreateTeacher)
	app.Get("/v1/teachers", tc.GetTeachers)
	app.Get("/v1/teachers/:id", tc.GetTeacherDetail)
	app.Put("/v1/teachers/:id", tc.UpdateTeacher)
	app.Delete("/v1/teachers/:id", tc.DeleteTeacher)

	// Jumbotron Routes
	app.Get("/v1/jumbotron", jc.GetJumbotron)
	app.Put("/v1/jumbotron", jc.UpsertJumbotron)

	// Contact Routes
	app.Get("/v1/contact", cc.GetContact)
	app.Put("/v1/contact", cc.UpsertContact)

	// Blog Routes
	app.Post("/v1/blogs", bc.CreateBlog)
	app.Get("/v1/blogs", bc.GetBlogs)
	app.Get("/v1/blogs/:slug", bc.GetBlogDetail)
	app.Put("/v1/blogs/:slug", bc.UpdateBlog)
	app.Delete("/v1/blogs/:slug", bc.DeleteBlog)
}