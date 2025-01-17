package checklist

import (
	ce "github.com/fazaalexander/BTSid-Case-Study.git/modules/entity"
	"gorm.io/gorm"
)

type ChecklistRepo interface {
	CreateChecklist(checklist *ce.Checklist) error
	DeleteChecklist(checklistId uint64, checklist ce.Checklist) error
	GetAllChecklists(checklists *[]ce.Checklist, offset int, pageSize int) ([]ce.Checklist, int64, error)
	GetAllChecklistItems(checklistId uint64, checklist *ce.Checklist) (*ce.Checklist, error)
	CreateChecklistItem(checklistDetail ce.ChecklistDetail) error
	GetChecklistItemDetail(checklistDetailId uint64, checklistDetail *ce.ChecklistDetail) (*ce.ChecklistDetail, error)
	UpdateChecklistItem(checklistDetailId uint64, request ce.CreateChecklistDetailRequest) error
	UpdateItemStatus(checklistDetailId uint64, checklistDetail *ce.ChecklistDetail) error
	DeleteChecklistItem(checklistDetailId uint64, checklistDetail *ce.ChecklistDetail) error
}

type checklistRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) ChecklistRepo {
	return &checklistRepo{
		db,
	}
}
