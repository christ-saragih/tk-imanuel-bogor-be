package services

import (
	"time"

	"github.com/christ-saragih/tk-imanuel-bogor-be/models"
	"github.com/christ-saragih/tk-imanuel-bogor-be/repositories"
	"github.com/google/uuid"
)

type ContactService interface {
	Get() (*models.Contact, error)
	Upsert(data *models.Contact) (*models.Contact, error)
}

type contactService struct {
	repo repositories.ContactRepository
}

func NewContactService(repo repositories.ContactRepository) ContactService {
	return &contactService{repo}
}

func (s *contactService) Get() (*models.Contact, error) {
	return s.repo.Get()
}

func (s *contactService) Upsert(contact *models.Contact) (*models.Contact, error) {
	existing, err := s.repo.Get()

	if err == nil && existing != nil && existing.InternalID != 0 {
		contact.InternalID = existing.InternalID
		contact.PublicID = existing.PublicID
	} else {
		contact.PublicID = uuid.New()
	}

	contact.LastUpdated = time.Now()

	if err := s.repo.Save(contact); err != nil {
		return nil, err
	}

	return contact, nil
}