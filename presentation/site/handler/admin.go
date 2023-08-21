package handler

import (
	"context"
	"net/http"
	"net/url"

	"pathfinder-family/model"
	"pathfinder-family/utils"

	"github.com/gin-gonic/gin"
)

// AdminProfile godoc
// @Tags Site
// @Summary Профиль администратора
// @Description Профиль администратора
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /admin/profile [get]
func (h *Handler) AdminProfile(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "AdminProfile")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		userInfo := utils.GetUserInfoFromGinContext(c)
		if userInfo == nil {
			location := url.URL{Path: "/user/login"}
			c.Redirect(http.StatusFound, location.RequestURI())
			return
		}
		if userInfo.Role != "Admin" {
			location := url.URL{Path: "/user/profile"}
			c.Redirect(http.StatusFound, location.RequestURI())
			return
		}

		c.HTML(http.StatusOK, "adminProfile.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Профиль администратора",
				Description: "Профиль администратора",
			},
		})
	}
}
