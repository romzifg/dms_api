package roles

import (
	"dms_api/modules/roles"
	"time"
)

type RolePermission struct {
	RolePermissionId    int64     	`gorm:"primary_key" json:"role_permission_id"`
	RoleId  			int64    	`gorm:"varchar" json:"role_id"`
	User				roles.Role	`gorm:"foreignKey:RoleId"`	
	Code    			string     	`gorm:"varchar" json:"code"`
	Description    		string     	`gorm:"varchar" json:"description"`
	CreatedBy 			int16     	`gorm:"int4" json:"created_by"`
	UpdatedBy 			int16     	`gorm:"int4" json:"updated_by"`
	DeletedAt 			time.Time 	`gorm:"timestamptz" json:"deleted_at"`
	DeletedBy 			int16     	`gorm:"int4" json:"deleted_by"`
}