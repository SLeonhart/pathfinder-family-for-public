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

// PrestigeClasses godoc
// @Tags Site
// @Summary Престиж-Классы
// @Description Престиж-Классы
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /class/prestige [get]
func (h *Handler) PrestigeClasses(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "PrestigeClasses")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "prestigeClass.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Престиж-Классы",
				Description: "Престиж-Классы",
			},
		})
	}
}

// PrestigeClassInfo godoc
// @Tags Site
// @Summary Информация по престиж-классу
// @Description Информация по престиж-классу
// @Produce html
// @Param alias path string true "Alias престиж-класса"
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /class/prestige/{alias} [get]
func (h *Handler) PrestigeClassInfo(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "PrestigeClassInfo")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		alias := c.Param("alias")

		title, err := h.postgres.GetNameByAlias(c, "prestige_class", alias, true)
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
			JsPaths: []string{"prestigeClass/info"},
		})
	}
}
