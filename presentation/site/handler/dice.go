package handler

import (
	"context"
	"net/http"

	"pathfinder-family/model"
	"pathfinder-family/utils"

	"github.com/gin-gonic/gin"
)

// Dice godoc
// @Tags Site
// @Summary Генератор бросков кубиков
// @Description Генератор бросков кубиков
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /dice [get]
func (h *Handler) Dice(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "Dice")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "dice.html", model.CommonPage{
			Page: model.Page{
				Cfg:              h.cfg.App,
				Url:              utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:            "Генератор бросков кубиков",
				Description:      "Генератор бросков кубиков",
				IsNotFixedNavBar: true,
			},
		})
	}
}
