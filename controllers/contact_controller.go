package controllers

import (
	"github.com/christ-saragih/tk-imanuel-bogor-be/models"
	"github.com/christ-saragih/tk-imanuel-bogor-be/services"
	"github.com/christ-saragih/tk-imanuel-bogor-be/utils"
	"github.com/gofiber/fiber/v2"
)

type ContactController struct {
	service services.ContactService
}

func NewContactController(service services.ContactService) *ContactController {
	return &ContactController{service}
}

func (c *ContactController) GetContact(ctx *fiber.Ctx) error {
	contact, err := c.service.Get()
	if err != nil {
		return utils.Success(ctx, "Contact data empty", nil)
	}

	return utils.Success(ctx, "Contact retrieved successfully", contact)
}

func (c *ContactController) UpsertContact(ctx *fiber.Ctx) error {
	contact := new(models.Contact)

	if err := ctx.BodyParser(contact); err != nil {
		return utils.BadRequest(ctx, "Failed to parse data", err.Error())
	}

	result, err := c.service.Upsert(contact)
	if err != nil {
		return utils.BadRequest(ctx, "Failed to upsert contact", err.Error())
	}

	return utils.Success(ctx, "Success upsert contact", result)
}