package models

type Comment struct {
	ID     string `gorm:"primary_key;unique" json:"id"` // Unique Comment ID
	UserID string `gorm:"not null" json:"user_id"` // User ID of the user that created the comment
	CaptureID string `gorm:"not null" json:"capture_id"` // Capture ID of the capture that the comment belongs to
	Body   string `gorm:"not null" json:"body"` // Body of the comment
	UpdatedAt   int64 `gorm:"autoUpdateTime" json:"updated_at"` // Unix timestamp of when the user was last updated
	CreatedAt   int64 `gorm:"autoCreateTime" json:"created_at"` // Unix timestamp of when the user was created
}

type CreateCommentRequest struct {
	Body   string `json:"body" binding:"required"`
}
