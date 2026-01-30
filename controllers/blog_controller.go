package controllers

import (
	"math"
	"strconv"
	"strings"

	"github.com/christ-saragih/tk-imanuel-bogor-be/models"
	"github.com/christ-saragih/tk-imanuel-bogor-be/services"
	"github.com/christ-saragih/tk-imanuel-bogor-be/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
)

type BlogController struct {
	service services.BlogService
}

func NewBlogController(service services.BlogService) *BlogController {
	return &BlogController{service}
}

func (c *BlogController) CreateBlog(ctx *fiber.Ctx) error {
	var blog models.Blog

	blog.Title = ctx.FormValue("title")
	blog.Excerpt = ctx.FormValue("excerpt")
	blog.Content = ctx.FormValue("content")

	// Parse Tags
	tagsStr := ctx.FormValue("tags")
	if tagsStr != "" {
		blog.Tags = strings.Split(tagsStr, ",")
		// Trim spaces
		for i, v := range blog.Tags {
			blog.Tags[i] = strings.TrimSpace(v)
		}
	}

	photoPath, err := utils.UploadFile(ctx, "image", "blogs")
	if err != nil {
		return utils.BadRequest(ctx, "Failed to upload image", err.Error())
	}
	blog.Image = photoPath

	if err := c.service.Create(&blog); err != nil {
		return utils.BadRequest(ctx, "Failed to create blog", err.Error())
	}

	return utils.Success(ctx, "Success create blog", blog)
}

func (c *BlogController) GetBlogs(ctx *fiber.Ctx) error {
	page, _ := strconv.Atoi(ctx.Query("page", "1"))
	limit, _ := strconv.Atoi(ctx.Query("limit", "10"))
	offset := (page - 1) * limit
	filter := ctx.Query("filter", "")
	sort := ctx.Query("sort", "")

	blogs, total, err := c.service.GetAll(filter, sort, limit, offset)
	if err != nil {
		return utils.BadRequest(ctx, "Failed to fetch blogs", err.Error())
	}

	var response []models.BlogListResponse
	_ = copier.Copy(&response, blogs)

	meta := utils.PaginationMeta{
		Page:      page,
		Limit:     limit,
		Total:     int(total),
		TotalPage: int(math.Ceil(float64(total) / float64(limit))),
		Filter:    filter,
		Sort:      sort,
	}

	return utils.SuccessPagination(ctx, "Blogs retrieved", response, meta)
}

func (c *BlogController) GetBlogDetail(ctx *fiber.Ctx) error {
	slug := ctx.Params("slug")
	blog, err := c.service.GetBySlug(slug)
	if err != nil {
		return utils.NotFound(ctx, "Blog not found", err.Error())
	}
	
	// Increment view count secara async
	go c.service.RecordView(slug) 

	return utils.Success(ctx, "Blog detail retrieved", blog)
}

func (c *BlogController) UpdateBlog(ctx *fiber.Ctx) error {
	slugParam := ctx.Params("slug")
	var blog models.Blog

	blog.Title = ctx.FormValue("title")
	blog.Excerpt = ctx.FormValue("excerpt")
	blog.Content = ctx.FormValue("content")

	tagsStr := ctx.FormValue("tags")
	if tagsStr != "" {
		blog.Tags = strings.Split(tagsStr, ",")
	}

	if _, err := ctx.FormFile("image"); err == nil {
		photoPath, err := utils.UploadFile(ctx, "image", "blogs")
		if err != nil {
			return utils.BadRequest(ctx, "Failed to upload image", err.Error())
		}
		blog.Image = photoPath
	}

	updated, err := c.service.Update(slugParam, &blog)
	if err != nil {
		return utils.BadRequest(ctx, "Failed to update blog", err.Error())
	}

	return utils.Success(ctx, "Success update blog", updated)
}

func (c *BlogController) DeleteBlog(ctx *fiber.Ctx) error {
	slug := ctx.Params("slug")
	if err := c.service.Delete(slug); err != nil {
		return utils.BadRequest(ctx, "Failed to delete blog", err.Error())
	}
	return utils.Success(ctx, "Success delete blog", nil)
}
