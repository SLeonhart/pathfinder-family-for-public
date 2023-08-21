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

// Npcs godoc
// @Tags Site
// @Summary Персонажи ведущего
// @Description Персонажи ведущего
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /npc [get]
func (h *Handler) Npcs(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "Npcs")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "npc.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Персонажи ведущего",
				Description: "Персонажи ведущего",
			},
		})
	}
}

// NpcInfo godoc
// @Tags Site
// @Summary Информация по персонажу ведущего
// @Description Информация по персонажу ведущего
// @Produce html
// @Param alias path string true "Alias персонажа ведущего"
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /npc/{alias} [get]
func (h *Handler) NpcInfo(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "NpcInfo")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		alias := c.Param("alias")

		title, err := h.postgres.GetNameByAlias(c, "class", alias, true)
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
			JsPaths: []string{"npc/info"},
		})
	}
}
