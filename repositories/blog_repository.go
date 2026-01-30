package repositories

import (
	"strings"

	"github.com/christ-saragih/tk-imanuel-bogor-be/config"
	"github.com/christ-saragih/tk-imanuel-bogor-be/models"
	"gorm.io/gorm"
)

type BlogRepository interface {
	Create(blog *models.Blog) error
	FindAll(filter, sort string, limit, offset int) ([]models.Blog, int64, error)
	FindBySlug(slug string) (*models.Blog, error)
	Update(blog *models.Blog) error
	Delete(blog *models.Blog) error
	IncrementViewCount(slug string) error
}

type blogRepository struct{}

func NewBlogRepository() BlogRepository {
	return &blogRepository{}
}

func (r *blogRepository) Create(blog *models.Blog) error {
	return config.DB.Create(blog).Error
}

func (r *blogRepository) FindAll(filter, sort string, limit, offset int) ([]models.Blog, int64, error) {
	var blogs []models.Blog
	var total int64

	db := config.DB.Model(&models.Blog{})

	if filter != "" {
		filterPattern := "%" + filter + "%"
		db = db.Where("title ILIKE ? OR content ILIKE ?", filterPattern, filterPattern)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if sort != "" {
		if strings.HasPrefix(sort, "-") {
			sort = strings.TrimPrefix(sort, "-") + " DESC"
		} else {
			sort += " ASC"
		}
		db = db.Order(sort)
	} else {
		db = db.Order("created_at DESC")
	}

	err := db.Limit(limit).Offset(offset).Find(&blogs).Error
	return blogs, total, err
}

func (r *blogRepository) FindBySlug(slug string) (*models.Blog, error) {
	var blog models.Blog
	err := config.DB.Where("slug = ?", slug).First(&blog).Error
	return &blog, err
}

func (r *blogRepository) Update(blog *models.Blog) error {
	return config.DB.Save(blog).Error
}

func (r *blogRepository) Delete(blog *models.Blog) error {
	return config.DB.Delete(blog).Error
}

func (r *blogRepository) IncrementViewCount(slug string) error {
	return config.DB.Model(&models.Blog{}).Where("slug = ?", slug).UpdateColumn("view_count", gorm.Expr("view_count + ?", 1)).Error
}
