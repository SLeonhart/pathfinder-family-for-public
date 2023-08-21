package handler

import (
	"context"
	"net/http"

	"pathfinder-family/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetNews godoc
// @Tags API
// @Summary Получение актуальных всех новостей
// @Description Получение актуальных всех новостей
// @Produce json
// @Param offset query int false "offset результата. По умолчанию 0"
// @Param limit query int false "limit результата. По умолчанию 1000"
// @Success 200 {object} []model.News "OK"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/news [get]
func (h *Handler) GetNews(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "GetNews")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		offsetStr := c.Query("offset")
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			offset = 0
		}

		limitStr := c.Query("limit")
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			limit = 1000
		}

		news, err := h.postgres.GetNews(c, offset, limit, true)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, news)
	}
}

// GetNewsLast godoc
// @Tags API
// @Summary Получение всех новостей
// @Description Получение всех новостей
// @Produce json
// @Success 200 {object} model.News "OK"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/news/last [get]
func (h *Handler) GetNewsLast(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "GetNewsLast")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		news, err := h.postgres.GetNews(c, 0, 1, false)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, news[0])
	}
}
