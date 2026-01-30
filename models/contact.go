package models

import (
    "time"

    "github.com/google/uuid"
)

type Contact struct {
    InternalID     int64     `json:"internal_id" db:"internal_id" gorm:"primaryKey;autoIncrement"`
    PublicID       uuid.UUID `json:"public_id" db:"public_id" gorm:"type:uuid;default:gen_random_uuid()"`
    
    Address        string    `json:"address" db:"address"`
    MapsEmbedUrl   string    `json:"maps_embed_url" db:"maps_embed_url"`
    MapsLink       string    `json:"maps_link" db:"maps_link"`
    
    Email          string    `json:"email" db:"email"`
    PhoneNumber    string    `json:"phone_number" db:"phone_number"`
    WhatsappNumber string    `json:"whatsapp_number" db:"whatsapp_number"`
    
    InstagramUrl   string    `json:"instagram_url" db:"instagram_url"`
    FacebookUrl    string    `json:"facebook_url" db:"facebook_url"`
    YoutubeUrl     string    `json:"youtube_url" db:"youtube_url"`
    TiktokUrl      string    `json:"tiktok_url" db:"tiktok_url"`
    
    OpeningHours   string    `json:"opening_hours" db:"opening_hours"`
    
    LastUpdated    time.Time `json:"last_updated" db:"last_updated"`
}