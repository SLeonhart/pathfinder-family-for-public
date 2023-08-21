package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"pathfinder-family/model"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/pgtype"
)

// GetAfflictions godoc
// @Tags API
// @Summary Список недугов
// @Description Список недугов
// @Produce json
// @Success 200 {object} []model.Affliction "OK"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/afflictions [get]
func (h *Handler) GetAfflictions(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "GetAfflictions")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		afflictions, err := h.postgres.GetAfflictions(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}
		for i := range afflictions {
			if afflictions[i].BookJson.Status == pgtype.Present {
				json.Unmarshal(afflictions[i].BookJson.Bytes, &afflictions[i].Book)
			}
			if afflictions[i].MainTypeJson.Status == pgtype.Present {
				json.Unmarshal(afflictions[i].MainTypeJson.Bytes, &afflictions[i].MainType)
			}
			if afflictions[i].SecondaryTypesJson.Status == pgtype.Present {
				json.Unmarshal(afflictions[i].SecondaryTypesJson.Bytes, &afflictions[i].SecondaryTypes)
			}
		}
		c.JSON(http.StatusOK, afflictions)
	}
}
