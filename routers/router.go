package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"wechat-bot-api/apis"
	"wechat-bot-api/middlewares"
)

func InitRouter() *gin.Engine  {
	r := gin.New()

	corsConf := cors.DefaultConfig()
	corsConf.AddAllowHeaders("Authorization")
	corsConf.AllowAllOrigins = true
	r.Use(cors.New(corsConf))

	r.POST("/admin/login", apis.AdminUserLogin)

	r.Use(middlewares.AuthMiddleware())
	{
		adminuser := r.Group("adminusers")
		{
			adminuser.GET(":id", apis.GetAdminUser)
			adminuser.GET("", apis.GetAdminUsersList)
			adminuser.POST("", apis.CreateAdminUser)
			adminuser.PUT(":id", apis.UpdateAdminUser)
			adminuser.PUT("pwd/:id", apis.UpdateAdminUserPwd)
			adminuser.DELETE("", apis.DeleteAdminUsers)
		}

		permission := r.Group("adminpermissions")
		{
			permission.GET(":id", apis.GetAdminPermission)
			permission.GET("", apis.GetAdminPermissions)
			permission.POST("", apis.CreateAdminPermission)
			permission.PUT(":id", apis.UpdateAdminPermission)
			permission.DELETE("", apis.DeleteAdminPermissions)
		}
	}

	return r
}

