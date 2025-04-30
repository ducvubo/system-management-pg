package service

import (
	"context"
	"system-management-pg/internal/model"
	"system-management-pg/pkg/response"
)

type (
	IInternalProposal interface {
		CreateInternalProposal(ctx context.Context, createInternalProposal *model.CreateInternalProposalDto, Account *model.Account) (err error, statusCode int)
		FindInternalProposal(ctx context.Context, ItnProposalId string, Account *model.Account) (out *model.InternalProposalOutput, err error, statusCode int)
		UpdateInternalProposal(ctx context.Context, updateInternalProposal *model.UpdateInternalProposalDto, Account *model.Account) (err error, statusCode int)
		DeleteInternalProposal(ctx context.Context, ItnProposalId string, Account *model.Account) (err error, statusCode int)
		RestoreInternalProposal(ctx context.Context, ItnProposalId string, Account *model.Account) (err error, statusCode int)
		GetAllInternalProposal(ctx context.Context, Limit int32, Offset int32, isDeleted int32, ItnProposalTitle string, Account *model.Account) (out response.ModelPagination[[]*model.InternalProposalOutput], err error, statusCode int)
		UpdateInternalProposalStatus(ctx context.Context, updateInternalProposalStatus *model.UpdateInternalProposalStatusDto, Account *model.Account) (err error, statusCode int)
	}
)

var (
	localInternalProposal IInternalProposal
)

func InternalProposal() IInternalProposal {
	if localInternalProposal == nil {
		panic("implement.............................................")
	}
	return localInternalProposal
}

func InitInternalProposal(i IInternalProposal) {
	localInternalProposal = i
}
