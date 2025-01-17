package checklist

import (
	ce "github.com/fazaalexander/BTSid-Case-Study.git/modules/entity"
	"github.com/labstack/echo/v4"
)

func (cr *checklistRepo) CreateChecklist(checklist *ce.Checklist) error {
	if err := cr.db.Create(&checklist).Error; err != nil {
		return echo.NewHTTPError(500, err)
	}

	return nil
}

func (cr *checklistRepo) DeleteChecklist(checklistId uint64, checklist ce.Checklist) error {
	var count int64
	if err := cr.db.Model(&ce.ChecklistDetail{}).Where("checklist_id = ?", checklistId).Count(&count).Error; err == nil {
		return echo.NewHTTPError(500, err)
	}

	if err := cr.db.Where("id = ?", checklistId).Delete(&checklist).Error; err != nil {
		return echo.NewHTTPError(500, err)
	}

	return nil
}

func (cr *checklistRepo) GetAllChecklists(checklists *[]ce.Checklist, offset, pageSize int) ([]ce.Checklist, int64, error) {
	var count int64
	if err := cr.db.Model(&ce.Checklist{}).Count(&count).Error; err != nil {
		return nil, 0, echo.NewHTTPError(500, err)
	}

	if err := cr.db.
		Offset(offset).
		Limit(pageSize).Order("created_at ASC").
		Find(&checklists).Error; err != nil {
		return nil, 0, echo.NewHTTPError(404, err)
	}

	return *checklists, count, nil
}

func (cr *checklistRepo) GetAllChecklistItems(checklistId uint64, checklist *ce.Checklist) (*ce.Checklist, error) {
	var count int64
	if err := cr.db.Model(&ce.ChecklistDetail{}).Where("checklist_id = ?", checklistId).Count(&count).Error; err != nil {
		return nil, echo.NewHTTPError(500, err)
	}

	if err := cr.db.
		Preload("ChecklistDetail").
		First(&checklist).Error; err != nil {
		return nil, echo.NewHTTPError(404, err)
	}

	return checklist, nil
}

func (cr *checklistRepo) CreateChecklistItem(checklistDetail ce.ChecklistDetail) error {
	if err := cr.db.Create(&checklistDetail).Error; err != nil {
		return echo.NewHTTPError(500, err)
	}

	return nil
}

func (cr *checklistRepo) GetChecklistItemDetail(checklistDetailId uint64, checklistDetail *ce.ChecklistDetail) (*ce.ChecklistDetail, error) {
	if err := cr.db.Where("id = ?", checklistDetailId).First(&checklistDetail).Error; err != nil {
		return nil, echo.NewHTTPError(404, err)
	}

	return nil, nil
}

func (cr *checklistRepo) UpdateChecklistItem(checklistDetailId uint64, request ce.CreateChecklistDetailRequest) error {
	if err := cr.db.Where("id = ?", checklistDetailId).First(&ce.ChecklistDetail{}).Error; err != nil {
		return echo.NewHTTPError(404, err)
	}

	if err := cr.db.Model(&ce.ChecklistDetail{}).Where("id = ?", checklistDetailId).Updates(ce.ChecklistDetail{Title: request.Title, Description: request.Description}).Error; err != nil {
		return echo.NewHTTPError(500, err)
	}

	return nil
}

func (cr *checklistRepo) UpdateItemStatus(checklistDetailId uint64, checklistDetail *ce.ChecklistDetail) error {
	if err := cr.db.Where("id = ?", checklistDetailId).First(&ce.ChecklistDetail{}).Error; err != nil {
		return echo.NewHTTPError(404, err)
	}

	if err := cr.db.Model(&ce.ChecklistDetail{}).Where("id = ?", checklistDetailId).Updates(ce.ChecklistDetail{IsCompleted: true}).Error; err != nil {
		return echo.NewHTTPError(500, err)
	}

	return nil
}

func (cr *checklistRepo) DeleteChecklistItem(checklistDetailId uint64, checklistDetail *ce.ChecklistDetail) error {
	if err := cr.db.Where("id = ?", checklistDetailId).Delete(checklistDetail).Error; err != nil {
		return echo.NewHTTPError(404, err)
	}

	return nil
}
