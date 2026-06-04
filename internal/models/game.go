package models

import "time"

type Game struct {
	ID          int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string `json:"name" gorm:"type:varchar(255);not null"`
	Slug        string `json:"slug" gorm:"type:varchar(255);not null"`
	Description string `json:"description" gorm:"type:text;not null"`
	ImageURL    string `json:"image_url" gorm:"type:varchar(255);not null"`

	IsActive  bool      `json:"is_active" gorm:"type:boolean;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp;not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp;not null"`
}
