package models

type Comment struct {
	ID     string `gorm:"primary_key;unique" json:"id"` // Unique Comment ID
	Author string `gorm:"not null" json:"author"` // Author of the comment
	Body   string `gorm:"not null" json:"body"` // Body of the comment
	UpdatedAt   int64 `gorm:"autoUpdateTime"` // Unix timestamp of when the comment was last updated
	CreatedAt   int64 `gorm:"autoCreateTime"` // Unix timestamp of when the comment was created
}

type CreateCommentRequest struct {
	Author string `json:"author" binding:"required"`
	Body   string `json:"body" binding:"required"`
}

// Get comments by array of IDs
func GetCaptureComments(ids []string) ([]Comment, error) {
	var comments []Comment

	if err := DB.Where("id IN (?)", ids).Find(&comments).Error; err != nil {
		return nil, err
	}

	return comments, nil
}