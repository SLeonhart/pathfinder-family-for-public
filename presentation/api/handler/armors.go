package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"pathfinder-family/model"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/pgtype"
)

// GetArmors godoc
// @Tags API
// @Summary Список доспехов
// @Description Список доспехов
// @Produce json
// @Success 200 {object} []model.Armor "OK"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/armors [get]
func (h *Handler) GetArmors(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "GetArmors")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		armors, err := h.postgres.GetArmors(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}
		for i := range armors {
			if armors[i].TypeJson.Status == pgtype.Present {
				json.Unmarshal(armors[i].TypeJson.Bytes, &armors[i].Type)
			}
			if armors[i].BookJson.Status == pgtype.Present {
				json.Unmarshal(armors[i].BookJson.Bytes, &armors[i].Book)
			}
		}
		c.JSON(http.StatusOK, armors)
	}
}
