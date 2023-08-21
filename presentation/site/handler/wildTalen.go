package handler

import (
	"context"
	"net/http"

	"pathfinder-family/model"
	"pathfinder-family/utils"

	"github.com/gin-gonic/gin"
)

// WildTalents godoc
// @Tags Site
// @Summary Дикие таланты
// @Description Дикие таланты
// @Produce html
// @Param alias path string true "Alias класса"
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /class/{alias}/wildTalent [get]
func (h *Handler) WildTalents(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "WildTalents")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		alias := c.Param("alias")
		if alias != "kineticist" {
			c.Status(http.StatusNotFound)
			return
		}

		c.HTML(http.StatusOK, "wildTalents.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Дикие таланты",
				Description: "Дикие таланты",
			},
		})
	}
}
