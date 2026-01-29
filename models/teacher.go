package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Teacher struct {
	InternalID int64          `json:"internal_id" db:"internal_id" gorm:"primaryKey;autoIncrement"`
	PublicID   uuid.UUID      `json:"public_id" db:"public_id" gorm:"type:uuid;default:gen_random_uuid()"`
	Name       string         `json:"name" db:"name"`
	Slug       string         `json:"slug" db:"slug" gorm:"unique"`
	Role       string         `json:"role" db:"role"`
	Photo      string         `json:"photo" db:"photo"`
	Color      string         `json:"color" db:"color"`
	Bio        string         `json:"bio" db:"bio"`
	Education  string         `json:"education" db:"education"`
	Experience int            `json:"experience" db:"experience"`
	FunFact    string         `json:"fun_fact" db:"fun_fact"`
	Quote      string         `json:"quote" db:"quote"`
	CreatedAt  time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at" db:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index"`
}

type TeacherListResponse struct {
	PublicID   uuid.UUID `json:"public_id"`
	Name  string `json:"name"`
    Slug  string `json:"slug"`
    Role  string `json:"role"`
    Photo string `json:"photo"`
    Color string `json:"color"`
	Bio  string `json:"bio"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"-"`
}