package handler

import (
	"context"
	"net/http"

	"pathfinder-family/model"
	"pathfinder-family/utils"

	"github.com/gin-gonic/gin"
)

// CalledShots godoc
// @Tags Site
// @Summary Прицельные Удары
// @Description Прицельные Удары
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /variantRules/calledShots [get]
func (h *Handler) CalledShots(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "CalledShots")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "calledShots.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Прицельные Удары",
				Description: "Прицельные Удары",
			},
		})
	}
}

// ArmorAsDamageReduction godoc
// @Tags Site
// @Summary Доспехи как снижение урона
// @Description Доспехи как снижение урона
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /variantRules/armorAsDamageReduction [get]
func (h *Handler) ArmorAsDamageReduction(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "ArmorAsDamageReduction")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "armorAsDamageReduction.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Доспехи как снижение урона",
				Description: "Доспехи как снижение урона",
			},
		})
	}
}

// Duels godoc
// @Tags Site
// @Summary Дуэли
// @Description Дуэли
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /rules/duels [get]
func (h *Handler) Duels(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "Duels")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "duels.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Дуэли",
				Description: "Дуэли",
			},
		})
	}
}
