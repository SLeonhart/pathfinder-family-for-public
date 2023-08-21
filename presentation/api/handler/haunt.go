package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"pathfinder-family/model"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/pgtype"
)

// GetHaunts godoc
// @Tags API
// @Summary Список видений
// @Description Список видений
// @Produce json
// @Success 200 {object} []model.Haunt "OK"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/haunts [get]
func (h *Handler) GetHaunts(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "GetHaunts")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		haunts, err := h.postgres.GetHaunts(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}
		for i := range haunts {
			if haunts[i].BookJson.Status == pgtype.Present {
				json.Unmarshal(haunts[i].BookJson.Bytes, &haunts[i].Book)
			}
		}
		c.JSON(http.StatusOK, haunts)
	}
}
