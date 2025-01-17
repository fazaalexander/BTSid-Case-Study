package entity

import (
	"time"

	"gorm.io/gorm"
)

type Checklist struct {
	*gorm.Model      `json:"-"`
	ID               uint64            `gorm:"primarykey" json:"ChecklistId"`
	UserID           uint64            `gorm:"not null" json:"UserId"`
	Title            string            `gorm:"not null" validate:"required" json:"title"`
	Description      string            `json:"description"`
	CreatedAt        time.Time         `json:"created_at"`
	UpdatedAt        time.Time         `json:"updated_at"`
	ChecklistDetails []ChecklistDetail `gorm:"foreignKey:ChecklistID" json:"checklist_detail,omitempty"`
}

type ChecklistRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
}
