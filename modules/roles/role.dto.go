package roles

import "time"

type UpdateRoleDto struct {
	RoleName  string    `json:"role_name"`
	Status    int16     `gorm:"int4" json:"status"`
	UpdatedAt time.Time `json:"updated_at" gorm:"timestamp"`
	UpdatedBy int16     `gorm:"int4" json:"updated_by"`
}