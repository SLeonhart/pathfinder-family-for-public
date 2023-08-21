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

// Feats godoc
// @Tags Site
// @Summary Черты
// @Description Черты
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /feat [get]
func (h *Handler) Feats(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "Feats")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "feat.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Черты",
				Description: "Черты",
			},
		})
	}
}

// FeatInfo godoc
// @Tags Site
// @Summary Информация по черте
// @Description Информация по черте
// @Produce html
// @Param alias path string true "Alias навыка"
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /feat/{alias} [get]
func (h *Handler) FeatInfo(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "FeatInfo")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		alias := c.Param("alias")

		title, err := h.postgres.GetNameByAlias(c, "feat", alias, true)
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
			JsPaths: []string{"feat/info"},
		})
	}
}
