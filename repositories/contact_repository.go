package repositories

import (
	"github.com/christ-saragih/tk-imanuel-bogor-be/config"
	"github.com/christ-saragih/tk-imanuel-bogor-be/models"
)

type ContactRepository interface {
	Get() (*models.Contact, error)
	Save(contact *models.Contact) error
}

type contactRepository struct{}

func NewContactRepository() ContactRepository {
	return &contactRepository{}
}

func (r *contactRepository) Get() (*models.Contact, error) {
	var contact models.Contact
	err := config.DB.First(&contact).Error
	return &contact, err
}

func (r *contactRepository) Save(contact *models.Contact) error {
	return config.DB.Save(contact).Error
}