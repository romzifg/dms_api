package drive

import (
	"dms_api/modules/drive_access"
	"time"
)

type Drive struct {
	DriveId     int       					`gorm:"primary_key" json:"drive_id"`
	DriveName   string    					`gorm:"varchar" json:"drive_name"`
	DriveAccessId int        				`gorm:"int4" json:"drive_access_id"`
	DriveAccess drive_access.DriveAccess 	`gorm:"foreignKey:DriveAccessId" json:"driveaccess"`	
	CreatedBy   int       					`gorm:"int4" json:"created_by"`
	CreatedAt   time.Time 					`gorm:"timestamptz" json:"created_at"`
	UpdatedBy   int       					`gorm:"int4" json:"updated_by"`
	UpdatedAt   time.Time 					`gorm:"timestamptz" json:"updated_at"`
	DeletedBy   int       					`gorm:"int4" json:"deleted_by"`
	DeletedAt   time.Time 					`gorm:"timestamptz" json:"deleted_at"`
}