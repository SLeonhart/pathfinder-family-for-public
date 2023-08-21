package handler

import (
	"context"
	"net/http"
	"net/url"

	"pathfinder-family/model"
	"pathfinder-family/utils"

	"github.com/gin-gonic/gin"
)

// Login godoc
// @Tags Site
// @Summary Авторизация
// @Description Авторизация
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /user/login [get]
func (h *Handler) Login(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "Login")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		userInfo := utils.GetUserInfoFromGinContext(c)
		if userInfo != nil {
			location := url.URL{Path: "/user/profile"}
			c.Redirect(http.StatusFound, location.RequestURI())
			return
		}

		c.HTML(http.StatusOK, "login.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Авторизация",
				Description: "Авторизация",
			},
		})
	}
}

// Profile godoc
// @Tags Site
// @Summary Профиль
// @Description Профиль
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /user/profile [get]
func (h *Handler) Profile(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "Profile")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		userInfo := utils.GetUserInfoFromGinContext(c)
		if userInfo == nil {
			location := url.URL{Path: "/user/login"}
			c.Redirect(http.StatusFound, location.RequestURI())
			return
		}

		c.HTML(http.StatusOK, "profile.html", model.ProfilePage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Профиль",
				Description: "Профиль",
			},
			Login: userInfo.Login,
			Email: userInfo.Email,
			Role:  userInfo.Role,
		})
	}
}

// Favourites godoc
// @Tags Site
// @Summary Избранное
// @Description Избранное
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /user/favourites [get]
func (h *Handler) Favourites(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "Favourites")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		userInfo := utils.GetUserInfoFromGinContext(c)
		if userInfo == nil {
			location := url.URL{Path: "/user/login"}
			c.Redirect(http.StatusFound, location.RequestURI())
			return
		}

		c.HTML(http.StatusOK, "favourites.html", model.ProfilePage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Избранное",
				Description: "Избранное",
			},
			Login: userInfo.Login,
			Email: userInfo.Email,
			Role:  userInfo.Role,
		})
	}
}

// Logout godoc
// @Tags Site
// @Summary Разлогин
// @Description Разлогин
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /user/logout [get]
func (h *Handler) Logout(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "Logout")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		userInfo := utils.GetUserInfoFromGinContext(c)
		if userInfo != nil {
			h.postgres.Logout(c, userInfo.Id)
			utils.RemoveAuthCookie(c, h.cfg.App.HostUrl)
			h.inmemory.Delete(model.AuthCacheKey{
				Token: *userInfo.Token,
			}.String())
		}

		location := url.URL{Path: "/user/login"}
		c.Redirect(http.StatusFound, location.RequestURI())
	}
}
