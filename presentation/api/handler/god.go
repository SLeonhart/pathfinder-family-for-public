package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"pathfinder-family/model"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/pgtype"
)

// GetGods godoc
// @Tags API
// @Summary Список божеств
// @Description Список божеств
// @Produce json
// @Success 200 {object} []model.God "OK"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/gods [get]
func (h *Handler) GetGods(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "GetGods")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		gods, err := h.postgres.GetGods(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}
		for i := range gods {
			if gods[i].BookJson.Status == pgtype.Present {
				json.Unmarshal(gods[i].BookJson.Bytes, &gods[i].Book)
			}
			if gods[i].GodTypeJson.Status == pgtype.Present {
				json.Unmarshal(gods[i].GodTypeJson.Bytes, &gods[i].GodType)
			}
			if gods[i].DomainsJson.Status == pgtype.Present {
				json.Unmarshal(gods[i].DomainsJson.Bytes, &gods[i].Domains)
			}
		}
		c.JSON(http.StatusOK, gods)
	}
}
