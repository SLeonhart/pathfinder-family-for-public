package handler

import (
	"context"
	"net/http"

	"pathfinder-family/model"
	"pathfinder-family/utils"

	"github.com/gin-gonic/gin"
)

// AdventureGenerator godoc
// @Tags Site
// @Summary Генератор случайных приключений
// @Description Генератор случайных приключений
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /generator/adventure [get]
func (h *Handler) AdventureGenerator(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "AdventureGenerator")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		c.HTML(http.StatusOK, "defaultPageGenerator.html", model.DeafultPageGenerator{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Генератор случайных приключений",
				Description: "Генератор случайных приключений",
			},
			JsPaths: []string{"adventure"},
		})
	}
}

// SmallProblemGenerator godoc
// @Tags Site
// @Summary Генератор маленьких проблем
// @Description Генератор маленьких проблем
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /generator/smallProblem [get]
func (h *Handler) SmallProblemGenerator(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "SmallProblemGenerator")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		c.HTML(http.StatusOK, "defaultPageGenerator.html", model.DeafultPageGenerator{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Генератор маленьких проблем",
				Description: "Генератор маленьких проблем",
			},
			JsPaths: []string{"smallProblem"},
		})
	}
}
