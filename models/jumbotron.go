package models

import (
	"time"

	"github.com/google/uuid"
)

type Jumbotron struct {
	InternalID  int64     		`json:"internal_id" db:"internal_id" gorm:"primaryKey;autoIncrement"`
	PublicID    uuid.UUID 		`json:"public_id" db:"public_id" gorm:"type:uuid;default:gen_random_uuid()"`
	Title       string    		`json:"title" db:"title"`
	Description string    		`json:"description" db:"description"`
	Image       string    		`json:"image" db:"image"`
	LastUpdated  time.Time      `json:"last_updated" db:"last_updated"`
}