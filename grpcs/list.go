package grpcs

import (
	"context"
	"errors"
	"linked-page/model"
	"linked-page/proto/list"
)

var listModel = new(model.ListModel)

type ListGRPC struct {
	list.UnimplementedListServer
}

func NewListGRPC() *ListGRPC {
	return &ListGRPC{}
}

func (l *ListGRPC) GetList(ctx context.Context, req *list.GetListRequest) (*list.GetListResponse, error) {
	if req.ListId != 0 {
		result, err := listModel.GetByListId(uint(req.ListId))
		if err != nil {
			return &list.GetListResponse{}, err
		}
		response := &list.GetListResponse{
			ListId: uint32(result.ListId),
			UserId: uint32(result.UserId),
			HeadId: uint32(result.HeadId),
		}
		return response, nil
	} else if req.UserId != 0 {
		result, err := listModel.GetByUserId(uint(req.UserId))
		if err != nil {
			return &list.GetListResponse{}, err
		}
		response := &list.GetListResponse{
			ListId: uint32(result.ListId),
			UserId: uint32(result.UserId),
			HeadId: uint32(result.HeadId),
		}
		return response, nil
	} else {
		return &list.GetListResponse{}, errors.New("field valid")
	}
}

func (l *ListGRPC) InsertList(ctx context.Context, req *list.InsertListRequest) (*list.SetHeadResponse, error) {
	listCreate := model.List{
		ListId: uint(req.ListId),
		UserId: uint(req.UserId),
		HeadId: uint(req.HeadId),
	}
	_, err := listModel.CreatePage(&listCreate)
	if err != nil {
		return &list.SetHeadResponse{IsSuccess: false}, err
	}
	return &list.SetHeadResponse{IsSuccess: true}, nil
}

func (l *ListGRPC) PatchList(ctx context.Context, req *list.PatchListRequest) (*list.SetHeadResponse, error) {
	listPatch := model.List{
		ListId: uint(req.ListId),
		UserId: uint(req.UserId),
		HeadId: uint(req.HeadId),
	}
	err := listModel.UpdateByListId(&listPatch)
	if err != nil {
		return &list.SetHeadResponse{IsSuccess: false}, err
	}
	return &list.SetHeadResponse{IsSuccess: true}, nil
}
