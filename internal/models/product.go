package models

import "time"

type Product struct {
	ID           int     `json:"id" gorm:"primaryKey;autoIncrement"`
	GameID       int     `json:"game_id" gorm:"type:integer;not null;index"`
	Name         string  `json:"name" gorm:"type:varchar(255);not null"`
	Price        float64 `json:"price" gorm:"type:decimal(10,2);not null"`
	ProviderCode string  `json:"provider_code" gorm:"type:varchar(255)"`

	IsActive  bool      `json:"is_active" gorm:"type:boolean;default:true"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
}
