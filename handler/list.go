package handler

import (
	"linked-page/dtos"
	"linked-page/model"
	"linked-page/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ListHandler struct{}

var listModel = new(model.ListModel)
var listService = new(service.ListService)

func (lh *ListHandler) GetListByUserId(c *gin.Context) {
	userId, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	list, err := listModel.GetByUserId(uint(userId))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Page not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": list,
	})
}

func (lh *ListHandler) GetListByListId(c *gin.Context) {
	userId, err := strconv.ParseUint(c.Param("list_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	list, err := listModel.GetByListId(uint(userId))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": list,
	})
}

func (lh *ListHandler) SetListByListId(c *gin.Context) {
	patchListRequest := dtos.PatchListRequest{}
	if err := c.ShouldBindJSON(&patchListRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	listId, err := strconv.ParseUint(c.Param("list_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if listId != uint64(patchListRequest.ListId) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "list id in slug and json are not identical"})
		return
	}
	err = listService.SetListByListId(&patchListRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "success")
}
