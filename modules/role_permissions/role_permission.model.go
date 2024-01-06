package role_permissions

import (
	"dms_api/modules/roles"
	"time"
)

type RolePermission struct {
	RolePermissionId    int64     	`gorm:"primary_key" json:"role_permission_id"`
	RoleId  			int64    	`gorm:"int4" json:"role_id"`
	Role				roles.Role	`gorm:"foreignKey:RoleId"`	
	Code    			string     	`gorm:"varchar" json:"code"`
	Description    		string     	`gorm:"varchar" json:"description"`
	CreatedBy int16     `gorm:"int4" json:"created_by"`
	CreatedAt time.Time `gorm:"timestamptz" json:"created_at"`
	UpdatedBy int16     `gorm:"int4" json:"updated_by"`
	UpdatedAt time.Time `gorm:"timestamptz" json:"updated_at"`
	DeletedBy int16     `gorm:"int4" json:"deleted_by"`
	DeletedAt time.Time `gorm:"timestamptz" json:"deleted_at"`
}