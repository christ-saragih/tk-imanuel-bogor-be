package controllers

import (
	"github.com/christ-saragih/tk-imanuel-bogor-be/models"
	"github.com/christ-saragih/tk-imanuel-bogor-be/services"
	"github.com/christ-saragih/tk-imanuel-bogor-be/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
)

type UserController struct {
	service services.UserService
}

func NewUserController(service services.UserService) *UserController {
	return &UserController{service}
}

func (c *UserController) Register(ctx *fiber.Ctx) error {
	user := new(models.User)

	if err := ctx.BodyParser(user); err != nil {
		return utils.BadRequest(ctx, "Failed to parsing data", err.Error())
	}

	if err := c.service.Register(user); err != nil {
		return utils.BadRequest(ctx, "Failed to register user", err.Error())
	}


	var userResp models.UserResponse
	_ = copier.Copy(&userResp,&user)

	return utils.Success(ctx, "Success register user", userResp)
}

func (c *UserController) Login(ctx *fiber.Ctx) error {
	var body struct {
		Email		string	`json:"email"`
		Password	string 	`json:"password"`
	}

	if err := ctx.BodyParser(&body); err != nil {
		return utils.BadRequest(ctx, "Invalid request", err.Error())
	}

	user, err := c.service.Login(body.Email, body.Password)

	if err != nil {
		return utils.Anauthorized(ctx, "Login failed", err.Error())
	}

	token, _ := utils.GenerateToken(user.InternalID, user.Role, user.Email, user.PublicID)

	refreshToken, _ := utils.GenerateRefreshToken(user.InternalID)

	var userResp models.UserResponse
	_ = copier.Copy(&userResp,&user)

	return utils.Success(ctx, "Login successful", fiber.Map{
		"user":         userResp,
		"token":        token,
		"refreshToken": refreshToken,
	})
}