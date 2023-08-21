package utils

import (
	"fmt"
	"net/url"
	"pathfinder-family/model"

	"github.com/gin-gonic/gin"
)

func getHost(host string) string {
	url, err := url.Parse(host)
	if err != nil {
		return host
	}
	return url.Hostname()
}

func CurrentUrl(c *gin.Context, host string) string {
	return fmt.Sprintf("%v%v", host, c.Request.RequestURI)
}

func SetAuthCookie(c *gin.Context, host string, token string) {
	c.SetCookie(model.TokenCookie, token, 10*365*24*60*60, "/", getHost(host), true, true)
}

func RemoveAuthCookie(c *gin.Context, host string) {
	c.SetCookie(model.TokenCookie, "", -1, "/", getHost(host), true, true)
}

func GetUserInfoFromGinContext(ctx *gin.Context) *model.UserData {
	userInfoData, ok := ctx.Get(model.UserDataKey)
	if !ok {
		return nil
	}

	userInfo, ok := userInfoData.(*model.UserData)
	if !ok {
		return nil
	}

	return userInfo
}
