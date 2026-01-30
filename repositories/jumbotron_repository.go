package repositories

import (
	"github.com/christ-saragih/tk-imanuel-bogor-be/config"
	"github.com/christ-saragih/tk-imanuel-bogor-be/models"
)

type JumbotronRepository interface {
	Get() (*models.Jumbotron, error)
	Save(jumbotron *models.Jumbotron) error
}

type jumbotronRepository struct{}

func NewJumbotronRepository() JumbotronRepository {
	return &jumbotronRepository{}
}

func (r *jumbotronRepository) Get() (*models.Jumbotron, error) {
	var jumbotron models.Jumbotron

	err := config.DB.First(&jumbotron).Error
	return &jumbotron, err
}

func (r *jumbotronRepository) Save(jumbotron *models.Jumbotron) error {
	return config.DB.Save(jumbotron).Error
}