package service_test

import (
	"linked-page/db"
	"linked-page/dtos"
	"linked-page/model"
	"linked-page/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetListByListId(t *testing.T) {
	// Initialize the database
	db.InitTestDB(db.DB)
	db.DB.Exec("DELETE FROM lists")
	// defer db.CloseTestDB(db.DB)

	// Initialize the list service
	listService := new(service.ListService)
	listModel := new(model.ListModel)

	// Insert a new list
	insertListRequest := dtos.InsertListRequest{
		ListId: 1,
		UserId: 1,
		HeadId: 1,
	}
	_, err := listService.InsertList(&insertListRequest)
	assert.NoError(t, err)

	// Update the list
	patchListRequest := dtos.PatchListRequest{
		ListId: 1,
		UserId: 1,
		HeadId: 2,
	}
	err = listService.SetListByListId(&patchListRequest)
	assert.NoError(t, err)

	// Verify that the list was updated correctly
	list, _ := listModel.GetByListId(1)
	assert.Equal(t, uint(2), list.HeadId)

	// Try to update a list that does not exist
	patchListRequest = dtos.PatchListRequest{
		ListId: 2,
		UserId: 1,
		HeadId: 2,
	}
	err = listService.SetListByListId(&patchListRequest)
	assert.Error(t, err)
}
