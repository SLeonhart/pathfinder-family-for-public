package handler

import (
	"context"
	"net/http"

	"pathfinder-family/model"
	"pathfinder-family/utils"

	"github.com/gin-gonic/gin"
)

// Spirits godoc
// @Tags Site
// @Summary Духи
// @Description Духи
// @Produce html
// @Param alias path string true "Alias класса"
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /class/{alias}/spirit [get]
func (h *Handler) Spirits(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "Spirits")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		alias := c.Param("alias")
		jsPath := alias

		if alias == "legendaryMedium" {
			jsPath = "medium"
		}

		if alias != "shaman" && alias != "medium" && alias != "legendaryMedium" {
			c.Status(http.StatusNotFound)
			return
		}

		c.HTML(http.StatusOK, "defaultPage.html", model.DeafultPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Духи",
				Description: "Духи",
			},
			Alias: alias,
			JsPaths: []string{
				"spirits/" + jsPath,
			},
		})
	}
}
