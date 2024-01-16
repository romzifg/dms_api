package roles

import (
	"time"
)

type Role struct {
	RoleId    int     `gorm:"primarykey" json:"role_id"`
	RoleName  string    `gorm:"varchar" json:"role_name"`
	Status    int     `gorm:"int" json:"status"`
	CreatedBy int     `gorm:"int" json:"created_by"`
	CreatedAt time.Time `gorm:"timestamptz" json:"created_at"`
	UpdatedBy int     `gorm:"int" json:"updated_by"`
	UpdatedAt time.Time `gorm:"timestamptz" json:"updated_at"`
	DeletedBy int     `gorm:"int" json:"deleted_by"`
	DeletedAt time.Time `gorm:"timestamptz" json:"deleted_at"`
}