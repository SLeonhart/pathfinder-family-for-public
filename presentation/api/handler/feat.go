package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"pathfinder-family/model"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/pgtype"
)

// GetFeats godoc
// @Tags API
// @Summary Список черт
// @Description Список черт
// @Produce json
// @Success 200 {object} []model.Feat "OK"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/feats [get]
func (h *Handler) GetFeats(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "GetFeats")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		feats, err := h.postgres.GetFeats(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}
		for i := range feats {
			if feats[i].BookJson.Status == pgtype.Present {
				json.Unmarshal(feats[i].BookJson.Bytes, &feats[i].Book)
			}
			if feats[i].FeatTypesJson.Status == pgtype.Present {
				json.Unmarshal(feats[i].FeatTypesJson.Bytes, &feats[i].FeatTypes)
			}
		}
		c.JSON(http.StatusOK, feats)
	}
}

// GetFeatInfo godoc
// @Tags API
// @Summary Информация по черте
// @Description Информация по черте
// @Produce json
// @Param alias query string true "Alias черты"
// @Success 200 {object} model.FeatInfo "OK"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/featInfo [get]
func (h *Handler) GetFeatInfo(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "GetFeatInfo")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		alias := c.Query("alias")

		featInfo, err := h.postgres.GetFeatInfo(c, alias)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}

		if featInfo.FeatTypesJson.Status == pgtype.Present {
			json.Unmarshal(featInfo.FeatTypesJson.Bytes, &featInfo.FeatTypes)
		}
		if featInfo.BookJson.Status == pgtype.Present {
			json.Unmarshal(featInfo.BookJson.Bytes, &featInfo.Book)
		}
		if featInfo.HelpersJson.Status == pgtype.Present {
			json.Unmarshal(featInfo.HelpersJson.Bytes, &featInfo.Helpers)
		}

		c.JSON(http.StatusOK, featInfo)
	}
}
