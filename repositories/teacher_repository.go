package repositories

import (
	"strings"

	"github.com/christ-saragih/tk-imanuel-bogor-be/config"
	"github.com/christ-saragih/tk-imanuel-bogor-be/models"
)

type TeacherRepository interface {
	Create(teacher *models.Teacher) error
	FindBySlug(slug string) (*models.Teacher, error)
	FindAll(filter, sort string, limit, offset int) ([]models.Teacher, int64, error)
	FindByPublicID(publicID string) (*models.Teacher, error)
	Update(teacher *models.Teacher) error
	Delete(teacher *models.Teacher) error
}

type teacherRepository struct{}

func NewTeacherRepository() TeacherRepository {
	return &teacherRepository{}
}

func (r *teacherRepository) Create(teacher *models.Teacher) error {
	return config.DB.Create(teacher).Error
}

func (r *teacherRepository) FindBySlug(slug string) (*models.Teacher, error) {
	var teacher models.Teacher
	err := config.DB.Where("slug = ?", slug).First(&teacher).Error
	return &teacher, err
}

func (r *teacherRepository) FindAll(filter, sort string, limit, offset int) ([]models.Teacher, int64, error) {
	var teachers []models.Teacher
	var total int64

	db := config.DB.Model(&models.Teacher{})

	// filtering
	if filter != "" {
		filterPattern := "%" + filter + "%"
		db = db.Where("name ILIKE ? OR role ILIKE ?", filterPattern, filterPattern)
	}

	// get total count
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// sorting
	if sort != "" {
		if sort == "-id" {
			sort = "-internal_id"
		} else if sort == "id" {
			sort = "internal_id"
		}

		if strings.HasPrefix(sort, "-") {
			sort = strings.TrimPrefix(sort, "-") + " DESC"
		} else {
			sort += " ASC"
		}

		db = db.Order(sort)
	}

	err := db.Limit(limit).Offset(offset).Find(&teachers).Error
	return teachers, total, err
}

func (r *teacherRepository) FindByPublicID(publicID string) (*models.Teacher, error) {
	var teacher models.Teacher
	err := config.DB.Where("public_id = ?", publicID).First(&teacher).Error
	return &teacher, err
}

func (r *teacherRepository) Update(teacher *models.Teacher) error {
	return config.DB.Save(teacher).Error
}

func (r *teacherRepository) Delete(teacher *models.Teacher) error {
	return config.DB.Delete(teacher).Error
}