package model

type EquipmentMaintenanceOutput struct {
	EqpMtnID               string  `json:"eqp_mtn_id"`
	EqpMtnName             string  `json:"eqp_mtn_name"`
	EqpMtnLocation         string  `json:"eqp_mtn_location,omitempty"`
	EqpMtnIssueDescription string  `json:"eqp_mtn_issue_description,omitempty"`
	EqpMtnStatus           string  `json:"eqp_mtn_status"`
	EqpMtnReportedBy       string  `json:"eqp_mtn_reported_by,omitempty"`
	EqpMtnPerformedBy      string  `json:"eqp_mtn_performed_by,omitempty"`
	EqpMtnDateReported     string  `json:"eqp_mtn_date_reported"`
	EqpMtnDateFixed        *string `json:"eqp_mtn_date_fixed,omitempty"`
	EqpMtnCost             *int64  `json:"eqp_mtn_cost,omitempty"`
	EqpMtnNote             string  `json:"eqp_mtn_note,omitempty"`
}

type CreateEquipmentMaintenanceDto struct {
	EqpMtnName             string  `json:"eqp_mtn_name" binding:"required"`
	EqpMtnLocation         string  `json:"eqp_mtn_location,omitempty"`
	EqpMtnIssueDescription string  `json:"eqp_mtn_issue_description,omitempty"`
	EqpMtnReportedBy       string  `json:"eqp_mtn_reported_by,omitempty"`
	EqpMtnPerformedBy      string  `json:"eqp_mtn_performed_by,omitempty"`
	EqpMtnDateReported     string  `json:"eqp_mtn_date_reported" binding:"required"`
	EqpMtnDateFixed        *string `json:"eqp_mtn_date_fixed,omitempty"`
	EqpMtnCost             *int64  `json:"eqp_mtn_cost,omitempty"`
	EqpMtnNote             string  `json:"eqp_mtn_note,omitempty"`
}

type UpdateEquipmentMaintenanceDto struct {
	EqpMtnID               string  `json:"eqp_mtn_id" binding:"required"`
	EqpMtnName             string  `json:"eqp_mtn_name" binding:"required"`
	EqpMtnLocation         string  `json:"eqp_mtn_location,omitempty"`
	EqpMtnIssueDescription string  `json:"eqp_mtn_issue_description,omitempty"`
	EqpMtnReportedBy       string  `json:"eqp_mtn_reported_by,omitempty"`
	EqpMtnPerformedBy      string  `json:"eqp_mtn_performed_by,omitempty"`
	EqpMtnDateReported     string  `json:"eqp_mtn_date_reported" binding:"required"`
	EqpMtnDateFixed        *string `json:"eqp_mtn_date_fixed,omitempty"`
	EqpMtnCost             *int64  `json:"eqp_mtn_cost,omitempty"`
	EqpMtnNote             string  `json:"eqp_mtn_note,omitempty"`
}

type UpdateEquipmentMaintenanceStatusDto struct {
	EqpMtnID     string `json:"eqp_mtn_id" binding:"required"`
	EqpMtnStatus string `json:"eqp_mtn_status" binding:"required"` // ENUM: pending, in_progress, done, rejected
}
