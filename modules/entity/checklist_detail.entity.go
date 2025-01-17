package entity

import (
	"time"

	"gorm.io/gorm"
)

type ChecklistDetail struct {
	*gorm.Model `json:"-"`
	ID          uint64      `gorm:"primarykey" json:"ItemId"`
	ChecklistID uint64      `gorm:"not null" json:"ChecklistId"`
	Title       string      `gorm:"not null" validate:"required" json:"title"`
	Description string      `json:"description"`
	IsCompleted bool        `gorm:"default:false" json:"is_completed"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
	Checklists  []Checklist `gorm:"foreignKey:UserID" json:"checklists,omitempty"`
}

type CreateChecklistDetailRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
}

type UpdateChecklistDetailRequest struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	IsCompleted *bool  `json:"is_completed,omitempty"`
}
