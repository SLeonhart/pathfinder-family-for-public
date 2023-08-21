package handler

import (
	"context"
	"fmt"
	"net/http"

	"pathfinder-family/model"
	"pathfinder-family/utils"

	"github.com/gin-gonic/gin"
)

// Search godoc
// @Tags Site
// @Summary Информация по характеристике
// @Description Информация по характеристике
// @Produce html
// @Param query query string true "Поисковая строка"
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /search/{query} [get]
func (h *Handler) Search(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "Search")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		query := c.Param("query")

		c.HTML(http.StatusOK, "defaultPage.html", model.DeafultPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       fmt.Sprintf("Результаты поиска: %v", query),
				Description: fmt.Sprintf("Результаты поиска: %v", query),
			},
			Alias:   query,
			JsPaths: []string{"search/main"},
		})
	}
}
