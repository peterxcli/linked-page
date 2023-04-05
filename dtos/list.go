package dtos

type PatchListRequest struct {
	ListId uint `gorm:"primaryKey" json:"list_id,omitempty" `
	UserId uint `json:"user_id,omitempty"`
	HeadId uint `json:"head_id,omitempty"`
}
