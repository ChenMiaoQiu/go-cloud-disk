package middleware

import (
	"strings"
	"time"

	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
	"github.com/ChenMiaoQiu/go-cloud-disk/utils"
	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// token format Authorization: "Bearer [token]"
		authorization := c.Request.Header.Get("Authorization")
		if authorization == "" {
			c.JSON(200, serializer.NotLogin("Need Token"))
			c.Abort()
			return
		}

		parts := strings.Split(authorization, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(200, serializer.NotLogin("Token format error"))
			c.Abort()
			return
		}

		// parse token
		claims, err := utils.ParseToken(parts[1])
		if err != nil {
			c.JSON(200, serializer.NotLogin("Token error"))
			c.Abort()
			return
		}

		// check if the token has expired
		if time.Now().Unix() > claims.ExpiresAt.Unix() {
			c.JSON(200, serializer.NotLogin("Token expiration"))
			c.Abort()
			return
		}

		c.Set("UserId", claims.UserId)
		c.Set("UserName", claims.UserName)
		c.Set("Status", claims.Status)

		c.Next()
	}
}
