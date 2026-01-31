package repositories

import (
	"github.com/christ-saragih/tk-imanuel-bogor-be/config"
	"github.com/christ-saragih/tk-imanuel-bogor-be/models"
)

type GalleryRepository interface {
	Create(gallery *models.Gallery) error
	FindAll(limit, offset int) ([]models.Gallery, int64, error)
	FindByPublicID(publicID string) (*models.Gallery, error)
	Update(gallery *models.Gallery) error
	Delete(gallery *models.Gallery) error
}

type galleryRepository struct{}

func NewGalleryRepository() GalleryRepository {
	return &galleryRepository{}
}

func (r *galleryRepository) Create(gallery *models.Gallery) error {
	return config.DB.Create(gallery).Error
}

func (r *galleryRepository) FindAll(limit, offset int) ([]models.Gallery, int64, error) {
	var galleries []models.Gallery
	var total int64

	db := config.DB.Model(&models.Gallery{})
	
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Order by terbaru
	err := db.Order("created_at DESC").Limit(limit).Offset(offset).Find(&galleries).Error
	return galleries, total, err
}

func (r *galleryRepository) FindByPublicID(publicID string) (*models.Gallery, error) {
	var gallery models.Gallery
	err := config.DB.Where("public_id = ?", publicID).First(&gallery).Error
	return &gallery, err
}

func (r *galleryRepository) Update(gallery *models.Gallery) error {
	return config.DB.Save(gallery).Error
}

func (r *galleryRepository) Delete(gallery *models.Gallery) error {
	return config.DB.Delete(gallery).Error
}
