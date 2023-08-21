package handler

import (
	"context"
	"net/http"

	"pathfinder-family/model"
	"pathfinder-family/utils"

	"github.com/gin-gonic/gin"
)

// WealthAndMoney godoc
// @Tags Site
// @Summary Богатство и деньги
// @Description Богатство и деньги
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /wealthAndMoney [get]
func (h *Handler) WealthAndMoney(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "WealthAndMoney")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "wealthAndMoney.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Богатство и деньги",
				Description: "Богатство и деньги",
			},
		})
	}
}

// WeaponsMainDescription godoc
// @Tags Site
// @Summary Общее описание оружия
// @Description Общее описание оружия
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /weapons/description [get]
func (h *Handler) WeaponsMainDescription(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "WeaponsMainDescription")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "weaponsMainDescription.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Описание оружия",
				Description: "Описание оружия",
			},
		})
	}
}

// Weapons godoc
// @Tags Site
// @Summary Оружие
// @Description Оружие
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /weapons [get]
func (h *Handler) Weapons(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "Weapons")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "weapons.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Оружие",
				Description: "Оружие",
			},
		})
	}
}

// ArmorsMainDescription godoc
// @Tags Site
// @Summary Общее описание доспехов
// @Description Общее описание доспехов
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /armors/description [get]
func (h *Handler) ArmorsMainDescription(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "ArmorsMainDescription")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "armorsMainDescription.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Описание доспехов",
				Description: "Описание доспехов",
			},
		})
	}
}

// Armors godoc
// @Tags Site
// @Summary Доспехи
// @Description Доспехи
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /armors [get]
func (h *Handler) Armors(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "Armors")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "armors.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Доспехи",
				Description: "Доспехи",
			},
		})
	}
}

// SpecialMaterials godoc
// @Tags Site
// @Summary Особые материалы
// @Description Особые материалы
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /specialMaterials [get]
func (h *Handler) SpecialMaterials(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "SpecialMaterials")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "specialMaterials.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Особые материалы",
				Description: "Особые материалы",
			},
		})
	}
}

// GoodsAndServices godoc
// @Tags Site
// @Summary Товары и услуги
// @Description Товары и услуги
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /goodsAndServices [get]
func (h *Handler) GoodsAndServices(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "GoodsAndServices")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "goodsAndServices.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Товары и услуги",
				Description: "Товары и услуги",
			},
		})
	}
}
