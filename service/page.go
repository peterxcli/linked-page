package service

import (
	"linked-page/dtos"
	"linked-page/model"

	"github.com/lib/pq"
)

type PageService struct{}

var pageModel = new(model.PageModel)

func (ps PageService) GetPage(pageId uint) (page *model.Page, err error) {
	page, error := pageModel.GetPageById(pageId)
	if error != nil {
		return page, error
	}
	return page, nil
}

func (ps PageService) SetPage(patchPage *dtos.PatchPageRequest) (err error) {
	article_ids := pq.Int64Array{}
	for _, articleId := range patchPage.ArticleIds {
		article_ids = append(article_ids, int64(articleId))
	}
	page := model.Page{
		PageId:     patchPage.PageId,
		ArticleIds: article_ids,
	}
	err = pageModel.UpdatePageContent(&page)
	if err != nil {
		return err
	}
	return nil
}

func (ps PageService) InsertPage(newPage *dtos.InsertPageRequest) (newPageId uint, err error) {
	article_ids := pq.Int64Array{}
	for _, articleId := range newPage.ArticleIds {
		article_ids = append(article_ids, int64(articleId))
	}
	if newPage.PrevPageId == 0 && newPage.NextPageId == 0 {
		page := model.Page{
			PageId:     newPage.PageId,
			ArticleIds: article_ids,
		}
		newPageId, err = pageModel.CreatePage(&page)
		if err != nil {
			return newPageId, err
		}
		return newPageId, err
	} else if newPage.PrevPageId == 0 {
		page := model.Page{
			PageId:     newPage.PageId,
			ArticleIds: article_ids,
			NextPageId: newPage.NextPageId,
		}
		newPageId, err = pageModel.CreatePage(&page)
		if err != nil {
			return newPageId, err
		}
		//next->prev_id = page.page_id
		err = pageModel.UpdatePrevPageId(&model.Page{PageId: page.NextPageId, PrevPageId: page.PageId})
		if err != nil {
			return newPageId, err
		}
		return newPageId, nil
	} else if newPage.NextPageId == 0 {
		page := model.Page{
			PageId:     newPage.PageId,
			ArticleIds: article_ids,
			PrevPageId: newPage.PrevPageId,
		}
		newPageId, err = pageModel.CreatePage(&page)
		if err != nil {
			return newPageId, err
		}
		//prev->next_id = page.page_id
		err = pageModel.UpdateNextPageId(&model.Page{PageId: page.PrevPageId, NextPageId: page.PageId})
		if err != nil {
			return newPageId, err
		}
		return newPageId, nil
	} else {
		page := model.Page{
			PageId:     newPage.PageId,
			ArticleIds: article_ids,
			PrevPageId: newPage.PrevPageId,
			NextPageId: newPage.NextPageId,
		}
		newPageId, err = pageModel.CreatePage(&page)
		if err != nil {
			return newPageId, err
		}

		//prev->next_id = page.page_id
		err = pageModel.UpdateNextPageId(&model.Page{PageId: page.PrevPageId, NextPageId: page.PageId})
		if err != nil {
			return newPageId, err
		}

		//next->prev_id = page.page_id
		err = pageModel.UpdatePrevPageId(&model.Page{PageId: page.NextPageId, PrevPageId: page.PageId})
		if err != nil {
			return newPageId, err
		}

		return newPageId, err
	}
}
