package dtos

type GetPageResponse struct {
	PageId         uint   `json:"page_id"`
	NextPageId     uint   `json:"next_page_id"`
	PreviousPageId uint   `json:"previous_page_id"`
	ArticleIds     []uint `json:"article_ids"`
}

type InsertPageRequest struct {
	PageId     uint   `json:"page_id"`
	PrevPageId uint   `json:"prev_page_id"`
	NextPageId uint   `json:"next_page_id"`
	ArticleIds []uint `json:"article_ids"`
}

type PatchPageRequest struct {
	PageId     uint   `json:"page_id"`
	ArticleIds []uint `json:"article_ids"`
}
