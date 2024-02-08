package auth

import (
	"time"
)

type Auth struct {
	UserId        	int                      `gorm:"primarykey" json:"user_id"`
	Username     	string                   `gorm:"varchar" json:"username"`
	Email 			string                   `gorm:"varchar" json:"email"`
	Password 		string                   `gorm:"varchar" json:"password"`
	CreatedAt     	time.Time                `gorm:"timestamptz" json:"created_at"`
	UpdatedBy     	int                      `gorm:"int4" json:"updated_by"`
	UpdatedAt     	time.Time                `gorm:"timestamptz" json:"updated_at"`
	DeletedBy     	int                      `gorm:"int4" json:"deleted_by"`
	DeletedAt     	time.Time                `gorm:"timestamptz" json:"deleted_at"`
}