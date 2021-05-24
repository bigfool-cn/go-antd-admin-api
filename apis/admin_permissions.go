package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"wechat-bot-api/models"
	"wechat-bot-api/utils"
)

type adminPermissionForm struct {
	Name          string   `json:"name"`
	Code          string   `json:"code"`
	ParentID      int64    `json:"parentId"`
	IsHelp        bool     `json:"isHelp"`
}

func CreateAdminPermission(ctx *gin.Context)  {
	var adminPermissionForm adminPermissionForm

	if err := ctx.BindJSON(&adminPermissionForm); err != nil {
		ctx.JSON(200,utils.Res{Code: 400,Message:"提交数据有误：" + err.Error()})
		return
	}
	if len(adminPermissionForm.Name) == 0 {
		ctx.JSON(200,utils.Res{Code: 400,Message: "请输入名称"})
		return
	}

	if !adminPermissionForm.IsHelp && len(adminPermissionForm.Code) == 0 {
		ctx.JSON(200,utils.Res{Code: 400,Message: "请输入权限标识"})
		return
	}

	var _adminPermission1 models.AdminPermission
	if _ = _adminPermission1.GetAdminPermissionByName(adminPermissionForm.Name); _adminPermission1.PermissionID != 0 {
		ctx.JSON(200,utils.Res{Code: 400,Message: "名称已存在"})
		return
	}

	if !adminPermissionForm.IsHelp {
		var _adminPermission2 models.AdminPermission
		if _ = _adminPermission2.GetAdminPermissionByCode(adminPermissionForm.Code); _adminPermission2.PermissionID != 0 {
			ctx.JSON(200,utils.Res{Code: 400,Message: "权限标识已存在"})
			return
		}
	}

	adminPermission := models.AdminPermission{
		Name: adminPermissionForm.Name,
		Code: adminPermissionForm.Code,
		ParentID: adminPermissionForm.ParentID,
		IsHelp: adminPermissionForm.IsHelp,
	}

	if err := adminPermission.CreateAdminPermission(); err != nil {
		ctx.JSON(200,utils.Res{Code: 400,Message: "新增权限失败"})
		return
	}

	ctx.JSON(200,utils.Res{Code: 0,Message: "新增权限成功"})
}

func UpdateAdminPermission(ctx *gin.Context)  {
	var adminPermissionForm adminPermissionForm

	if err := ctx.BindJSON(&adminPermissionForm); err != nil {
		ctx.JSON(200,utils.Res{Code: 400,Message:"提交数据有误：" + err.Error()})
		return
	}

	permissionId,err := strconv.ParseInt(ctx.Param("id"),10,10)
	if err != nil {
		ctx.JSON(200,utils.Res{Code: 400,Message:"获取参数失败：" + err.Error()})
		return
	}

	if len(adminPermissionForm.Name) == 0 {
		ctx.JSON(200,utils.Res{Code: 400,Message: "请输入名称"})
		return
	}

	if !adminPermissionForm.IsHelp && len(adminPermissionForm.Code) == 0 {
		ctx.JSON(200,utils.Res{Code: 400,Message: "请输入权限标识"})
		return
	}

	var _adminPermission1 models.AdminPermission
	if _ = _adminPermission1.GetAdminPermissionByName(adminPermissionForm.Name); _adminPermission1.PermissionID != 0 && _adminPermission1.PermissionID != permissionId {
		ctx.JSON(200,utils.Res{Code: 400,Message: "名称已存在"})
		return
	}

	if !adminPermissionForm.IsHelp {
		var _adminPermission2 models.AdminPermission
		if _ = _adminPermission2.GetAdminPermissionByCode(adminPermissionForm.Code); _adminPermission2.PermissionID != 0 && _adminPermission2.PermissionID != permissionId {
			ctx.JSON(200,utils.Res{Code: 400,Message: "权限标识已存在"})
			return
		}
		fmt.Println(_adminPermission2.PermissionID)
		fmt.Println(permissionId)
	}

	adminPermission := models.AdminPermission{
		PermissionID: permissionId,
		Name: adminPermissionForm.Name,
		Code: adminPermissionForm.Code,
		ParentID: adminPermissionForm.ParentID,
		IsHelp: adminPermissionForm.IsHelp,
	}

	if err := adminPermission.UpdateAdminPermission(); err != nil {
		fmt.Println(err)
		ctx.JSON(200,utils.Res{Code: 400,Message: "修改权限失败"})
		return
	}

	ctx.JSON(200,utils.Res{Code: 0,Message: "修改权限成功"})
}


func DeleteAdminPermissions(ctx *gin.Context)  {
	var ids []int64

	if err := ctx.BindJSON(&ids); err != nil {
		ctx.JSON(200,utils.Res{Code: 400,Message: "提交数据有误：" + err.Error()})
		return
	}

	var adminPermission models.AdminPermission
	if err := adminPermission.DeleteAdminPermissions(ids); err != nil {
		ctx.JSON(200,utils.Res{Code: 400,Message: "删除成功"})
		return
	}

	ctx.JSON(200,utils.Res{Code: 0,Message: "删除成功："})
}

func GetAdminPermission(ctx *gin.Context)  {
	id,_ := strconv.ParseInt(ctx.Param("id"),10,10)

	if id <= 0 {
		ctx.JSON(200,utils.Res{Code: 400,Message: "未找到数据"})
		return
	}

	var adminPermission models.AdminPermission
	adminPermission.PermissionID = id

	if err := adminPermission.GetAdminPermission();err != nil {
		ctx.JSON(200,utils.Res{Code: 400,Message: "获取数据失败"})
		return
	}
	ctx.JSON(200, utils.Res{Code: 0,Message: "获取成功",Data: adminPermission})
}

func GetAdminPermissions(ctx *gin.Context)  {
	var (
		adminPermissionModel      models.AdminPermission
		adminPermissions          []models.AdminPermission
		adminPermissionCondtion   models.AdminPermissionCondition
		err                       error
	)

	_ = ctx.BindQuery(&adminPermissionCondtion)

	if adminPermissions, err = adminPermissionModel.GetAdminPermissions(adminPermissionCondtion); err != nil {
		fmt.Println(err)
		ctx.JSON(200,utils.Res{Code:400,Message:"获取失败"})
		return
	}

	ctx.JSON(200,utils.Res{Code:0,Message:"获取成功",Data:adminPermissions})
}
