package models

import "github.com/lib/pq"

type Capture struct {
	ID     string   `gorm:"primary_key;unique" json:"id"` // Unique Capture ID
	Author string `gorm:"not null" json:"author"` // Author of the capture
	Source string `gorm:"not null" json:"source"` // Source of the capture
	Comments pq.StringArray `gorm:"type:text[]" json:"comments"` // Array of comment IDs
	UpdatedAt   int64 `gorm:"autoUpdateTime"` // Unix timestamp of when the capture was last updated
	CreatedAt   int64 `gorm:"autoCreateTime"` // Unix timestamp of when the capture was created
}

type CreateCaptureRequest struct {
	Author string `json:"author" binding:"required"`
	Source string `json:"source" binding:"required"`
}