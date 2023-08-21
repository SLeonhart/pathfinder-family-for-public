package handler

import (
	"context"
	"net/http"

	"pathfinder-family/model"
	"pathfinder-family/utils"

	"github.com/gin-gonic/gin"
)

// Dungeons godoc
// @Tags Site
// @Summary Подземелья
// @Description Подземелья
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /dungeons [get]
func (h *Handler) Dungeons(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "Dungeons")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "dungeons.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Подземелья",
				Description: "Подземелья",
			},
		})
	}
}

// Wilderness godoc
// @Tags Site
// @Summary Дикая местность
// @Description Дикая местность
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /wilderness [get]
func (h *Handler) Wilderness(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "Wilderness")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "wilderness.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Дикая местность",
				Description: "Дикая местность",
			},
		})
	}
}

// UrbanAdventures godoc
// @Tags Site
// @Summary Города
// @Description Города
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /urbanAdventures [get]
func (h *Handler) UrbanAdventures(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "UrbanAdventures")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "urbanAdventures.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Города",
				Description: "Города",
			},
		})
	}
}

// Weather godoc
// @Tags Site
// @Summary Погода
// @Description Погода
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /weather [get]
func (h *Handler) Weather(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "Weather")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "weather.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Погода",
				Description: "Погода",
			},
		})
	}
}

// ThePlanes godoc
// @Tags Site
// @Summary Планы
// @Description Планы
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /thePlanes [get]
func (h *Handler) ThePlanes(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "ThePlanes")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "thePlanes.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Планы",
				Description: "Планы",
			},
		})
	}
}

// EnvironmentalRules godoc
// @Tags Site
// @Summary Опасная среда
// @Description Опасная среда
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /environmentalRules [get]
func (h *Handler) EnvironmentalRules(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "EnvironmentalRules")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "environmentalRules.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Опасная среда",
				Description: "Опасная среда",
			},
		})
	}
}

// Hazards godoc
// @Tags Site
// @Summary Опасности
// @Description Опасности
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /environment/hazards [get]
func (h *Handler) Hazards(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "Hazards")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "hazards.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Опасности",
				Description: "Опасности",
			},
		})
	}
}
