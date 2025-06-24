package internalproposal

import (
	internalproposal "system-management-pg/internal/controller/internal-proposal"
	"system-management-pg/internal/middlewares"

	"github.com/gin-gonic/gin"
)

type InternalProposalRouter struct{}

func (pr *InternalProposalRouter) InitInternalProposalRouter(Router *gin.RouterGroup) {

	internalProposalRouterPrivate := Router.Group("/internal-proposal")
	internalProposalRouterPrivate.Use(middlewares.AuthenMiddlewareAccount())
	internalProposalRouterPrivate.Use(middlewares.LogApiMiddleware())
	{
		internalProposalRouterPrivate.POST("", internalproposal.InternalProposal.CreateInternalProposal)
		internalProposalRouterPrivate.GET("/:id", internalproposal.InternalProposal.FindInternalProposal)
		internalProposalRouterPrivate.PATCH("", internalproposal.InternalProposal.UpdateInternalProposal)
		internalProposalRouterPrivate.DELETE("/:id", internalproposal.InternalProposal.DeleteInternalProposal)
		internalProposalRouterPrivate.PATCH("/restore/:id", internalproposal.InternalProposal.RestoreInternalProposal)
		internalProposalRouterPrivate.GET("", internalproposal.InternalProposal.GetAllInternalProposal)
		internalProposalRouterPrivate.GET("/recycle", internalproposal.InternalProposal.GetAllInternalProposalRecycle)
		internalProposalRouterPrivate.PATCH("/update-status", internalproposal.InternalProposal.UpdateInternalProposalStatus)
	}

}
