package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"pathfinder-family/model"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/pgtype"
)

// GetShamanSpirits godoc
// @Tags API
// @Summary Список духов
// @Description Список духов
// @Produce json
// @Success 200 {object} []model.ShamanSpiritInfo "OK"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/spirits/shaman [get]
func (h *Handler) GetShamanSpirits(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "GetShamanSpirits")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		spirits, err := h.postgres.GetShamanSpirits(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}
		for i := range spirits {
			if spirits[i].HexesJson.Status == pgtype.Present {
				json.Unmarshal(spirits[i].HexesJson.Bytes, &spirits[i].Hexes)
			}
			if spirits[i].SpiritAbilityJson.Status == pgtype.Present {
				json.Unmarshal(spirits[i].SpiritAbilityJson.Bytes, &spirits[i].SpiritAbility)
			}
			if spirits[i].GreaterSpiritAbilityJson.Status == pgtype.Present {
				json.Unmarshal(spirits[i].GreaterSpiritAbilityJson.Bytes, &spirits[i].GreaterSpiritAbility)
			}
			if spirits[i].TrueSpiritAbilityJson.Status == pgtype.Present {
				json.Unmarshal(spirits[i].TrueSpiritAbilityJson.Bytes, &spirits[i].TrueSpiritAbility)
			}
			if spirits[i].ParentJson.Status == pgtype.Present {
				json.Unmarshal(spirits[i].ParentJson.Bytes, &spirits[i].Parent)
			}
			if spirits[i].SpellsJson.Status == pgtype.Present {
				json.Unmarshal(spirits[i].SpellsJson.Bytes, &spirits[i].Spells)
			}
			if spirits[i].BookJson.Status == pgtype.Present {
				json.Unmarshal(spirits[i].BookJson.Bytes, &spirits[i].Book)
			}
		}
		c.JSON(http.StatusOK, spirits)
	}
}

// GetMediumSpirits godoc
// @Tags API
// @Summary Список духов
// @Description Список духов
// @Produce json
// @Success 200 {object} []model.MediumSpiritInfo "OK"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/spirits/{alias} [get]
func (h *Handler) GetMediumSpirits(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "GetMediumSpirits")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		alias := c.Param("alias")
		if alias != "legendaryMedium" && alias != "medium" {
			c.Status(http.StatusNotFound)
			return
		}

		spirits, err := h.postgres.GetMediumSpirits(c, alias)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}
		for i := range spirits {
			if spirits[i].SpiritPowerBaseJson.Status == pgtype.Present {
				json.Unmarshal(spirits[i].SpiritPowerBaseJson.Bytes, &spirits[i].SpiritPowerBase)
			}
			if spirits[i].SpiritPowerIntermediateJson.Status == pgtype.Present {
				json.Unmarshal(spirits[i].SpiritPowerIntermediateJson.Bytes, &spirits[i].SpiritPowerIntermediate)
			}
			if spirits[i].SpiritPowerGreaterJson.Status == pgtype.Present {
				json.Unmarshal(spirits[i].SpiritPowerGreaterJson.Bytes, &spirits[i].SpiritPowerGreater)
			}
			if spirits[i].SpiritPowerSupremeJson.Status == pgtype.Present {
				json.Unmarshal(spirits[i].SpiritPowerSupremeJson.Bytes, &spirits[i].SpiritPowerSupreme)
			}
			if spirits[i].BookJson.Status == pgtype.Present {
				json.Unmarshal(spirits[i].BookJson.Bytes, &spirits[i].Book)
			}
		}
		c.JSON(http.StatusOK, spirits)
	}
}
