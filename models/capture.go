package models

import "github.com/lib/pq"

type Capture struct {
	ID     string   `gorm:"primary_key;unique" json:"id"` // Unique Capture ID
	Author string `json:"author"` // Author of the capture 
	Source string `json:"source"` // Source of the capture
	Comments pq.StringArray `gorm:"type:text[]" json:"comments"` // Array of comment IDs
}

type CreateCaptureRequest struct {
	Author string `json:"author" binding:"required"`
	Source string `json:"source" binding:"required"`
}