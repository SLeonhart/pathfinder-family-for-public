package handler

import (
	"context"
	"net/http"

	"pathfinder-family/model"
	"pathfinder-family/utils"

	"github.com/gin-gonic/gin"
)

// HowCombatWorks godoc
// @Tags Site
// @Summary Как работает система боя
// @Description Как работает система боя
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /howCombatWorks [get]
func (h *Handler) HowCombatWorks(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "HowCombatWorks")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "howCombatWorks.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Как работает система боя",
				Description: "Как работает система боя",
			},
		})
	}
}

// CombatStatistics godoc
// @Tags Site
// @Summary Основные понятия боевой системы
// @Description Основные понятия боевой системы
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /combatStatistics [get]
func (h *Handler) CombatStatistics(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "CombatStatistics")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "combatStatistics.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Основные понятия боевой системы",
				Description: "Основные понятия боевой системы",
			},
		})
	}
}

// ActionsInCombat godoc
// @Tags Site
// @Summary Действия в бою
// @Description Действия в бою
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /actionsInCombat [get]
func (h *Handler) ActionsInCombat(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "ActionsInCombat")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "actionsInCombat.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Действия в бою",
				Description: "Действия в бою",
			},
		})
	}
}

// InjuryAndDeath godoc
// @Tags Site
// @Summary Ранения и смерть
// @Description Ранения и смерть
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /injuryAndDeath [get]
func (h *Handler) InjuryAndDeath(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "InjuryAndDeath")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "injuryAndDeath.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Ранения и смерть",
				Description: "Ранения и смерть",
			},
		})
	}
}

// MovementAndDistance godoc
// @Tags Site
// @Summary Перемещение, позиция и дистанция
// @Description Перемещение, позиция и дистанция
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /movementAndDistance [get]
func (h *Handler) MovementAndDistance(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "MovementAndDistance")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "movementAndDistance.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Перемещение, позиция и дистанция",
				Description: "Перемещение, позиция и дистанция",
			},
		})
	}
}

// BigAndLittleCreatures godoc
// @Tags Site
// @Summary Большие и маленькие существа в бою
// @Description Большие и маленькие существа в бою
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /bigAndLittleCreatures [get]
func (h *Handler) BigAndLittleCreatures(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "BigAndLittleCreatures")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "bigAndLittleCreatures.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Большие и маленькие существа в бою",
				Description: "Большие и маленькие существа в бою",
			},
		})
	}
}

// CombatModifiers godoc
// @Tags Site
// @Summary Боевые модификаторы
// @Description Боевые модификаторы
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /combatModifiers [get]
func (h *Handler) CombatModifiers(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "CombatModifiers")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "combatModifiers.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Боевые модификаторы",
				Description: "Боевые модификаторы",
			},
		})
	}
}

// SpecialAttacks godoc
// @Tags Site
// @Summary Особые атаки
// @Description Особые атаки
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /specialAttacks [get]
func (h *Handler) SpecialAttacks(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "SpecialAttacks")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "specialAttacks.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Особые атаки",
				Description: "Особые атаки",
			},
		})
	}
}

// SpecialInitiativeActions godoc
// @Tags Site
// @Summary Сдвиг инициативы
// @Description Сдвиг инициативы
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /specialInitiativeActions [get]
func (h *Handler) SpecialInitiativeActions(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "SpecialInitiativeActions")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "specialInitiativeActions.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Сдвиг инициативы",
				Description: "Сдвиг инициативы",
			},
		})
	}
}
