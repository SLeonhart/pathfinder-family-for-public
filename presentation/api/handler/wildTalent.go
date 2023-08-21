package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"pathfinder-family/model"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/pgtype"
)

// GetWildTalents godoc
// @Tags API
// @Summary Список диких талантов
// @Description Список диких талантов
// @Produce json
// @Success 200 {object} []model.WildTalent "OK"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/wildTalents [get]
func (h *Handler) GetWildTalents(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "GetWildTalents")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		wildTalents, err := h.postgres.GetWildTalents(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}
		for i := range wildTalents {
			if wildTalents[i].TypeJson.Status == pgtype.Present {
				json.Unmarshal(wildTalents[i].TypeJson.Bytes, &wildTalents[i].Type)
			}
			if wildTalents[i].BookJson.Status == pgtype.Present {
				json.Unmarshal(wildTalents[i].BookJson.Bytes, &wildTalents[i].Book)
			}
		}
		c.JSON(http.StatusOK, wildTalents)
	}
}
