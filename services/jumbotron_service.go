package services

import (
	"os"
	"time"

	"github.com/christ-saragih/tk-imanuel-bogor-be/models"
	"github.com/christ-saragih/tk-imanuel-bogor-be/repositories"
	"github.com/google/uuid"
)

type JumbotronService interface {
	Get() (*models.Jumbotron, error)
	Upsert(data *models.Jumbotron) (*models.Jumbotron, error)
}

type jumbotronService struct {
	repo repositories.JumbotronRepository
}

func NewJumbotronService(repo repositories.JumbotronRepository) JumbotronService {
	return &jumbotronService{repo}
}

func (s *jumbotronService) Get() (*models.Jumbotron, error) {
	return s.repo.Get()
}

func (s *jumbotronService) Upsert(jumbotron *models.Jumbotron) (*models.Jumbotron, error) {
	existing, err := s.repo.Get()

	if err == nil && existing != nil && existing.InternalID != 0 {
		// update
		jumbotron.InternalID = existing.InternalID
		jumbotron.PublicID = existing.PublicID

		if jumbotron.Image != "" {
			if existing.Image != "" {
				_ = os.Remove(existing.Image)
			}
		} else {
			jumbotron.Image = existing.Image
		}
	} else {
		// insert
		jumbotron.PublicID = uuid.New()
	}

	jumbotron.LastUpdated = time.Now()

	if err := s.repo.Save(jumbotron); err != nil {
		return nil, err
	}

	return jumbotron, nil
}