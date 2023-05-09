package models

type User struct {
	ID     string `gorm:"primary_key;unique" json:"id"` // Unique User ID
	Username   string `gorm:"not null" json:"username"` // Username of the user
	Email  string `gorm:"not null;unique" json:"email"` // Email of the user
	Collections []Collection `gorm:"foreignKey:UserID;default:'{}'" json:"-"` // Array of collections 
	Captures []Capture `gorm:"foreignKey:UserID;default:'{}'" json:"-"` // Array of captures
	AuthToken AuthToken `gorm:"embedded" json:"-"` // Auth token of the user
	UpdatedAt   int64 `gorm:"autoUpdateTime" json:"updated_at"` // Unix timestamp of when the user was last updated
	CreatedAt   int64 `gorm:"autoCreateTime" json:"created_at"` // Unix timestamp of when the user was created
}

type CreateUserRequest struct {
	Email string `json:"email" binding:"required"`
}

type UpdateUserRequest struct {
	Username string `json:"username"`
}

type AuthToken struct {
	Token string `json:"token"`
	ExpiresAt int64 `json:"expires_at"`
}

type UserAuthChallengeStatus struct {
	Success bool `json:"success"`
	AuthStep string `json:"authStep"`
	ErrorMsg string `json:"errorMsg,omitempty"` 
	Session string `json:"session,omitempty"`
}

type UserAuthChallengeRequest struct {
	Email string `json:"email" binding:"required"`
	Token string `json:"token"` // If token is not provided, a new token will be generated and sent to the user's email
}
