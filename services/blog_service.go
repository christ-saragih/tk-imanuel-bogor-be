package services

import (
	"fmt"
	"os"

	"github.com/christ-saragih/tk-imanuel-bogor-be/models"
	"github.com/christ-saragih/tk-imanuel-bogor-be/repositories"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
)

type BlogService interface {
	Create(blog *models.Blog) error
	GetAll(filter, sort string, limit, offset int) ([]models.Blog, int64, error)
	GetBySlug(slug string) (*models.Blog, error)
	Update(slug string, blogData *models.Blog) (*models.Blog, error)
	Delete(slug string) error
	RecordView(slug string) error
}

type blogService struct {
	repo repositories.BlogRepository
}

func NewBlogService(repo repositories.BlogRepository) BlogService {
	return &blogService{repo}
}

func (s *blogService) Create(blog *models.Blog) error {
	blog.PublicID = uuid.New()
	blog.Slug = slug.Make(blog.Title)
	
	checkBlog, err := s.repo.FindBySlug(blog.Slug)

	if err == nil && checkBlog.InternalID != 0 {
		return fmt.Errorf("judul artikel '%s' sudah digunakan", blog.Title)
	}

	return s.repo.Create(blog)
}

func (s *blogService) GetAll(filter, sort string, limit, offset int) ([]models.Blog, int64, error) {
	return s.repo.FindAll(filter, sort, limit, offset)
}

func (s *blogService) GetBySlug(slug string) (*models.Blog, error) {
	return s.repo.FindBySlug(slug)
}

func (s *blogService) Update(slugParam string, blogData *models.Blog) (*models.Blog, error) {
	existing, err := s.repo.FindBySlug(slugParam)
	if err != nil {
		return nil, err
	}

	// Update Fields
	if blogData.Title != "" && blogData.Title != existing.Title {
		newSlug := slug.Make(blogData.Title)

		checkBlog, err := s.repo.FindBySlug(newSlug)

		if err == nil && checkBlog.InternalID != existing.InternalID {
            return nil, fmt.Errorf("judul artikel '%s' sudah digunakan", blogData.Title)
        }

		existing.Title = blogData.Title
		existing.Slug = newSlug
	}

	if blogData.Excerpt != "" { existing.Excerpt = blogData.Excerpt }
	if blogData.Content != "" { existing.Content = blogData.Content }
	
	// jika field tags dikirim di request, maka timpa yang lama.
	if len(blogData.Tags) > 0 {
		existing.Tags = blogData.Tags
	}

	if blogData.Image != "" {
		if existing.Image != "" {
			_ = os.Remove(existing.Image)
		}
		existing.Image = blogData.Image
	}

	if err := s.repo.Update(existing); err != nil {
		return nil, err
	}
	return existing, nil
}

func (s *blogService) Delete(slugParam string) error {
	existing, err := s.repo.FindBySlug(slugParam)
	if err != nil {
		return err
	}
	return s.repo.Delete(existing)
}

func (s *blogService) RecordView(slugParam string) error {
	return s.repo.IncrementViewCount(slugParam)
}
