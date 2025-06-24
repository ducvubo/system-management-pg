package internalnote

import (
	internalnote "system-management-pg/internal/controller/internal-note"
	"system-management-pg/internal/middlewares"

	"github.com/gin-gonic/gin"
)

type InternalNoteRouter struct{}

func (pr *InternalNoteRouter) InitInternalNoteRouter(Router *gin.RouterGroup) {

	internalProposalRouterPrivate := Router.Group("/internal-note")
	internalProposalRouterPrivate.Use(middlewares.AuthenMiddlewareAccount())
	internalProposalRouterPrivate.Use(middlewares.LogApiMiddleware())
	{
		internalProposalRouterPrivate.POST("", internalnote.InternalNote.CreateInternalNote)
		internalProposalRouterPrivate.GET("/:id", internalnote.InternalNote.FindInternalNote)
		internalProposalRouterPrivate.PATCH("", internalnote.InternalNote.UpdateInternalNote)
		internalProposalRouterPrivate.DELETE("/:id", internalnote.InternalNote.DeleteInternalNote)
		internalProposalRouterPrivate.PATCH("/restore/:id", internalnote.InternalNote.RestoreInternalNote)
		internalProposalRouterPrivate.GET("", internalnote.InternalNote.GetAllInternalNote)
		internalProposalRouterPrivate.GET("/recycle", internalnote.InternalNote.GetAllInternalNoteRecycle)
	}

}
