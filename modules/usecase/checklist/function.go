package checklist

import (
	"time"

	vld "github.com/fazaalexander/BTSid-Case-Study.git/helper/validator"
	ce "github.com/fazaalexander/BTSid-Case-Study.git/modules/entity"
)

func (cc *checklistUseCase) CreateChecklist(checklistRequest ce.ChecklistRequest) error {
	if err := vld.Validation(checklistRequest); err != nil {
		return err
	}

	checklist := &ce.Checklist{
		Title:       checklistRequest.Title,
		Description: checklistRequest.Description,
		CreatedAt:   time.Now(),
	}

	return cc.checklistRepo.CreateChecklist(checklist)
}
func (cc *checklistUseCase) DeleteChecklist(checklistId uint64, checklist ce.Checklist) error {
	return cc.checklistRepo.DeleteChecklist(checklistId, checklist)
}

func (cc *checklistUseCase) GetAllChecklists(checklists *[]ce.Checklist, offset, pageSize int) ([]ce.Checklist, int64, error) {
	return cc.checklistRepo.GetAllChecklists(checklists, offset, pageSize)
}

func (cc *checklistUseCase) GetAllChecklistItems(checklistId uint64, checklist *ce.Checklist) (*ce.Checklist, error) {
	return cc.checklistRepo.GetAllChecklistItems(checklistId, checklist)
}

func (cc *checklistUseCase) CreateChecklistItem(checklistId uint64, request ce.CreateChecklistDetailRequest) error {
	if err := vld.Validation(request); err != nil {
		return err
	}

	checklistDetail := &ce.ChecklistDetail{
		ChecklistID: checklistId,
		Title:       request.Title,
		Description: request.Description,
		CreatedAt:   time.Now(),
	}

	return cc.checklistRepo.CreateChecklistItem(*checklistDetail)
}

func (cc *checklistUseCase) GetChecklistItemDetail(checklistDetailId uint64, checklistDetail *ce.ChecklistDetail) (*ce.ChecklistDetail, error) {
	return cc.checklistRepo.GetChecklistItemDetail(checklistDetailId, checklistDetail)
}

func (cc *checklistUseCase) UpdateChecklistItem(checklistDetailId uint64, request ce.CreateChecklistDetailRequest) error {
	return cc.checklistRepo.UpdateChecklistItem(checklistDetailId, request)
}

func (cc *checklistUseCase) UpdateItemStatus(checklistDetailId uint64, checklistDetail *ce.ChecklistDetail) error {
	return cc.checklistRepo.UpdateItemStatus(checklistDetailId, checklistDetail)
}

func (cc *checklistUseCase) DeleteChecklistItem(checklistDetailId uint64, checklistDetail *ce.ChecklistDetail) error {
	return cc.checklistRepo.DeleteChecklistItem(checklistDetailId, checklistDetail)
}
