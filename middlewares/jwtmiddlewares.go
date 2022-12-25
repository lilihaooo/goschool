package middlewares

import (
	"aschool/conn"
	"aschool/models"
	"aschool/util"
	"github.com/gin-gonic/gin"
	//"crud/app/global"
	"net/http"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {

		uri := context.Request.RequestURI
		if uri == "/admin/login" {
			context.Next()
			return
		}

		tokenStr := context.GetHeader("token")
		if tokenStr == "" {
			context.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "请登陆, 权限不足..."})
			context.Abort()
			return
		}
		token, claims, err := util.ParseToken(tokenStr)
		if err != nil || !token.Valid {
			context.JSON(401, gin.H{"code": 401, "msg": err.Error()})
			context.Abort()
			return
		}

		var user models.User
		conn.DB.First(&user, claims.UserId)
		if user.Id == 0 {
			context.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "请登陆, 权限不足..."})
			context.Abort()
			return
		}
		// 将 claims 中的信息存储在 context 中
		context.Set("claims", *claims)
	}
}
