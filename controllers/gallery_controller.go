package controllers

import (
	"math"
	"strconv"

	"github.com/christ-saragih/tk-imanuel-bogor-be/models"
	"github.com/christ-saragih/tk-imanuel-bogor-be/services"
	"github.com/christ-saragih/tk-imanuel-bogor-be/utils"
	"github.com/gofiber/fiber/v2"
)

type GalleryController struct {
	service services.GalleryService
}

func NewGalleryController(service services.GalleryService) *GalleryController {
	return &GalleryController{service}
}

func (c *GalleryController) CreateGallery(ctx *fiber.Ctx) error {
	var gallery models.Gallery

	gallery.Title = ctx.FormValue("title")
	
	// Image Wajib
	photoPath, err := utils.UploadFile(ctx, "image", "galleries")
	if err != nil {
		return utils.BadRequest(ctx, "Upload image required", err.Error())
	}
	gallery.Image = photoPath

	if err := c.service.Create(&gallery); err != nil {
		return utils.BadRequest(ctx, "Failed to create gallery", err.Error())
	}

	return utils.Success(ctx, "Success create gallery", gallery)
}

func (c *GalleryController) GetGalleries(ctx *fiber.Ctx) error {
	page, _ := strconv.Atoi(ctx.Query("page", "1"))
	limit, _ := strconv.Atoi(ctx.Query("limit", "12")) // Default limit lebih banyak untuk galeri
	offset := (page - 1) * limit

	galleries, total, err := c.service.GetAll(limit, offset)
	if err != nil {
		return utils.BadRequest(ctx, "Failed to fetch galleries", err.Error())
	}

	meta := utils.PaginationMeta{
		Page:      page,
		Limit:     limit,
		Total:     int(total),
		TotalPage: int(math.Ceil(float64(total) / float64(limit))),
	}

	return utils.SuccessPagination(ctx, "Galleries retrieved", galleries, meta)
}

func (c *GalleryController) UpdateGallery(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var gallery models.Gallery

	gallery.Title = ctx.FormValue("title")

	// Image Optional saat update
	if _, err := ctx.FormFile("image"); err == nil {
		photoPath, err := utils.UploadFile(ctx, "image", "galleries")
		if err != nil {
			return utils.BadRequest(ctx, "Failed to upload image", err.Error())
		}
		gallery.Image = photoPath
	}

	updated, err := c.service.Update(id, &gallery)
	if err != nil {
		return utils.BadRequest(ctx, "Failed to update gallery", err.Error())
	}

	return utils.Success(ctx, "Success update gallery", updated)
}

func (c *GalleryController) DeleteGallery(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if err := c.service.Delete(id); err != nil {
		return utils.BadRequest(ctx, "Failed to delete gallery", err.Error())
	}
	return utils.Success(ctx, "Success delete gallery", nil)
}
