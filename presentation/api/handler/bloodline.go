package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"pathfinder-family/model"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/pgtype"
)

// GetBloodlines godoc
// @Tags API
// @Summary Список наследий
// @Description Список наследий
// @Produce json
// @Param classAlias path string true "Alias класса"
// @Success 200 {object} []model.Bloodline "OK"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/bloodlines [get]
func (h *Handler) GetBloodlines(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "GetBloodlines")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		classAlias := c.Query("classAlias")

		bloodlines, err := h.postgres.GetBloodlines(c, classAlias)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}
		for i := range bloodlines {
			if bloodlines[i].BookJson.Status == pgtype.Present {
				json.Unmarshal(bloodlines[i].BookJson.Bytes, &bloodlines[i].Book)
			}
		}
		c.JSON(http.StatusOK, bloodlines)
	}
}

// GetBloodlineInfo godoc
// @Tags API
// @Summary Информация по наследию
// @Description Информация по наследию
// @Produce json
// @Param classAlias path string true "Alias класса"
// @Param alias path string true "Alias наследия"
// @Success 200 {object} model.BloodlineInfo "OK"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/bloodlineInfo [get]
func (h *Handler) GetBloodlineInfo(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "GetBloodlineInfo")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		alias := c.Query("alias")
		classAlias := c.Query("classAlias")

		bloodlineInfo, err := h.postgres.GetBloodlineInfo(c, classAlias, alias)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}

		if bloodlineInfo.BookJson.Status == pgtype.Present {
			json.Unmarshal(bloodlineInfo.BookJson.Bytes, &bloodlineInfo.Book)
		}

		c.JSON(http.StatusOK, bloodlineInfo)
	}
}
