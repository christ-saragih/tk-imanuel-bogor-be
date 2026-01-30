package controllers

import (
	"github.com/christ-saragih/tk-imanuel-bogor-be/models"
	"github.com/christ-saragih/tk-imanuel-bogor-be/services"
	"github.com/christ-saragih/tk-imanuel-bogor-be/utils"
	"github.com/gofiber/fiber/v2"
)

type JumbotronController struct {
	service services.JumbotronService
}

func NewJumbotronController(service services.JumbotronService) *JumbotronController {
	return &JumbotronController{service}
}

func (c *JumbotronController) GetJumbotron(ctx *fiber.Ctx) error {
	jumbotron, err := c.service.Get()
	if err != nil {
		return utils.Success(ctx, "Jumbotron data empty", nil)
	}

	return utils.Success(ctx, "Jumbotron retrieved successfully", jumbotron)
}

func (c *JumbotronController) UpsertJumbotron(ctx *fiber.Ctx) error {
	var jumbotron models.Jumbotron

	jumbotron.Title = ctx.FormValue("title")
	jumbotron.Description = ctx.FormValue("description")

	imagePath, err := utils.UploadFile(ctx, "image", "jumbotron")
	if err != nil {
		return utils.BadRequest(ctx, "Failed to upload image", err.Error())
	}
	jumbotron.Image = imagePath

	result, err := c.service.Upsert(&jumbotron)
	if err != nil {
		return utils.BadRequest(ctx, "Failed to upsert jumbotron", err.Error())
	}

	return utils.Success(ctx, "Success upsert jumbotron", result)
}