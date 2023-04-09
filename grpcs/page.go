package grpcs

import (
	"context"
	"linked-page/model"
	"linked-page/proto/page"

	"github.com/lib/pq"
)

var pageModel = new(model.PageModel)

type PageGRPC struct {
	page.UnimplementedPageServer
}

func NewPageGRPC() *PageGRPC {
	return &PageGRPC{}
}

func (p *PageGRPC) GetPage(ctx context.Context, req *page.GetPageRequest) (pageResponse *page.GetPageResponse, err error) {
	pageRes, err := pageModel.GetPageById(uint(req.PageId))
	if err != nil {
		return pageResponse, err
	}
	article_ids := []uint32{}
	for _, articleId := range pageRes.ArticleIds {
		article_ids = append(article_ids, uint32(articleId))
	}
	pageResponse = &page.GetPageResponse{
		PageId:     uint32(pageRes.PageId),
		NextPageId: uint32(pageRes.NextPageId),
		PrevPageId: uint32(pageRes.PrevPageId),
		ArticleIds: article_ids,
	}
	return pageResponse, nil
}

func (p *PageGRPC) SetPage(ctx context.Context, req *page.SetPageRequest) (pageResponse *page.SetPageResponse, err error) {
	article_ids := pq.Int64Array{}
	for _, articleId := range req.ArticleIds {
		article_ids = append(article_ids, int64(articleId))
	}
	pageUpdate := model.Page{
		PageId:     uint(req.PageId),
		ArticleIds: article_ids,
	}
	err = pageModel.UpdatePageContent(&pageUpdate)
	if err != nil {
		return &page.SetPageResponse{IsSuccess: false}, err
	}
	return &page.SetPageResponse{IsSuccess: true}, nil
}

func (p *PageGRPC) InsertPage(ctx context.Context, req *page.InsertPageRequest) (pageResponse *page.InsertPageResponse, err error) {
	article_ids := pq.Int64Array{}
	for _, articleId := range req.ArticleIds {
		article_ids = append(article_ids, int64(articleId))
	}
	if req.PrevPageId == 0 && req.NextPageId == 0 {
		pageCreate := model.Page{
			PageId:     uint(req.PageId),
			ArticleIds: article_ids,
		}
		newPageId, err := pageModel.CreatePage(&pageCreate)
		if err != nil {
			return &page.InsertPageResponse{IsSuccess: false}, err
		}
		return &page.InsertPageResponse{IsSuccess: true, NewPageId: uint32(newPageId)}, err
	} else if req.PrevPageId == 0 {
		pageCreate := model.Page{
			PageId:     uint(req.PageId),
			ArticleIds: article_ids,
		}
		newPageId, err := pageModel.CreatePage(&pageCreate)
		if err != nil {
			return &page.InsertPageResponse{IsSuccess: false}, err
		}
		//next->prev_id = page.page_id
		err = pageModel.UpdatePrevPageId(&model.Page{PageId: pageCreate.NextPageId, PrevPageId: pageCreate.PageId})
		if err != nil {
			return &page.InsertPageResponse{IsSuccess: false}, err
		}
		return &page.InsertPageResponse{IsSuccess: true, NewPageId: uint32(newPageId)}, err
	} else if req.NextPageId == 0 {
		pageCreate := model.Page{
			PageId:     uint(req.PageId),
			ArticleIds: article_ids,
		}
		newPageId, err := pageModel.CreatePage(&pageCreate)
		if err != nil {
			return &page.InsertPageResponse{IsSuccess: false}, err
		}
		//prev->next_id = page.page_id
		err = pageModel.UpdateNextPageId(&model.Page{PageId: pageCreate.PrevPageId, NextPageId: pageCreate.PageId})
		if err != nil {
			return &page.InsertPageResponse{IsSuccess: false}, err
		}
		return &page.InsertPageResponse{IsSuccess: true, NewPageId: uint32(newPageId)}, err
	} else {
		pageCreate := model.Page{
			PageId:     uint(req.PageId),
			ArticleIds: article_ids,
		}
		newPageId, err := pageModel.CreatePage(&pageCreate)
		if err != nil {
			return &page.InsertPageResponse{IsSuccess: false}, err
		}
		//prev->next_id = page.page_id
		err = pageModel.UpdateNextPageId(&model.Page{PageId: pageCreate.PrevPageId, NextPageId: pageCreate.PageId})
		if err != nil {
			return &page.InsertPageResponse{IsSuccess: false}, err
		}

		//next->prev_id = page.page_id
		err = pageModel.UpdatePrevPageId(&model.Page{PageId: pageCreate.NextPageId, PrevPageId: pageCreate.PageId})
		if err != nil {
			return &page.InsertPageResponse{IsSuccess: false}, err
		}
		return &page.InsertPageResponse{IsSuccess: true, NewPageId: uint32(newPageId)}, err
	}
}

func (p *PageGRPC) DeletePage(ctx context.Context, req *page.DeletePageRequest) (pageResponse *page.DeletePageResponse, err error) {
	rowAffected, err := pageModel.DeletePageById(int64(req.PageId))
	if err != nil {
		return &page.DeletePageResponse{}, err
	}
	return &page.DeletePageResponse{RowAffected: uint32(rowAffected)}, nil
}
