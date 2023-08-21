package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"pathfinder-family/model"

	"github.com/gin-gonic/gin"
)

// GetTranslations godoc
// @Tags API
// @Summary Получение переводов в группе
// @Description Получение переводов в группе
// @Produce json
// @Param alias query string true "Тип переводов"
// @Success 200 {object} []model.Translation "OK"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/translations [get]
func (h *Handler) GetTranslations(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "GetTranslations")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		alias := c.Query("alias")

		translations, err := h.postgres.GetTranslations(c, alias)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}
		for i := range translations {
			json.Unmarshal(translations[i].ItemsJson.Bytes, &translations[i].Items)
		}

		c.JSON(http.StatusOK, translations)
	}
}
