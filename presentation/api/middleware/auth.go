package middleware

import (
	"context"
	"net/http"
	"pathfinder-family/data/cache/cacheInterface"
	"pathfinder-family/data/db/dbInterface"
	"pathfinder-family/model"
	"pathfinder-family/utils"

	"github.com/gin-gonic/gin"
)

func Auth(ctx context.Context, inmemory cacheInterface.IInMemory, postgres dbInterface.IPostgres) func(*gin.Context) {
	return func(c *gin.Context) {
		token := ""
		if cookie, err := c.Request.Cookie(model.TokenCookie); err == nil && cookie != nil {
			token = cookie.Value
		}
		if token == "" {
			c.Next()
			return
		}

		var userData *model.UserData

		сacheKey := model.AuthCacheKey{
			Token: token,
		}
		res := inmemory.Get(сacheKey.String())
		isFromCache := false
		if res != nil {
			var ok bool
			if userData, ok = res.(*model.UserData); ok && userData != nil {
				isFromCache = true
			}
		}
		if !isFromCache {
			var err error
			userData, err = postgres.GetUser(ctx, token)
			if err != nil {
				c.Next()
				return
			}
			go inmemory.Set(сacheKey.String(), userData)
		}

		c.Set(model.UserDataKey, userData)

		c.Next()
	}
}

func RequireAdmin() func(*gin.Context) {
	return func(c *gin.Context) {
		userInfo := utils.GetUserInfoFromGinContext(c)
		if userInfo == nil || userInfo.Role != "Admin" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.ErrorResponse{
				Message: "Unauthorized",
			})
			return
		}

		c.Next()
	}
}

func RequireContent() func(*gin.Context) {
	return func(c *gin.Context) {
		userInfo := utils.GetUserInfoFromGinContext(c)
		if userInfo == nil || (userInfo.Role != "Admin" && userInfo.Role != "Content") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.ErrorResponse{
				Message: "Unauthorized",
			})
			return
		}

		c.Next()
	}
}
