package checklist

import (
	ce "github.com/fazaalexander/BTSid-Case-Study.git/modules/entity"
	cr "github.com/fazaalexander/BTSid-Case-Study.git/modules/repository/checklist"
)

type ChecklistUseCase interface {
	CreateChecklist(checklistRequest ce.ChecklistRequest) error
	DeleteChecklist(checklistId uint64, checklist ce.Checklist) error
	GetAllChecklists(checklists *[]ce.Checklist, offset int, pageSize int) ([]ce.Checklist, int64, error)
	GetAllChecklistItems(checklistId uint64, checklist *ce.Checklist) (*ce.Checklist, error)
	CreateChecklistItem(checklistId uint64, request ce.CreateChecklistDetailRequest) error
	GetChecklistItemDetail(checklistDetailId uint64, checklistDetail *ce.ChecklistDetail) (*ce.ChecklistDetail, error)
	UpdateChecklistItem(checklistDetailId uint64, request ce.CreateChecklistDetailRequest) error
	UpdateItemStatus(checklistDetailId uint64, checklistDetail *ce.ChecklistDetail) error
	DeleteChecklistItem(checklistDetailId uint64, checklistDetail *ce.ChecklistDetail) error
}

type checklistUseCase struct {
	checklistRepo cr.ChecklistRepo
}

func New(checklistRepo cr.ChecklistRepo) *checklistUseCase {
	return &checklistUseCase{
		checklistRepo,
	}
}
