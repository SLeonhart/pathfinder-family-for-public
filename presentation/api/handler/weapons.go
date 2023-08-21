package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"pathfinder-family/model"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/pgtype"
)

// GetWeapons godoc
// @Tags API
// @Summary Список оружия
// @Description Список оружия
// @Produce json
// @Success 200 {object} []model.Weapon "OK"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/weapons [get]
func (h *Handler) GetWeapons(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "GetWeapons")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		weapons, err := h.postgres.GetWeapons(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}
		for i := range weapons {
			if weapons[i].ParentsArray.Status == pgtype.Present {
				weapons[i].Parents = make([]string, 0, len(weapons[i].ParentsArray.Elements))
				for _, element := range weapons[i].ParentsArray.Elements {
					weapons[i].Parents = append(weapons[i].Parents, element.String)
				}
			}
			if weapons[i].ProficientCategoryJson.Status == pgtype.Present {
				json.Unmarshal(weapons[i].ProficientCategoryJson.Bytes, &weapons[i].ProficientCategory)
			}
			if weapons[i].RangeCategoryJson.Status == pgtype.Present {
				json.Unmarshal(weapons[i].RangeCategoryJson.Bytes, &weapons[i].RangeCategory)
			}
			if weapons[i].EncumbranceCategoryJson.Status == pgtype.Present {
				json.Unmarshal(weapons[i].EncumbranceCategoryJson.Bytes, &weapons[i].EncumbranceCategory)
			}
			if weapons[i].BookJson.Status == pgtype.Present {
				json.Unmarshal(weapons[i].BookJson.Bytes, &weapons[i].Book)
			}
			if weapons[i].ChildsJson.Status == pgtype.Present {
				json.Unmarshal(weapons[i].ChildsJson.Bytes, &weapons[i].Childs)
			}
		}
		c.JSON(http.StatusOK, weapons)
	}
}
