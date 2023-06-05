package models

import (
	"time"

	"github.com/google/uuid"
)

type Phone struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
	Phone       string    `gorm:"uniqueIndex;not null" json:"phone,omitempty"`
	Description string    `gorm:"not null" json:"description,omitempty"`
	IsMobile    bool      `gorm:"not null" json:"is_mobile,omitempty"`
	User        uuid.UUID `gorm:"not null" json:"user,omitempty"`
	CreatedAt   time.Time `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt   time.Time `gorm:"not null" json:"updated_at,omitempty"`
}

type CreatePhoneRequest struct {
	Phone       string    `json:"phone" binding:"required"`
	Description string    `json:"description" binding:"required"`
	IsMobile    bool      `json:"is_mobile,omitempty"`
	User        string    `json:"user,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

type UpdatePhone struct {
	Phone       string    `json:"phone,omitempty"`
	Description string    `json:"description,omitempty"`
	IsMobile    bool      `json:"is_mobile,omitempty"`
	User        string    `json:"user,omitempty"`
	CreateAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

type PhoneResponse struct {
	ID         uuid.UUID `json:"id,omitempty"`
	Phone       string    `gorm:"uniqueIndex;not null" json:"phone,omitempty"`
	Description string    `gorm:"not null" json:"description,omitempty"`
	IsMobile    bool      `gorm:"not null" json:"is_mobile,omitempty"`
	CreatedAt   time.Time `gorm:"not null" json:"created_at,omitempty"`
}