package services

import (
	"fmt"
	"os"

	"github.com/christ-saragih/tk-imanuel-bogor-be/models"
	"github.com/christ-saragih/tk-imanuel-bogor-be/repositories"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type TeacherService interface {
	Create(teacher *models.Teacher) error
	GetAll(filter, sort string, limit, offset int) ([]models.Teacher, int64, error)
	GetByPublicID(publicID string) (*models.Teacher, error)
	Update(publicID string, teacher *models.Teacher) (*models.Teacher, error)
	Delete(publicID string) error
}

type teacherService struct {
	repo repositories.TeacherRepository
}

func NewTeacherService(repo repositories.TeacherRepository) TeacherService {
	return &teacherService{repo}
}

func (s *teacherService) Create(teacher *models.Teacher) error {
	teacher.PublicID = uuid.New()
	teacher.Slug = s.generateUniqueSlug(teacher.Name)

	return s.repo.Create(teacher)
}

func (s *teacherService) generateUniqueSlug(name string) string {
	baseSlug := slug.Make(name)
	uniqueSlug := baseSlug
	counter := 1

	for {
		_, err := s.repo.FindBySlug(uniqueSlug)

		if err != nil && err == gorm.ErrRecordNotFound {
            break
        }

		uniqueSlug = fmt.Sprintf("%s-%d", baseSlug, counter)
        counter++
	}

	return uniqueSlug
}

func (s *teacherService) GetAll(filter, sort string, limit, offset int) ([]models.Teacher, int64, error) {
	return s.repo.FindAll(filter, sort, limit, offset)
}

func (s *teacherService) GetByPublicID(publicID string) (*models.Teacher, error) {
	return s.repo.FindByPublicID(publicID)
}

func (s *teacherService) Update(publicID string, teacher *models.Teacher) (*models.Teacher, error) {
	existingTeacher, err := s.repo.FindByPublicID(publicID)
	if err != nil {
		return nil, err
	}

	if teacher.Name != "" && teacher.Name != existingTeacher.Name {
        existingTeacher.Name = teacher.Name
        existingTeacher.Slug = s.generateUniqueSlug(teacher.Name)
    }

	if teacher.Role != "" { existingTeacher.Role = teacher.Role }
    if teacher.Bio != "" { existingTeacher.Bio = teacher.Bio }
    if teacher.Education != "" { existingTeacher.Education = teacher.Education }
    if teacher.Experience != 0 { existingTeacher.Experience = teacher.Experience }
    if teacher.FunFact != "" { existingTeacher.FunFact = teacher.FunFact }
    if teacher.Quote != "" { existingTeacher.Quote = teacher.Quote }
    if teacher.Color != "" { existingTeacher.Color = teacher.Color }

	if teacher.Photo != "" {
        
		if existingTeacher.Photo != "" {
			if _, err := os.Stat(existingTeacher.Photo); err == nil {
				_ = os.Remove(existingTeacher.Photo)
			}
		}
        existingTeacher.Photo = teacher.Photo
    }

	if err := s.repo.Update(existingTeacher); err != nil {
        return nil, err
    }

    return existingTeacher, nil

}

func (s *teacherService) Delete(publicID string) error {
	teacher, err := s.repo.FindByPublicID(publicID)
	if err != nil {
		return err
	}

	return s.repo.Delete(teacher)
}