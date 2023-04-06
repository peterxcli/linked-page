package service_test

import (
	"linked-page/db"
	"linked-page/dtos"
	"linked-page/service"
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	// _ "github.com/stretchr/testify/assert"
)

func TestInsertPage(t *testing.T) {
	// Initialize the database
	db.InitTestDB(db.DB)
	// defer db.CloseTestDB(db.DB)
	db.DB.Exec("DELETE FROM pages")

	// Initialize the page service
	pageService := new(service.PageService)

	// Insert the first page
	insertPageRequest := dtos.InsertPageRequest{
		PageId:     1,
		ArticleIds: []uint{1, 2, 3},
	}
	_, err := pageService.InsertPage(&insertPageRequest)
	assert.NoError(t, err)

	// Insert the second page after the first
	insertPageRequest = dtos.InsertPageRequest{
		PageId:     2,
		ArticleIds: []uint{4, 5, 6},
		PrevPageId: 1,
	}
	_, err = pageService.InsertPage(&insertPageRequest)
	assert.NoError(t, err)

	// Insert the third page after the second
	insertPageRequest = dtos.InsertPageRequest{
		PageId:     3,
		ArticleIds: []uint{7, 8, 9},
		PrevPageId: 2,
	}
	_, err = pageService.InsertPage(&insertPageRequest)
	assert.NoError(t, err)

	// Verify that the pages are linked correctly
	page1, _ := pageService.GetPage(1)
	page2, _ := pageService.GetPage(2)
	page3, _ := pageService.GetPage(3)
	assert.Equal(t, uint(0), page1.PrevPageId)
	assert.Equal(t, uint(2), page1.NextPageId)
	assert.Equal(t, uint(1), page2.PrevPageId)
	assert.Equal(t, uint(3), page2.NextPageId)
	assert.Equal(t, uint(2), page3.PrevPageId)
	assert.Equal(t, uint(0), page3.NextPageId)

	// Insert the new page between the first and the second page
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	min := 1
	max := math.MaxUint32
	new_page_id := uint(rng.Intn(max-min) + min)
	insertPageRequest = dtos.InsertPageRequest{
		PageId:     new_page_id,
		ArticleIds: []uint{7, 8, 9},
		PrevPageId: 1,
		NextPageId: 2,
	}
	_, err = pageService.InsertPage(&insertPageRequest)
	assert.NoError(t, err)
	page1, _ = pageService.GetPage(1)
	page2, _ = pageService.GetPage(2)
	new_page, _ := pageService.GetPage(new_page_id)
	assert.Equal(t, uint(0), page1.PrevPageId)
	assert.Equal(t, uint(new_page_id), page1.NextPageId)
	assert.Equal(t, uint(1), new_page.PrevPageId)
	assert.Equal(t, uint(2), new_page.NextPageId)
	assert.Equal(t, uint(new_page_id), page2.PrevPageId)
	assert.Equal(t, uint(3), page2.NextPageId)
}

func TestUpdatePage(t *testing.T) {
	// Initialize the database
	db.InitTestDB(db.DB)
	// defer db.CloseTestDB(db.DB)
	db.DB.Exec("DELETE FROM pages")

	// Initialize the page service
	pageService := new(service.PageService)

	// Insert a page
	insertPageRequest := dtos.InsertPageRequest{
		PageId:     1,
		ArticleIds: []uint{1, 2, 3},
	}
	_, err := pageService.InsertPage(&insertPageRequest)
	assert.NoError(t, err)

	// Update the page
	patchPageRequest := dtos.PatchPageRequest{
		PageId:     1,
		ArticleIds: []uint{4, 5, 6},
	}
	err = pageService.SetPage(&patchPageRequest)
	assert.NoError(t, err)

	// Verify that the page was updated correctly
	page, _ := pageService.GetPage(1)
	actual := make([]uint, len(page.ArticleIds))
	for i, id := range page.ArticleIds {
		actual[i] = uint(id)
	}

	assert.Equal(t, []uint{4, 5, 6}, actual)

}
