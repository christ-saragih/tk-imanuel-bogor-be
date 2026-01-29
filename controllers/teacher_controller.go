package controllers

import (
	"math"
	"strconv"

	"github.com/christ-saragih/tk-imanuel-bogor-be/models"
	"github.com/christ-saragih/tk-imanuel-bogor-be/services"
	"github.com/christ-saragih/tk-imanuel-bogor-be/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
)

type TeacherController struct {
	services services.TeacherService
}

func NewTeacherController(services services.TeacherService) *TeacherController {
	return &TeacherController{services}
}

func (c *TeacherController) CreateTeacher(ctx *fiber.Ctx) error {
	var teacher models.Teacher

    teacher.Name = ctx.FormValue("name")
    teacher.Role = ctx.FormValue("role")
    teacher.Bio = ctx.FormValue("bio")
    teacher.Education = ctx.FormValue("education")
    teacher.FunFact = ctx.FormValue("fun_fact")
    teacher.Quote = ctx.FormValue("quote")
    teacher.Color = ctx.FormValue("color")
    
    if exp, err := strconv.Atoi(ctx.FormValue("experience")); err == nil {
        teacher.Experience = exp
    }

	photoPath, err := utils.UploadFile(ctx, "photo", "teachers")
	if err != nil {
		return utils.BadRequest(ctx, "Failed to upload photo", err.Error())
	}
	teacher.Photo = photoPath

	if err := c.services.Create(&teacher); err != nil {
		return utils.BadRequest(ctx, "Failed to create teacher", err.Error())
	}

	return utils.Success(ctx, "Success create teacher", teacher)
}

func (c *TeacherController) GetTeachers(ctx *fiber.Ctx) error {
	page, _ := strconv.Atoi(ctx.Query("page", "1"))
	limit, _ := strconv.Atoi(ctx.Query("limit", "10"))
	offset := (page - 1) * limit

	filter := ctx.Query("filter", "")
	sort := ctx.Query("sort", "")

	teachers, total, err := c.services.GetAll(filter, sort, limit, offset)
	if err != nil {
		return utils.BadRequest(ctx, "Failed to get teachers", err.Error())
	}

	var teacherList []models.TeacherListResponse
	_ = copier.Copy(&teacherList, teachers)

	meta := utils.PaginationMeta{
		Page:       page,
		Limit:      limit,
		Total:      int(total),
		TotalPage: int(math.Ceil(float64(total)/float64(limit))),
		Filter: filter,
		Sort: sort,
	}

	if total == 0 {
		return utils.NotFoundPagination(ctx, "No teachers found", teacherList, meta)
	}

	return utils.SuccessPagination(ctx, "Teacher retrieved successfully", teacherList, meta)
}

func (c *TeacherController) GetTeacherDetail(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	teacher, err := c.services.GetByPublicID(id)

	if err != nil {
		return utils.NotFound(ctx, "Teacher not found", err.Error())
	}

	return utils.Success(ctx, "Teacher retrieved successfully", teacher)
}

func (c *TeacherController) UpdateTeacher(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
    var teacher models.Teacher

    teacher.Name = ctx.FormValue("name")
    teacher.Role = ctx.FormValue("role")
    teacher.Bio = ctx.FormValue("bio")
    teacher.Education = ctx.FormValue("education")
    teacher.FunFact = ctx.FormValue("fun_fact")
    teacher.Quote = ctx.FormValue("quote")
    teacher.Color = ctx.FormValue("color")
    
    if expStr := ctx.FormValue("experience"); expStr != "" {
        if exp, err := strconv.Atoi(expStr); err == nil {
            teacher.Experience = exp
        }
    }

	if _, err := ctx.FormFile("photo"); err == nil {
        photoPath, err := utils.UploadFile(ctx, "photo", "teachers")
        if err != nil {
            return utils.BadRequest(ctx, "Failed to upload photo", err.Error())
        }
        teacher.Photo = photoPath
    }

	updatedTeacher, err := c.services.Update(id, &teacher)
    if err != nil {
        return utils.BadRequest(ctx, "Failed to update teacher", err.Error())
    }

    return utils.Success(ctx, "Success update teacher", updatedTeacher)
}

func (c *TeacherController) DeleteTeacher(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	if err := c.services.Delete(id); err != nil {
		return utils.BadRequest(ctx, "Failed to delete teacher", err.Error())
	}

	return utils.Success(ctx, "Success delete teacher", id)

}