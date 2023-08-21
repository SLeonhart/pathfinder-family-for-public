package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"pathfinder-family/model"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/pgtype"
)

// GetTraits godoc
// @Tags API
// @Summary Список штрихов
// @Description Список штрихов
// @Produce json
// @Success 200 {object} []model.Trait "OK"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/traits [get]
func (h *Handler) GetTraits(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "GetTraits")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		traits, err := h.postgres.GetTraits(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}
		for i := range traits {
			if traits[i].BookJson.Status == pgtype.Present {
				json.Unmarshal(traits[i].BookJson.Bytes, &traits[i].Book)
			}
			if traits[i].TraitTypeJson.Status == pgtype.Present {
				json.Unmarshal(traits[i].TraitTypeJson.Bytes, &traits[i].TraitType)
			}
		}
		c.JSON(http.StatusOK, traits)
	}
}
