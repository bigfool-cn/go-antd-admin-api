package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wechat-bot-api/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			token = ctx.Query("token")
		}

		claims, err := utils.Jwt.ParseToken(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized,utils.Res{Code:401,Message:"token已失效"})
			ctx.Abort()
			return
		}
		ctx.Set("loginUserId",claims.UserInfo.UserId)
		ctx.Set("loginUserName",claims.UserInfo.UserName)
		ctx.Next()
	}
}
