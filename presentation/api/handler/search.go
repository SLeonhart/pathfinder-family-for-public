package handler

import (
	"context"
	"net/http"

	"pathfinder-family/model"

	"github.com/gin-gonic/gin"
)

// GetSearchResult godoc
// @Tags API
// @Summary Поиск
// @Description Поиск
// @Produce json
// @Param query query string true "Поисковая строка"
// @Success 200 {object} []model.SearchResult "OK"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/search [get]
func (h *Handler) GetSearchResult(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "GetSearchResult")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		query := c.Query("query")

		searchResult, err := h.elasticSearchService.Get(c, query)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}
		res := make([]model.SearchResult, 0)
		for i := range searchResult {
			res = append(res, model.SearchResult{
				Url:   searchResult[i].Source.Url,
				H1:    searchResult[i].Source.H1,
				Type:  searchResult[i].Source.Type,
				Score: searchResult[i].Score,
			})
		}
		c.JSON(http.StatusOK, res)
	}
}
