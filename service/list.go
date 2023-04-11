package service

import (
	"linked-page/dtos"
	"linked-page/model"
)

type ListServiceInterface interface {
	SetListByListId(patchList *dtos.PatchListRequest) (err error)
	InsertList(insertList *dtos.InsertListRequest) (newListId uint, err error)
}

type ListService struct{}

var listModel = new(model.ListModel)

func (ls *ListService) SetListByListId(patchList *dtos.PatchListRequest) (err error) {
	list := model.List{
		ListId: patchList.ListId,
		UserId: patchList.UserId,
		HeadId: patchList.HeadId,
	}
	err = listModel.UpdateByListId(&list)
	if err != nil {
		return err
	}
	return nil
}

func (ls *ListService) InsertList(insertList *dtos.InsertListRequest) (newListId uint, err error) {
	list := model.List{
		ListId: insertList.ListId,
		UserId: insertList.UserId,
		HeadId: insertList.HeadId,
	}
	newListId, err = listModel.CreatePage(&list)
	if err != nil {
		return newListId, err
	}
	return newListId, err
}
