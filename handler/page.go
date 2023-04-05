package handler

import (
	"fmt"
	"linked-page/dtos"
	"linked-page/model"
	"linked-page/service"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type PageHandler struct{}

var pageService = new(service.PageService)

func (ph *PageHandler) GetPage(c *gin.Context) {
	pageId, err := strconv.ParseUint(c.Param("page_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	page, err := pageService.GetPage(uint(pageId))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Page not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": page,
	})
}

func (ph *PageHandler) SetPage(c *gin.Context) {
	patchPageRequest := dtos.PatchPageRequest{}
	if err := c.ShouldBindJSON(&patchPageRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pageId, err := strconv.ParseUint(c.Param("page_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if pageId != uint64(patchPageRequest.PageId) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "page id in slug and json are not identical"})
		return
	}
	err = pageService.SetPage(&patchPageRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "success")
}

func (ph *PageHandler) InsertPage(c *gin.Context) {
	insertPageRequest := dtos.InsertPageRequest{}
	if err := c.ShouldBindJSON(&insertPageRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if c.Param("page_id") == "" {
		// fmt.Printf("%#v", insertPageRequest)
		source := rand.NewSource(time.Now().UnixNano())
		rng := rand.New(source)
		min := 1
		max := math.MaxUint32
		pageId := uint(rng.Intn(max-min) + min)
		insertPageRequest.PageId = pageId
	} else {
		pageId, err := strconv.ParseUint(c.Param("page_id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if pageId != uint64(insertPageRequest.PageId) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "page id in slug and json are not identical"})
			return
		}
	}

	newPageId, err := pageService.InsertPage(&insertPageRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"new_page_id": newPageId})
}

var pageModel = new(model.PageModel)

func (ph *PageHandler) DeleteCertainHourBeforeHandler(c *gin.Context) {
	hourParam := c.Query("hour")
	hour, err := strconv.Atoi(hourParam)
	// fmt.Println("hour: ", hour)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid hour parameter"})
		return
	}
	rowsAffected, err := pageModel.DeletePageCertainHourBefore(hour)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("%d pages deleted successfully", rowsAffected)})
}
