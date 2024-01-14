package drive

type DriveUpdateDto struct {
	DriveName     string `gorm:"varchar" json:"drive_name"`
	DriveAccessId int    `gorm:"int4" json:"drive_access_id"`
}