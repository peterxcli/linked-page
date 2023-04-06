package handler

import (
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

type ListHandler struct{}

var listModel = new(model.ListModel)
var listService = new(service.ListService)

// GetListByUserId godoc
// @Summary Get all lists by user ID
// @Description Get all lists belonging to a user by their ID
// @Tags List
// @Param user_id path uint true "User ID"
// @Produce json
// @Success 200 {object} string "list content"
// @Failure 400 {object} string "error"
// @Failure 404 {object} string "error"
// @Router /list/user/{user_id} [get]
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

// GetListByListId godoc
// @Summary Get a list by ID
// @Description Get a list by its ID
// @Tags List
// @Param list_id path uint true "List ID"
// @Produce json
// @Success 200 {object} string "list content"
// @Failure 400 {object} string "error"
// @Failure 404 {object} string "error"
// @Router /list/{list_id} [get]
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

// SetListByListId godoc
// @Summary Update a list by ID
// @Description Update an existing list by its ID
// @Tags List
// @Param list_id path uint true "List ID"
// @Param patchListRequest body dtos.PatchListRequest true "Patch List Request"
// @Produce json
// @Success 200 {string} string "success"
// @Failure 400 {object} string "error"
// @Router /list/{list_id} [patch]
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

// InsertList godoc
// @Summary Creates a new list
// @Description Creates a new list and returns the ID of the newly created list.
// @ID create-list
// @Produce json
// @Tags List
// @Param list_id path int64 false "List ID (if empty, it will be generated randomly)"
// @Param body body dtos.InsertListRequest true "List object that needs to be added"
// @Success 200 {object} string "returns the ID of the newly created list"
// @Failure 400 {object} string "error message"
// @Router /list/{list_id} [post]
func (lh *ListHandler) InsertList(c *gin.Context) {
	insertListRequest := dtos.InsertListRequest{}
	if err := c.ShouldBindJSON(&insertListRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if c.Param("list_id") == "" {
		// fmt.Printf("%#v", insertPageRequest)
		source := rand.NewSource(time.Now().UnixNano())
		rng := rand.New(source)
		min := 1
		max := math.MaxUint32
		listId := uint(rng.Intn(max-min) + min)
		insertListRequest.ListId = listId
	} else {
		pageId, err := strconv.ParseUint(c.Param("list_id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if pageId != uint64(insertListRequest.ListId) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "list id in slug and json are not identical"})
			return
		}
	}

	newListId, err := listService.InsertList(&insertListRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"new_list_id": newListId})
}
