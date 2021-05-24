package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
	orm "wechat-bot-api/db"
	"wechat-bot-api/utils"
)

type AdminUser struct {
	AdminID      int64   `gorm:"primary_key;auto_increment" json:"id"`
	Name         string  `json:"name"`
	Password     string  `gorm:"size:32" json:"password"`
	Status       int     `gorm:"default:1" json:"status"`
	UpdateTime   string  `gorm:"default:NULL" json:"updateTime"`
	CreateTime   string  `json:"createTime"`
}

type AdminUserView struct {
	AdminID      int64   `json:"id"`
	Name         string  `json:"name"`
	Status       int     `json:"status"`
	UpdateTime   string  `json:"updateTime"`
	CreateTime   string  `json:"createTime"`
}

func (au *AdminUser) CreateAdminUser(permissionIds []int64) (err error) {
	tran := orm.Eloquent.Begin()

	au.CreateTime = utils.GetCurrentTime()
	if err = orm.Eloquent.Table("admin_users").Create(&au).Error; err != nil {
		tran.Rollback()
		return
	}
	fmt.Println(au)
	var adminUsersPermission AdminUsersPermission
	for _, permissionId := range permissionIds {
		adminUsersPermission.AdminId = au.AdminID
		adminUsersPermission.PermissionId = permissionId
		if err = adminUsersPermission.CreateAdminUsersPermission(); err != nil {
			tran.Rollback()
			return
		}
	}

	tran.Commit()
	return
}

func (au *AdminUser) UpdateAdminUserByPermissions(permissionIds []int64) (err error)  {
	tran := orm.Eloquent.Begin()

	var (
		adminUsersPermission AdminUsersPermission
		oldPermissionIds []int64
	)

	aups, err := adminUsersPermission.GetAdminUsersPermissionsByAdminId(au.AdminID)

	if err != nil {
		return err
	}

	for _, aup := range aups{
		oldPermissionIds = append(oldPermissionIds,aup.PermissionId)
	}

	if len(permissionIds) > 0 {
		if len(oldPermissionIds) > 0 {
			// 新取消
			delPermissionIds, err := utils.Difference(oldPermissionIds,permissionIds)
			if err != nil {
				tran.Rollback()
				return err
			}
			if len(delPermissionIds.([]int64)) > 0 {
				if err = adminUsersPermission.DeleteAdminUsersPermissionsByUser(au.AdminID,delPermissionIds.([]int64)); err != nil {
					tran.Rollback()
					return err
				}
			}

			// 新增加
			addPermissionIds, err := utils.Difference(permissionIds,oldPermissionIds)
			if err != nil {
				tran.Rollback()
				return err
			}
			if len(addPermissionIds.([]int64)) > 0 {
				for _, permissionId := range addPermissionIds.([]int64) {
					adminUsersPermission.AdminId = au.AdminID
					adminUsersPermission.PermissionId = permissionId
					if err = adminUsersPermission.CreateAdminUsersPermission(); err != nil {
						tran.Rollback()
						return err
					}
				}
			}
		} else {
			// 全新加
			for _, permissionId := range permissionIds {
				adminUsersPermission.AdminId = au.AdminID
				adminUsersPermission.PermissionId = permissionId
				if err = adminUsersPermission.CreateAdminUsersPermission(); err != nil {
					tran.Rollback()
					return err
				}
			}
		}
	}

	au.UpdateTime = utils.GetCurrentTime()
	if err = orm.Eloquent.Table("admin_users").Omit("create_time").Save(&au).Error; err != nil {
		tran.Rollback()
		return err
	}

	tran.Commit()
	return nil
}

func (au *AdminUser) UpdateAdminUser() error {
	au.UpdateTime = utils.GetCurrentTime()
	if err := orm.Eloquent.Table("admin_users").Omit("create_time").Save(&au).Error; err != nil {
		return err
	}

	return nil
}

func (au * AdminUser) DeleteAdminUsers(ids []int64) error {
	tran := orm.Eloquent.Begin()

	if err := orm.Eloquent.Table("admin_users").Where("admin_id in (?)",ids).Delete(&au).Error; err != nil {
		tran.Rollback()
		return err
	}

	var aup AdminUsersPermission
	if err := aup.DeleteAdminUsersPermissionsByAminId(ids); err != nil {
		tran.Rollback()
		return err
	}

	tran.Commit()
	return nil
}

func (au *AdminUser) GetAdminUserById(id int64) (err error) {
	if err = orm.Eloquent.Table("admin_users").Where("admin_id = ?",id).Take(&au).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			err = nil
		}
		return
	}
	return
}

func (au *AdminUser) GetAdminUserByName(name string) (err error) {
	if err = orm.Eloquent.Table("admin_users").Where("name = ?",name).Take(&au).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			err = nil
		}
		return
	}
	return
}

func (au *AdminUser) GetAdminUser() (err error) {
	if err = orm.Eloquent.Table("admin_users").Where(&au).Take(&au).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			err = nil
		}
		return
	}
	return
}

func (au *AdminUser) GetAdminUserPage(pageSize int, page int, name string, status int) (adminUsers []AdminUserView,total int64,err error)  {
	table := orm.Eloquent.Table("admin_users")
	if name != "" {
		table = table.Where("name = ?",name)
	}
	if status != -1 {
		table = table.Where("status = ?",status)
	}
	if err = table.Offset((page -1) * pageSize).Limit(pageSize).Order("create_time desc").Find(&adminUsers).Error; err != nil {
		if err == gorm.ErrRecordNotFound  {
			err = nil
		}
	}
	table.Count(&total)
	adminUsers = au.formatTime(adminUsers)
	return
}

func (au *AdminUser) formatTime(adminUsers []AdminUserView) []AdminUserView {
	for idx,adminUser := range adminUsers {
		createTime,_ :=  time.Parse("2006-01-02T15:04:05+08:00",adminUser.CreateTime)

		adminUser.CreateTime = createTime.Format("2006-01-02 15:04:05")

		if adminUser.UpdateTime != "" {
			updateTime,_ :=  time.Parse("2006-01-02T15:04:05+08:00",adminUser.UpdateTime)

			adminUser.UpdateTime = updateTime.Format("2006-01-02 15:04:05")
		}
		adminUsers[idx] = adminUser
	}
	return adminUsers
}

