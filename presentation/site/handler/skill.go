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

// Skills godoc
// @Tags Site
// @Summary Навыки
// @Description Навыки
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /skill [get]
func (h *Handler) Skills(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "Skills")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "skill.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Навыки",
				Description: "Навыки",
			},
		})
	}
}

// SkillInfo godoc
// @Tags Site
// @Summary Информация по навыку
// @Description Информация по навыку
// @Produce html
// @Param alias path string true "Alias навыка"
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /skill/{alias} [get]
func (h *Handler) SkillInfo(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "SkillInfo")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		alias := c.Param("alias")

		title, err := h.postgres.GetNameByAlias(c, "skill", alias, true)
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
			JsPaths: []string{"skill/info"},
		})
	}
}
