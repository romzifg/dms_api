package roles

import "time"

type Role struct {
	RoleId    int64     `gorm:"primary_key" json:"role_id"`
	RoleName  string    `gorm:"varchar" json:"role_name"`
	Status    int16     `gorm:"int4" json:"status"`
	CreatedBy int16     `gorm:"int4" json:"created_by"`
	CreatedAt time.Time `gorm:"timestamptz" json:"created_at"`
	UpdatedBy int16     `gorm:"int4" json:"updated_by"`
	UpdatedAt time.Time `gorm:"timestamptz" json:"updated_at"`
	DeletedBy int16     `gorm:"int4" json:"deleted_by"`
	DeletedAt time.Time `gorm:"timestamptz" json:"deleted_at"`
}