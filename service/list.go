package service

import (
	"linked-page/dtos"
	"linked-page/model"
)

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
