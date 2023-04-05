package router

import (
	"linked-page/handler"

	"github.com/gin-gonic/gin"
)

var pageHandler = new(handler.PageHandler)

func PageRoute(router *gin.RouterGroup) {
	pageRouter := router.Group("/page")
	{
		pageRouter.GET("/:page_id", pageHandler.GetPage)
		pageRouter.PATCH("/:page_id", pageHandler.SetPage)
		pageRouter.POST("/:page_id", pageHandler.InsertPage)
		pageRouter.POST("/", pageHandler.InsertPage)
		pageRouter.DELETE("/hours", pageHandler.DeleteCertainHourBeforeHandler)
	}
}
