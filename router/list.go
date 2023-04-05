package router

import (
	"linked-page/handler"

	"github.com/gin-gonic/gin"
)

var listHandler = new(handler.ListHandler)

func ListRoute(router *gin.RouterGroup) {
	pageRouter := router.Group("/list")
	{
		pageRouter.GET("/:list_id", listHandler.GetListByListId)
		pageRouter.PATCH("/:list_id", listHandler.SetListByListId)
		pageRouter.GET("/user/:user_id", listHandler.GetListByUserId)
	}
}
