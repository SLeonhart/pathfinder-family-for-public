package handler

import (
	"context"
	"net/http"

	"pathfinder-family/model"
	"pathfinder-family/utils"

	"github.com/gin-gonic/gin"
)

// Afflictions godoc
// @Tags Site
// @Summary Недуги
// @Description Недуги
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /afflictions [get]
func (h *Handler) Afflictions(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "Afflictions")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "afflictions.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Недуги",
				Description: "Недуги",
			},
		})
	}
}
