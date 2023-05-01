package models

type Comment struct {
	ID     string `gorm:"primary_key;unique" json:"id"` // Unique Comment ID
	Author string `json:"author"` // Author of the comment
	Body   string `json:"body"` // Body of the comment
}

type CreateCommentRequest struct {
	Author string `json:"author" binding:"required"`
	Body   string `json:"body" binding:"required"`
	CaptureID string `json:"captureId" binding:"required"`
}

// Get comments by array of IDs
func GetCaptureComments(ids []string) ([]Comment, error) {
	var comments []Comment

	if err := DB.Where("id IN (?)", ids).Find(&comments).Error; err != nil {
		return nil, err
	}

	return comments, nil
}