package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"regexp"
	"strconv"
	"wechat-bot-api/configs"
	"wechat-bot-api/models"
	"wechat-bot-api/utils"
)

func AdminUserLogin(ctx *gin.Context) {
	adminUserForm := struct {
		Name     string `json:"username"`
		Password string `json:"password"`
	}{}

	if err := ctx.BindJSON(&adminUserForm); err != nil {
		ctx.JSON(200,utils.Res{Code: 400,Message:"提交数据有误：" + err.Error()})
		return
	}

	if len(adminUserForm.Name) == 0 {
		ctx.JSON(200,utils.Res{Code: 400,Message: "请填写用户名"})
		return
	}

	if len(adminUserForm.Password) == 0 {
		ctx.JSON(200,utils.Res{Code: 400,Message: "请填写密码"})
		return
	}

	var (
		adminUsersModel models.AdminUser
		adminUserPermissionModel models.AdminUsersPermission
		adminPermissionModel models.AdminPermission
	)

	if err := adminUsersModel.GetAdminUserByName(adminUserForm.Name);err != nil || adminUsersModel.AdminID == 0 {
		fmt.Println(adminUsersModel)
		ctx.JSON(200,utils.Res{Code: 400,Message: "用户不存在"})
		return
	}

	passwordMd5 := utils.MD5(adminUserForm.Password + configs.Conf.App.SecretKey)

	if adminUsersModel.Password != passwordMd5 {
		ctx.JSON(200,utils.Res{Code: 400,Message: "密码错误"})
		return
	}

	token, err := utils.Jwt.GenerateToken(adminUsersModel.AdminID,adminUsersModel.Name)

	if err != nil {
		ctx.JSON(200,utils.Res{Code: 400,Message: "登录失败"})
		return
	}

	userPermissions, err := adminUserPermissionModel.GetAdminUsersPermissionsByAdminId(adminUsersModel.AdminID)
	if err != nil {
		ctx.JSON(200,utils.Res{Code: 400,Message: "登录失败"})
		return
	}

	var permissionIds []int64

	for _, userPermission := range userPermissions{
		permissionIds = append(permissionIds,userPermission.PermissionId)
	}

	permissions, err := adminPermissionModel.GetAdminPermissionsByIds(permissionIds)
	if err != nil {
		ctx.JSON(200,utils.Res{Code: 400,Message: "登录失败"})
		return
	}

	ctx.JSON(200,utils.Res{Code: 0,Message: "ok",Data: struct {
		Token string `json:"token"`
		Permissions []models.AdminPermission `json:"permissions"`
	}{Token: token,Permissions: permissions}})
}

func CreateAdminUser(ctx *gin.Context)  {
	adminUserForm := struct {
		Name          string   `json:"name"`
		Password      string   `json:"password"`
		RePassword    string   `json:"repassword"`
		Status        int      `json:"status"`
		PermissionIds []int64  `json:"permissions"`
	}{}

	if err := ctx.BindJSON(&adminUserForm); err != nil {
		ctx.JSON(200,utils.Res{Code: 400,Message:"提交数据有误：" + err.Error()})
		return
	}

	if len(adminUserForm.Name) < 4 {
		ctx.JSON(200,utils.Res{Code: 400,Message: "用户名最少4个字符"})
		return
	}

	if bol,err := regexp.Match("^[a-zA-Z0-9_]{4,}$",[]byte(adminUserForm.Name)); !bol || err != nil {
		ctx.JSON(200,utils.Res{Code: 400,Message: "用户名只能是数字字母下划线组合"})
		return
	}

	if len(adminUserForm.Password) < 6 {
		ctx.JSON(200,utils.Res{Code: 400,Message: "密码最少6个字符"})
		return
	}
	if adminUserForm.Password != adminUserForm.RePassword {
		ctx.JSON(200,utils.Res{Code: 400,Message: "两次密码输入不一致"})
		return
	}

	var _adminUser models.AdminUser
	if _ = _adminUser.GetAdminUserByName(adminUserForm.Name); _adminUser.AdminID != 0 {
		ctx.JSON(200,utils.Res{Code: 400,Message: "用户名已存在"})
		return
	}

	adminUser := models.AdminUser{
		Name: adminUserForm.Name,
		Password: utils.MD5(adminUserForm.Password + configs.Conf.App.SecretKey),
		Status: adminUserForm.Status,
	}

	if err := adminUser.CreateAdminUser(adminUserForm.PermissionIds); err != nil {
		ctx.JSON(200,utils.Res{Code: 400,Message: "新增用户失败"})
		return
	}

	ctx.JSON(200,utils.Res{Code: 0,Message: "新增用户成功"})
}

func UpdateAdminUser(ctx *gin.Context)  {
	adminUserForm := struct {
		Name          string   `json:"name"`
		Status        int      `json:"status"`
		PermissionIds []int64  `json:"permissions"`
	}{}

	adminId,err := strconv.ParseInt(ctx.Param("id"),10,10)
	if err != nil {
		ctx.JSON(200,utils.Res{Code: 400,Message:"获取参数失败：" + err.Error()})
		return
	}

	if err := ctx.BindJSON(&adminUserForm); err != nil {
		ctx.JSON(200,utils.Res{Code: 400,Message:"提交数据有误：" + err.Error()})
		return
	}

	if len(adminUserForm.Name) < 4 {
		ctx.JSON(200,utils.Res{Code: 400,Message: "用户名最少4个字符"})
		return
	}

	if bol,err := regexp.Match("^[a-zA-Z0-9_]{4,}$",[]byte(adminUserForm.Name)); !bol || err != nil {
		ctx.JSON(200,utils.Res{Code: 400,Message: "用户名只能是数字字母下划线组合"})
		return
	}

	var adminUser models.AdminUser

	if err := adminUser.GetAdminUserByName(adminUserForm.Name); err == nil && adminUser.AdminID > 0 && adminUser.AdminID != adminId {
		ctx.JSON(200,utils.Res{Code: 400,Message: "用户名已存在"})
		return
	}

	if err := adminUser.GetAdminUserById(adminId); err != nil || adminUser.AdminID == 0 {
		ctx.JSON(200,utils.Res{Code: 400,Message: "该用户名不存在"})
		return
	}

	adminUser.Name = adminUserForm.Name
	adminUser.Status = adminUserForm.Status

	if err := adminUser.UpdateAdminUserByPermissions(adminUserForm.PermissionIds); err != nil {
		ctx.JSON(200,utils.Res{Code: 400,Message: "修改用户失败"})
		return
	}

	ctx.JSON(200,utils.Res{Code: 0,Message: "修改用户成功"})
}

func UpdateAdminUserPwd(ctx *gin.Context)  {
	pwdForm := struct {
		Password      string   `json:"password"`
		RePassword    string   `json:"repassword"`
	}{}

	adminId,err := strconv.ParseInt(ctx.Param("id"),10,10)
	if err != nil {
		ctx.JSON(200,utils.Res{Code: 400,Message:"获取参数失败：" + err.Error()})
		return
	}

	if err := ctx.BindJSON(&pwdForm); err != nil {
		ctx.JSON(200,utils.Res{Code: 400,Message: "提交数据有误：" + err.Error()})
		return
	}

	if len(pwdForm.Password) < 6 {
		ctx.JSON(200,utils.Res{Code: 400,Message: "密码最少6个字符"})
		return
	}

	if pwdForm.Password != pwdForm.RePassword {
		ctx.JSON(200,utils.Res{Code: 400,Message: "两次密码输入不一致"})
		return
	}

	var adminUser models.AdminUser

	if err := adminUser.GetAdminUserById(adminId); err != nil || adminUser.AdminID == 0 {
		ctx.JSON(200,utils.Res{Code: 400,Message: "该用户名不存在"})
		return
	}

	loginUserId, _ := ctx.Get("loginUserId")
	if adminUser.Name == "admin" && loginUserId != adminUser.AdminID{
		ctx.JSON(200,utils.Res{Code: 400,Message: "非[admin]用户禁止修改[admin]密码"})
		return
	}

	adminUser.Password = utils.MD5(pwdForm.Password + configs.Conf.App.SecretKey)

	if err := adminUser.UpdateAdminUser(); err != nil {
		ctx.JSON(200,utils.Res{Code: 400,Message: "修改密码失败"})
		return
	}

	ctx.JSON(200,utils.Res{Code: 0,Message: "修改密码成功"})
}

func DeleteAdminUsers(ctx *gin.Context)  {
	var ids []int64
	if err := ctx.BindJSON(&ids); err != nil {
		ctx.JSON(200,utils.Res{Code: 400,Message: "提交数据有误：" + err.Error()})
		return
	}

	var adminUser models.AdminUser
	if err := adminUser.DeleteAdminUsers(ids); err != nil {
		ctx.JSON(200,utils.Res{Code: 400,Message: "删除失败"})
		return
	}

	ctx.JSON(200,utils.Res{Code: 0,Message: "删除成功"})
}

func GetAdminUser(ctx *gin.Context)  {
	id,_ := strconv.ParseInt(ctx.Param("id"),10,10)

	if id <= 0 {
		ctx.JSON(200,utils.Res{Code: 400,Message: "未找到数据"})
		return
	}

	var (
		adminUser models.AdminUser
		adminUsersPermission models.AdminUsersPermission
	)
	adminUser.AdminID = id

	if err := adminUser.GetAdminUser();err != nil {
		ctx.JSON(200,utils.Res{Code: 400,Message: "获取数据失败"})
		return
	}
	adminUser.Password = ""

	permissions,err := adminUsersPermission.GetAdminUsersPermissionsByAdminId(id)
	if err != nil {
		ctx.JSON(200,utils.Res{Code: 400,Message: "获取数据失败"})
		return
	}

	ctx.JSON(200, utils.Res{Code: 0,Message: "获取成功",Data: struct {
		AdminUser models.AdminUser `json:"adminUser"`
		AdminUsersPermissions []models.AdminUsersPermission `json:"permissions"`
	}{AdminUser: adminUser,AdminUsersPermissions: permissions}})
}

func GetAdminUsersList(ctx *gin.Context)  {

	var (
		adminUserModel models.AdminUser
		adminUser      [] models.AdminUserView
		total          int64
		err            error
	)

	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize","20"))
	page, _ := strconv.Atoi(ctx.DefaultQuery("page","1"))
	name := ctx.DefaultQuery("name","")
	status, _ := strconv.Atoi(ctx.DefaultQuery("status","-1"))
	if adminUser, total, err = adminUserModel.GetAdminUserPage(pageSize,page,name,status); err != nil {
		ctx.JSON(200,utils.Res{Code:400,Message:"获取失败"})
		return
	}

	type adminUsers struct {
		AdminUsers  []models.AdminUserView   `json:"rows"`
		Total       int64                    `json:"total"`
	}

	ctx.JSON(200,utils.Res{Code:0,Message:"获取成功",Data:&adminUsers{AdminUsers:adminUser,Total:total}})
}
