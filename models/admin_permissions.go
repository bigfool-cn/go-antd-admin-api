package models

import (
	"github.com/jinzhu/gorm"
	"time"
	orm "wechat-bot-api/db"
	"wechat-bot-api/utils"
)

type AdminPermission struct {
	PermissionID int64   `gorm:"primary_key;auto_increment" json:"id"`
	Name         string  `gorm:"unique;size:50" json:"name"`
	Code         string  `gorm:"unique" json:"code"`
	ParentID     int64   `gorm:"default:0" json:"parentId"`
	IsHelp       bool    `gorm:"default:0" json:"isHelp"`
	UpdateTime   string  `gorm:"default:NULL" json:"updateTime"`
	CreateTime   string  `json:"createTime"`
}

type AdminPermissionCondition struct {
	Name  string `form:"name"`
}

func (ap *AdminPermission) CreateAdminPermission() (err error) {
	ap.CreateTime = utils.GetCurrentTime()
	if err = orm.Eloquent.Table("admin_permissions").Create(&ap).Error; err != nil {
		return
	}
	return
}

func (ap *AdminPermission) UpdateAdminPermission() (err error) {
	ap.UpdateTime = utils.GetCurrentTime()
	if err = orm.Eloquent.Table("admin_permissions").Omit("create_time").Save(&ap).Error; err != nil {
		return
	}
	return
}

func (ap *AdminPermission) DeleteAdminPermissions(permissionIds []int64) (err error)  {
	tran := orm.Eloquent.Begin()

	if err = orm.Eloquent.Table("admin_permissions").Where("permission_id in (?)",permissionIds).Delete(&ap).Error; err != nil {
		tran.Rollback()
		return
	}

	var aup AdminUsersPermission
	if err = aup.DeleteAdminUsersPermissionsByPermissionId(permissionIds); err != nil {
		tran.Rollback()
		return
	}

	tran.Commit()
	return
}

func (ap *AdminPermission) GetAdminPermission() (err error) {
	if err = orm.Eloquent.Table("admin_permissions").Where(&ap).Take(&ap).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			err = nil
		}
		return
	}
	return
}

func (ap *AdminPermission) GetAdminPermissionByName(name string) (err error) {
	if err = orm.Eloquent.Table("admin_permissions").Where("name = ?",name).Take(&ap).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			err = nil
		}
		return
	}
	return
}

func (ap *AdminPermission) GetAdminPermissionByCode(code string) (err error) {
	if err = orm.Eloquent.Table("admin_permissions").Where("code = ?",code).Take(&ap).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			err = nil
		}
		return
	}
	return
}

func (ap *AdminPermission) GetAdminPermissionsByIds(permissionIds []int64) (adminPermissions []AdminPermission,err error) {
	if err = orm.Eloquent.Table("admin_permissions").Where("permission_id in (?)",permissionIds).Find(&adminPermissions).Error; err != nil {
		return
	}
	return
}

func (ap *AdminPermission) GetAdminPermissions(condition AdminPermissionCondition) (adminPermissions []AdminPermission,err error)  {
	table := orm.Eloquent.Table("admin_permissions")
	if condition.Name != "" {
		table = table.Where("name = ?",condition.Name)
	}

	if err = table.Order("create_time desc").Find(&adminPermissions).Error; err != nil {
		if err == gorm.ErrRecordNotFound  {
			err = nil
		}
	}

	adminPermissions = ap.formatTime(adminPermissions)
	return
}

func (ap *AdminPermission) formatTime(adminPermissions []AdminPermission) []AdminPermission {
	for idx,adminPermission := range adminPermissions {
		createTime,_ :=  time.Parse("2006-01-02T15:04:05+08:00",adminPermission.CreateTime)

		adminPermission.CreateTime = createTime.Format("2006-01-02 15:04:05")

		if adminPermission.UpdateTime != "" {
			updateTime,_ :=  time.Parse("2006-01-02T15:04:05+08:00",adminPermission.UpdateTime)

			adminPermission.UpdateTime = updateTime.Format("2006-01-02 15:04:05")
		}
		adminPermissions[idx] = adminPermission
	}
	return adminPermissions
}
