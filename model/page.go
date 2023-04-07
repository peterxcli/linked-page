package model

import (
	"linked-page/db"
	"time"

	"github.com/lib/pq"
)

type PageModel struct{}

type Page struct {
	PageId     uint          `gorm:"primaryKey" json:"page_id,omitempty" `
	CreatedAt  *time.Time    `json:"created_at,omitempty"`
	UpdatedAt  *time.Time    `json:"updated_at,omitempty"`
	PrevPageId uint          `json:"prev_page_id,omitempty"`
	NextPageId uint          `json:"next_page_id,omitempty"`
	ArticleIds pq.Int64Array `gorm:"type:bigint[]" json:"article_ids"`
}

func (m PageModel) GetPageById(pageId uint) (page *Page, err error) {
	page = &Page{}
	err = db.DB.Where(&Page{PageId: pageId}).First(&page).Error
	if err != nil {
		return page, err
	}
	return page, nil
}

func (m PageModel) CreatePage(page *Page) (id uint, err error) {
	err = db.DB.Create(page).Error
	if err != nil {
		return 0, err
	}
	return page.PageId, nil
}

func (m PageModel) UpdatePrevPageId(page *Page) (err error) {
	err = db.DB.Model(&Page{}).Where("page_id = ?", page.PageId).Update("prev_page_id", page.PrevPageId).Error
	if err != nil {
		return
	}
	return nil
}

func (m PageModel) UpdateNextPageId(page *Page) (err error) {
	err = db.DB.Model(&Page{}).Where("page_id = ?", page.PageId).Update("next_page_id", page.NextPageId).Error
	if err != nil {
		return
	}
	return nil
}

func (m PageModel) UpdatePageContent(page *Page) (err error) {
	//transform page struct to map[string]interface{}
	// pageJson, _ := json.Marshal(page)
	// var pageMap map[string]interface{}
	// _ = json.Unmarshal(pageJson, &pageMap)
	err = db.DB.Model(&Page{}).Where("page_id = ?", page.PageId).Updates(page).Error
	if err != nil {
		return
	}
	return nil
}

func (m PageModel) DeletePageCertainHourBefore(hour int) (rowsAffected int, err error) {
	threshold := time.Now().Add(-time.Duration(hour) * time.Hour)
	result := db.DB.Where("created_at < ?", threshold).Delete(&Page{})
	if result.Error != nil {
		return rowsAffected, result.Error
	}
	return int(result.RowsAffected), nil
}

func (m PageModel) DeletePageById(pageId int64) error {
	err := db.DB.Delete(&Page{}, pageId).Error
	if err != nil {
		return err
	}
	return nil
}

func (m PageModel) SeedData() {
	// Check if any data already exists in the table
	db.DB.Exec("DELETE FROM pages")
	var count int64
	db.DB.Model(&Page{}).Count(&count)
	if count == 0 {
		// Create sample articles
		pages := []Page{
			{
				PageId:     1,
				NextPageId: 2,
				ArticleIds: []int64{1, 2, 3},
			},
			{
				PageId:     2,
				PrevPageId: 1,
				NextPageId: 3,
				ArticleIds: []int64{4, 5, 6},
			},
			{
				PageId:     3,
				PrevPageId: 2,
				NextPageId: 4,
				ArticleIds: []int64{7, 8, 9},
			},
			{
				PageId:     4,
				PrevPageId: 3,
				ArticleIds: []int64{10},
			},
			{
				PageId:     5,
				NextPageId: 6,
				ArticleIds: []int64{11, 12, 13},
			},
			{
				PageId:     6,
				PrevPageId: 5,
				NextPageId: 7,
				ArticleIds: []int64{14, 15, 16},
			},
			{
				PageId:     7,
				PrevPageId: 6,
				NextPageId: 8,
				ArticleIds: []int64{17, 18, 19},
			},
			{
				PageId:     8,
				PrevPageId: 7,
				ArticleIds: []int64{20},
			},
			{
				PageId:     9,
				NextPageId: 10,
				ArticleIds: []int64{21, 22, 23},
			},
			{
				PageId:     10,
				PrevPageId: 9,
				NextPageId: 11,
				ArticleIds: []int64{24, 25, 26},
			},
			{
				PageId:     11,
				PrevPageId: 10,
				NextPageId: 12,
				ArticleIds: []int64{27, 28, 29},
			},
			{
				PageId:     12,
				PrevPageId: 11,
				ArticleIds: []int64{30},
			},
			{
				PageId:     13,
				NextPageId: 14,
				ArticleIds: []int64{1, 11, 21},
			},
			{
				PageId:     14,
				PrevPageId: 13,
				NextPageId: 15,
				ArticleIds: []int64{2, 12, 22},
			},
			{
				PageId:     15,
				PrevPageId: 14,
				NextPageId: 16,
				ArticleIds: []int64{3, 13, 23},
			},
			{
				PageId:     16,
				PrevPageId: 15,
				ArticleIds: []int64{4},
			},
		}
		// db.DB.Create(&pages)
		// Insert sample pages into the database
		for _, page := range pages {
			err := db.DB.Create(&page).Error
			if err != nil {
				panic("Failed to seed page data: " + err.Error())
			}
		}
	}
}
