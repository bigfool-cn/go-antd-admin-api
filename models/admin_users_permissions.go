package models

import (
	"github.com/jinzhu/gorm"
	orm "wechat-bot-api/db"
	"wechat-bot-api/utils"
)

type AdminUsersPermission struct {
	AdminId      int64   `json:"adminId"`
	PermissionId int64   `json:"permissionId"`
	CreateTime   string  `json:"createTime"`
}

func (aup *AdminUsersPermission) CreateAdminUsersPermission() (err error) {
	aup.CreateTime = utils.GetCurrentTime()
	if err = orm.Eloquent.Table("admin_users_permissions").Create(&aup).Error; err != nil {
		return
	}
	return
}

func (aup *AdminUsersPermission) DeleteAdminUsersPermissionsByAminId(adminIds []int64) (err error)  {
	if err = orm.Eloquent.Table("admin_users_permissions").Where("admin_id in (?)",adminIds).Delete(&aup).Error; err != nil {
		return
	}
	return
}

func (aup *AdminUsersPermission) DeleteAdminUsersPermissionsByPermissionId(permissionIds []int64) (err error)  {
	if err = orm.Eloquent.Table("admin_users_permissions").Where("permission_id in (?)",permissionIds).Delete(&aup).Error; err != nil {
		return
	}
	return
}

func (aup *AdminUsersPermission) DeleteAdminUsersPermissionsByUser(adminId int64,permissionIds []int64) (err error)  {
	if err = orm.Eloquent.Table("admin_users_permissions").Where("admin_id = ?",adminId).Where("permission_id in (?)",permissionIds).Delete(&aup).Error; err != nil {
		return
	}
	return
}

func (aup *AdminUsersPermission) GetAdminUsersPermissionsByAdminId(adminId int64) (adminUserPermissions []AdminUsersPermission,err error) {
	if err = orm.Eloquent.Table("admin_users_permissions").Where("admin_id = ?",adminId).Find(&adminUserPermissions).Error; err != nil {
		if err == gorm.ErrRecordNotFound  {
			err = nil
		}
		return
	}
	return
}
