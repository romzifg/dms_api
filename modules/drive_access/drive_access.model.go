package drive_access

import "time"

type DriveAccess struct {
	DriveAccessId   int       `gorm:"primaryKey" json:"drive_access_id"`
	DriveAccessName string    `gorm:"varchar" json:"drive_access_name"`
	CreatedBy       int       `gorm:"int4" json:"created_by"`
	CreatedAt       time.Time `gorm:"timestamptz" json:"created_at"`
	UpdatedBy       int       `gorm:"int4" json:"updated_by"`
	UpdatedAt       time.Time `gorm:"timestamptz" json:"updated_at"`
	DeletedBy       int       `gorm:"int4" json:"deleted_by"`
	DeletedAt       time.Time `gorm:"timestamptz" json:"deleted_at"`
}