package checklist

import (
	cc "github.com/fazaalexander/BTSid-Case-Study.git/modules/usecase/checklist"
)

type ChecklistHandler struct {
	checklistUseCase cc.ChecklistUseCase
}

func New(authUseCase cc.ChecklistUseCase) *ChecklistHandler {
	return &ChecklistHandler{
		authUseCase,
	}
}
