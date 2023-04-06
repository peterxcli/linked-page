package model

import (
	"linked-page/db"
	"time"
)

type List struct {
	ListId    uint       `gorm:"primaryKey" json:"list_id,omitempty" `
	UserId    uint       `json:"user_id,omitempty"`
	HeadId    uint       `json:"head_id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type ListModel struct{}

func (m ListModel) GetByUserId(userId uint) (list *List, err error) {
	err = db.DB.Where("user_id = ?", userId).First(list).Error
	if err != nil {
		return nil, err
	}
	return list, err
}

func (m ListModel) GetByListId(listId uint) (list *List, err error) {
	list = &List{}
	err = db.DB.Where("list_id = ?", listId).First(list).Error
	if err != nil {
		return nil, err
	}
	return list, err
}

func (m ListModel) CreatePage(list *List) (id uint, err error) {
	err = db.DB.Create(list).Error
	if err != nil {
		return 0, err
	}
	return list.ListId, nil
}

func (m ListModel) UpdateByListId(list *List) (err error) {
	_, err = m.GetByListId(list.ListId)
	if err != nil {
		return err
	}
	result := db.DB.Model(&List{}).Where("list_id = ?", list.ListId).Updates(list)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (m ListModel) UpdateByUserId(list *List) (err error) {
	_, err = m.GetByUserId(list.UserId)
	if err != nil {
		return err
	}
	return db.DB.Model(&List{}).Where("user_id = ?", list.UserId).Updates(list).Error
}

func (m ListModel) SeedData() {
	// Check if any data already exists in the table
	db.DB.Exec("DELETE FROM lists")
	var count int64
	db.DB.Model(&List{}).Count(&count)
	if count == 0 {
		// Create sample articles
		lists := []List{
			{
				ListId: 1,
				HeadId: 1,
			},
			{
				ListId: 2,
				HeadId: 5,
			},
			{
				ListId: 3,
				HeadId: 9,
			},
			{
				ListId: 4,
				UserId: 1,
				HeadId: 13,
			},
		}
		return
		// Insert sample pages into the database
		for _, list := range lists {
			err := db.DB.Create(&list).Error
			if err != nil {
				panic("Failed to seed list data: " + err.Error())
			}
		}
	}
}
