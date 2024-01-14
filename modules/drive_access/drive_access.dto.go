package drive_access

type DriveAccessUpdateDto struct {
	DriveAccessName string `gorm:"varchar" json:"drive_access_name"`
}