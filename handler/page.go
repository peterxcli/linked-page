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

// GetPage godoc
// @Summary Get a page by ID
// @Description Get a page by its ID
// @Tags Page
// @Param page_id path uint true "Page ID"
// @Produce json
// @Success 200 {object} string "error"
// @Failure 400 {object} string "error"
// @Failure 404 {object} string "page content"
// @Router /page/{page_id} [get]
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

// SetPage godoc
// @Summary Update a page
// @Description Update an existing page by its ID
// @Tags Page
// @Param page_id path uint true "Page ID"
// @Param PrevPageId body uint false "Previous page ID"
// @Param NextPageId body uint false "Next page ID"
// @Param ArticleIds body []int64 false "Article IDs"
// @Produce json
// @Success 200 {object} string "success"
// @Failure 400 {object} string "error"
// @Failure 404 {object} string "error"
// @Router /page/{page_id} [patch]
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

// InsertPage godoc
// @Summary Insert a new page, if the pageId doesnt specify, then the newPageId would be a random number
// @Tags Page
// @Accept json
// @Produce json
// @Param insertPageRequest body dtos.InsertPageRequest true "Insert Page Request"
// @Success 200 {integer} integer "Page ID"
// @Failure 400 {object} string "error"
// @Router /page [post]
// @Description If PrevPageId is provided but NextPageId is not, the new page will be inserted after the page with the specified PrevPageId, forming a linked-list-like structure. Similarly, if NextPageId is provided but PrevPageId is not, the new page will be inserted before the page with the specified NextPageId, forming a linked-list-like structure.
// @Tags Page
// @Accept json
// @Produce json
// @Param insertPageRequest body dtos.InsertPageRequest true "Insert Page Request"
// @Success 200 {integer} integer "Page ID"
// @Failure 400 {object} string "error"
// @Router /page [post]
// @Router /page/{page_id} [post]
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

// DeleteCertainHourBeforeHandler godoc
// @Summary Delete all pages with updated time before a certain hour
// @Description Delete all pages with updated time before a certain hour
// @Tags Page
// @Param hour query int true "Hour"
// @Produce json
// @Success 200 {object} string "rowAffected"
// @Failure 400 {object} string "error"
// @Router /page/hours [delete]
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
