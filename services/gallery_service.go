package services

import (
	"os"

	"github.com/christ-saragih/tk-imanuel-bogor-be/models"
	"github.com/christ-saragih/tk-imanuel-bogor-be/repositories"
	"github.com/google/uuid"
)

type GalleryService interface {
	Create(gallery *models.Gallery) error
	GetAll(limit, offset int) ([]models.Gallery, int64, error)
	Update(publicID string, galleryData *models.Gallery) (*models.Gallery, error)
	Delete(publicID string) error
}

type galleryService struct {
	repo repositories.GalleryRepository
}

func NewGalleryService(repo repositories.GalleryRepository) GalleryService {
	return &galleryService{repo}
}

func (s *galleryService) Create(gallery *models.Gallery) error {
	gallery.PublicID = uuid.New()
	return s.repo.Create(gallery)
}

func (s *galleryService) GetAll(limit, offset int) ([]models.Gallery, int64, error) {
	return s.repo.FindAll(limit, offset)
}

func (s *galleryService) Update(publicID string, gallery *models.Gallery) (*models.Gallery, error) {
	existing, err := s.repo.FindByPublicID(publicID)
	if err != nil {
		return nil, err
	}

	if gallery.Title != "" {
		existing.Title = gallery.Title
	}

	if gallery.Image != "" {
		// Hapus image lama
		if existing.Image != "" {
			_ = os.Remove(existing.Image)
		}
		existing.Image = gallery.Image
	}

	if err := s.repo.Update(existing); err != nil {
		return nil, err
	}

	return existing, nil
}

func (s *galleryService) Delete(publicID string) error {
	gallery, err := s.repo.FindByPublicID(publicID)
	if err != nil {
		return err
	}
	
	if gallery.Image != "" {
		_ = os.Remove(gallery.Image)
	}

	return s.repo.Delete(gallery)
}
