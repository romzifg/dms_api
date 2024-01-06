package role_permissions

type RolePermissionDto struct {
	Code        string `gorm:"varchar" json:"code"`
	Description string `gorm:"varchar" json:"description"`
}

type RolePermissionInputDto struct {
	RoleId     int64               `gorm:"varchar" json:"role_id"`
	Permission []RolePermissionDto `json:"permissions"`
}
