package middlewares

import (
	"github.com/gin-gonic/gin"
	"wechat-bot-api/utils"
)

func OptMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		loginUserName, isExist := context.Get("loginUserName")
		if (method == "POST" || method == "PUT" || method == "DELETE") && (!isExist || loginUserName != "admin") {
			context.JSON(200,utils.Res{Code:0,Message:"假装成功"})
			context.Abort()
			return
		}

		context.Next()
	}
}
