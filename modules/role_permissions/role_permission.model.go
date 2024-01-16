package role_permissions

import (
	"dms_api/modules/roles"
	"time"
)

type RolePermission struct {
	RolePermissionId    int     	`gorm:"primarykey" json:"role_permission_id"`
	RoleId  			int    	`gorm:"int" json:"role_id"`
	Role				roles.Role	`gorm:"foreignKey: role_id; references: role_id" json:"role"`	
	Code    			string     	`gorm:"varchar" json:"code"`
	Description    		string     	`gorm:"varchar" json:"description"`
	CreatedBy int     `gorm:"int4" json:"created_by"`
	CreatedAt time.Time `gorm:"timestamptz" json:"created_at"`
	UpdatedBy int     `gorm:"int4" json:"updated_by"`
	UpdatedAt time.Time `gorm:"timestamptz" json:"updated_at"`
	DeletedBy int     `gorm:"int4" json:"deleted_by"`
	DeletedAt time.Time `gorm:"timestamptz" json:"deleted_at"`
}