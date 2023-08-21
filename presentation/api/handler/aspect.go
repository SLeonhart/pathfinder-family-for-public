package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"pathfinder-family/model"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/pgtype"
)

// GetAspects godoc
// @Tags API
// @Summary Список аспектов
// @Description Список аспектов
// @Produce json
// @Success 200 {object} []model.Aspect "OK"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/aspects [get]
func (h *Handler) GetAspects(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "GetAspects")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		aspects, err := h.postgres.GetAspects(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}
		for i := range aspects {
			if aspects[i].BookJson.Status == pgtype.Present {
				json.Unmarshal(aspects[i].BookJson.Bytes, &aspects[i].Book)
			}
		}
		c.JSON(http.StatusOK, aspects)
	}
}
