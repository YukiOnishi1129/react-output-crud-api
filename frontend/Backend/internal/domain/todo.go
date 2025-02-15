package domain

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	ID        uuid.UUID  `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Title     string     `json:"title"`
	Content   *string    `json:"content"`
	UserID    uuid.UUID  `json:"user_id" gorm:"type:uuid"`
	User      *User      `json:"user" gorm:"foreignKey:UserID"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
} 