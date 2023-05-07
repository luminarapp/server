package models

type Capture struct {
	ID     string   `gorm:"primary_key;unique" json:"id"` // Unique Capture ID
	UserID string `gorm:"not null" json:"user_id"` // User ID of the user that owns the capture
	CollectionID string `gorm:"not null" json:"collection_id"` // Collection ID of the collection that the capture belongs to
	// TODO: Add this back as might be useful for frontend
	// Collection Collection `gorm:"foreignKey:CollectionID" json:"collection"` // Collection that the capture belongs to 
	Reference string `gorm:"not null" json:"reference"` // Reference of the capture //TODO: Think about renaming this to something more descriptive
	Comments []Comment `gorm:"foreignKey:CaptureID;default:'{}'" json:"comments"` // Array of comments
	UpdatedAt   int64 `gorm:"autoUpdateTime" json:"updated_at"` // Unix timestamp of when the user was last updated
	CreatedAt   int64 `gorm:"autoCreateTime" json:"created_at"` // Unix timestamp of when the user was created
}

type CreateCaptureRequest struct {
	Reference string `json:"reference" binding:"required"`
	CollectionID string `json:"collection_id"`
}

// GetCapturesByUserId gets captures by user ID
func GetCapturesByUserId(id string) ([]Capture, error) {
	var captures []Capture

	if err := DB.Preload("Comments").Where("user_id = ?", id).Find(&captures).Error; err != nil {
		return nil, err
	}

	return captures, nil
}