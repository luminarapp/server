package models

import "github.com/lib/pq"

type Collection struct {
	ID     string `gorm:"primary_key;unique" json:"id"` // Unique Collection ID 
	Name string `json:"name"` // Name of the collection 
	Description string `json:"description"` // Description of the collection 
	Captures pq.StringArray `gorm:"type:text[]" json:"captures"` // Array of capture IDs 
}

type CreateCollectionRequest struct {
	Name string `json:"name" binding:"required"`
	Description string `json:"description"` 
}

type AddCaptureToCollectionRequest struct {
	CaptureID string `json:"captureId" binding:"required"`
}

// Get captures by array of IDs
func GetCollectionCaptures(ids []string) ([]Capture, error) {
	var captures []Capture

	if err := DB.Where("id IN (?)", ids).Find(&captures).Error; err != nil {
		return nil, err
	}

	return captures, nil
}