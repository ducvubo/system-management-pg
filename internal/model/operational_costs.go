package model

type OperationalCostsOutput struct {
	OperaCostID          string `json:"opera_cost_id"`
	OperaCostType        string `json:"opera_cost_type"`
	OperaCostAmount      *int64 `json:"opera_cost_amount"`
	OperaCostDescription string `json:"opera_cost_description,omitempty"`
	OperaCostDate        string `json:"opera_cost_date"`
	OperaCostStatus      string `json:"opera_cost_status"`
}

type CreateOperationalCostsDto struct {
	OperaCostType        string `json:"opera_cost_type" binding:"required"`
	OperaCostAmount      *int64 `json:"opera_cost_amount" binding:"required"`
	OperaCostDescription string `json:"opera_cost_description" binding:"required"`
	OperaCostDate        string `json:"opera_cost_date" binding:"required"`
}

type UpdateOperationalCostsDto struct {
	OperaCostID          string `json:"opera_cost_id" binding:"required"`
	OperaCostType        string `json:"opera_cost_type" binding:"required"`
	OperaCostAmount      *int64 `json:"opera_cost_amount" binding:"required"`
	OperaCostDescription string `json:"opera_cost_description" binding:"required"`
	OperaCostDate        string `json:"opera_cost_date" binding:"required"`
}

type UpdateOperationalCostsStatusDto struct {
	OperaCostID     string `json:"opera_cost_id" binding:"required"`
	OperaCostStatus string `json:"opera_cost_status" binding:"required"` // ENUM: pending, in_progress, done, rejected
}
