package handler

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"pathfinder-family/model"
	"pathfinder-family/utils"

	"github.com/gin-gonic/gin"
)

// Abilities godoc
// @Tags Site
// @Summary Характеристики
// @Description Характеристики
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /ability [get]
func (h *Handler) Abilities(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "Abilities")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "ability.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Характеристики",
				Description: "Характеристики",
			},
		})
	}
}

// AbilityInfo godoc
// @Tags Site
// @Summary Информация по характеристике
// @Description Информация по характеристике
// @Produce html
// @Param alias path string true "Alias характеристики"
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /ability/{alias} [get]
func (h *Handler) AbilityInfo(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "AbilityInfo")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		alias := c.Param("alias")

		title, err := h.postgres.GetNameByAlias(c, "ability", alias, true)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				c.Status(http.StatusNotFound)
			} else {
				c.Status(http.StatusInternalServerError)
			}
			return
		}

		c.HTML(http.StatusOK, "defaultPage.html", model.DeafultPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       *title,
				Description: *title,
			},
			Alias:   alias,
			JsPaths: []string{"ability/info"},
		})
	}
}
