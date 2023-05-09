package models

type Collection struct {
	ID     string `gorm:"primary_key;unique" json:"id"` // Unique Collection ID
	UserID string `gorm:"not null" json:"user_id"` // User ID of the user that created the collection
	Name string `gorm:"not null" json:"name"` // Name of the collection
	Description string `gorm:"not null" json:"description"` // Description of the collection
	Captures []Capture `gorm:"foreignKey:CollectionID;default:'{}'" json:"captures"` // Array of captures
	UpdatedAt   int64 `gorm:"autoUpdateTime" json:"updated_at"` // Unix timestamp of when the user was last updated
	CreatedAt   int64 `gorm:"autoCreateTime" json:"created_at"` // Unix timestamp of when the user was created
}

type CreateCollectionRequest struct {
	Name string `json:"name" binding:"required"`
	Description string `json:"description"` 
}

type UpdateCollectionRequest struct {
	Name string `json:"name"`
	Description string `json:"description"`
}

// Get collections by user ID
func GetCollectionsByUserId(id string) ([]Collection, error) {
	var collections []Collection

	// Dont preload as used for user collections router which doesnt need captures
	if err := DB.Where("user_id = ?", id).Find(&collections).Error; err != nil {
		return nil, err
	}

	return collections, nil
}
