package handler

import (
	"context"
	"net/http"

	"pathfinder-family/model"
	"pathfinder-family/utils"

	"github.com/gin-gonic/gin"
)

// StartingCampaign godoc
// @Tags Site
// @Summary Начало кампании
// @Description Начало кампании
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /startingCampaign [get]
func (h *Handler) StartingCampaign(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "StartingCampaign")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "startingCampaign.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Начало кампании",
				Description: "Начало кампании",
			},
		})
	}
}

// BuildingAdventure godoc
// @Tags Site
// @Summary Создание приключения
// @Description Создание приключения
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /buildingAdventure [get]
func (h *Handler) BuildingAdventure(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "BuildingAdventure")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "buildingAdventure.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Создание приключения",
				Description: "Создание приключения",
			},
		})
	}
}

// PreparingGame godoc
// @Tags Site
// @Summary Подготовка к игре
// @Description Подготовка к игре
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /preparingGame [get]
func (h *Handler) PreparingGame(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "PreparingGame")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "preparingGame.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Подготовка к игре",
				Description: "Подготовка к игре",
			},
		})
	}
}

// DuringGame godoc
// @Tags Site
// @Summary Во время игры
// @Description Во время игры
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /duringGame [get]
func (h *Handler) DuringGame(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "DuringGame")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "duringGame.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Во время игры",
				Description: "Во время игры",
			},
		})
	}
}

// CampaignTips godoc
// @Tags Site
// @Summary Как создавать кампанию
// @Description Как создавать кампанию
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /campaignTips [get]
func (h *Handler) CampaignTips(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "CampaignTips")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "campaignTips.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Как создавать кампанию",
				Description: "Как создавать кампанию",
			},
		})
	}
}

// EndingCampaign godoc
// @Tags Site
// @Summary Завершение кампании
// @Description Завершение кампании
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /endingCampaign [get]
func (h *Handler) EndingCampaign(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "EndingCampaign")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "endingCampaign.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Завершение кампании",
				Description: "Завершение кампании",
			},
		})
	}
}
