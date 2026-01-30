package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Blog struct {
	InternalID int64     `json:"internal_id" db:"internal_id" gorm:"primaryKey;autoIncrement"`
	PublicID   uuid.UUID `json:"public_id" db:"public_id" gorm:"type:uuid;default:gen_random_uuid()"`

	Slug    string `json:"slug" db:"slug" gorm:"unique"`
	Title   string `json:"title" db:"title" gorm:"unique"`
	Excerpt string `json:"excerpt" db:"excerpt"`
	Content string `json:"content" db:"content"`
	Image   string `json:"image" db:"image"`

	Tags pq.StringArray `json:"tags" db:"tags" gorm:"type:text[]"`

	ViewCount int `json:"view_count" db:"view_count" gorm:"default:0"`

	CreatedAt time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt time.Time      `json:"updated_at" db:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}


type BlogListResponse struct {
	PublicID   uuid.UUID `json:"public_id"`
	Slug    string `json:"slug"`
	Title   string `json:"title"`
	Excerpt string `json:"excerpt"`
	Image   string `json:"image"`
	Tags      []string `json:"tags"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"-"`
}