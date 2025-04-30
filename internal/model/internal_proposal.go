package model

type InternalProposalOutput struct {
	ItnProposalId      string `json:"itn_proposal_id" binding:"required"`
	ItnProposalTitle   string `json:"itn_proposal_title" binding:"required"`
	ItnProposalContent string `json:"itn_proposal_content" binding:"required"`
	ItnProposalType    string `json:"itn_proposal_type" binding:"required"`
	ItnProposalStatus  string `json:"itn_proposal_status" binding:"required"`
	Isdeleted          int32  `json:"isDeleted"`
}

type CreateInternalProposalDto struct {
	ItnProposalTitle   string `json:"itn_proposal_title" binding:"required"`
	ItnProposalContent string `json:"itn_proposal_content" binding:"required"`
	ItnProposalType    string `json:"itn_proposal_type" binding:"required"`
}

type UpdateInternalProposalDto struct {
	ItnProposalId      string `json:"itn_proposal_id" binding:"required"`
	ItnProposalTitle   string `json:"itn_proposal_title" binding:"required"`
	ItnProposalContent string `json:"itn_proposal_content" binding:"required"`
	ItnProposalType    string `json:"itn_proposal_type" binding:"required"`
}

type UpdateInternalProposalStatusDto struct {
	ItnProposalId     string `json:"itn_proposal_id" binding:"required"`
	ItnProposalStatus string `json:"itn_proposal_status" binding:"required"`
}