package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"pathfinder-family/model"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/pgtype"
)

// GetTraps godoc
// @Tags API
// @Summary Список ловушек
// @Description Список ловушек
// @Produce json
// @Success 200 {object} []model.Trap "OK"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/traps [get]
func (h *Handler) GetTraps(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "GetTraps")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		traps, err := h.postgres.GetTraps(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}
		for i := range traps {
			if traps[i].BookJson.Status == pgtype.Present {
				json.Unmarshal(traps[i].BookJson.Bytes, &traps[i].Book)
			}
		}
		c.JSON(http.StatusOK, traps)
	}
}
