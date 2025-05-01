package model

type OperationManualOutput struct {
	OperaManualID      string `json:"opera_manual_id" binding:"required"`
	OperaManualTitle   string `json:"opera_manual_title" binding:"required"`
	OperaManualContent string `json:"opera_manual_content" binding:"required"`
	OperaManualType    string `json:"opera_manual_type" binding:"required"`
	OperaManualNote    string `json:"opera_manual_note" binding:"required"`
	OperaManualStatus  string `json:"opera_manual_status" binding:"required"`
	Isdeleted          int32  `json:"isDeleted"`
}

type CreateOperationManualDto struct {
	OperaManualTitle   string `json:"opera_manual_title" binding:"required"`
	OperaManualContent string `json:"opera_manual_content" binding:"required"`
	OperaManualType    string `json:"opera_manual_type" binding:"required"`
	OperaManualNote    string `json:"opera_manual_note"`
}

type UpdateOperationManualDto struct {
	OperaManualID      string `json:"opera_manual_id" binding:"required"`
	OperaManualTitle   string `json:"opera_manual_title" binding:"required"`
	OperaManualContent string `json:"opera_manual_content" binding:"required"`
	OperaManualType    string `json:"opera_manual_type" binding:"required"`
	OperaManualNote    string `json:"opera_manual_note" binding:"required"`
}

type UpdateOperationManualStatusDto struct {
	OperaManualID     string `json:"opera_manual_id" binding:"required"`
	OperaManualStatus string `json:"opera_manual_status" binding:"required"`
}