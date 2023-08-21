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

// Bloodlines godoc
// @Tags Site
// @Summary Наследия
// @Description Наследия
// @Produce html
// @Param alias path string true "Alias класса"
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /class/{alias}/bloodline [get]
func (h *Handler) Bloodlines(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "Bloodlines")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		alias := c.Param("alias")
		title := "Наследия"
		if alias == "sorcerer" {
			title = "Наследия чародеев"
		} else if alias == "bloodrager" {
			title = "Наследия кровавой ярости"
		} else {
			c.Status(http.StatusNotFound)
			return
		}

		c.HTML(http.StatusOK, "bloodline.html", model.AliasPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       title,
				Description: title,
			},
			Alias: alias,
		})
	}
}

// BloodlineInfo godoc
// @Tags Site
// @Summary Информация по наследию
// @Description Информация по наследию
// @Produce html
// @Param alias path string true "Alias класса"
// @Param bloodlineAlias path string true "Alias наследия"
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /class/{alias}/bloodline/{bloodlineAlias} [get]
func (h *Handler) BloodlineInfo(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "BloodlineInfo")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		alias := c.Param("alias")
		bloodlineAlias := c.Param("bloodlineAlias")

		title, err := h.postgres.GetBloodlineName(c, alias, bloodlineAlias)
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
			Alias:   bloodlineAlias,
			Alias2:  &alias,
			JsPaths: []string{"bloodline/info"},
		})
	}
}
