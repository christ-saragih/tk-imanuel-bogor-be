package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Gallery struct {
	InternalID int64     `json:"internal_id" db:"internal_id" gorm:"primaryKey;autoIncrement"`
	PublicID   uuid.UUID `json:"public_id" db:"public_id" gorm:"type:uuid;default:gen_random_uuid()"`

	Title string `json:"title" db:"title" gorm:"unique"`
	Image   string `json:"image" db:"image"`

	CreatedAt time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt time.Time      `json:"updated_at" db:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}