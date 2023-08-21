package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"pathfinder-family/model"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/pgtype"
)

// GetPrestigeClasses godoc
// @Tags API
// @Summary Список престиж-классов
// @Description Список престиж-классов
// @Produce json
// @Success 200 {object} []model.Class "OK"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/prestigeClasses [get]
func (h *Handler) GetPrestigeClasses(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "GetPrestigeClasses")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		prestigeClasses, err := h.postgres.GetPrestigeClasses(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}
		for i := range prestigeClasses {
			json.Unmarshal(prestigeClasses[i].BookJson.Bytes, &prestigeClasses[i].Book)
		}
		c.JSON(http.StatusOK, prestigeClasses)
	}
}

// GetPrestigeClassInfo godoc
// @Tags API
// @Summary Информация по престиж-классу
// @Description Информация по престиж-классу
// @Produce json
// @Param alias query string true "Alias престиж-класса"
// @Success 200 {object} model.PrestigeClassInfo "OK"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/prestigeClassInfo [get]
func (h *Handler) GetPrestigeClassInfo(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "GetPrestigeClassInfo")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		alias := c.Query("alias")

		prestigeClassInfo, err := h.postgres.GetPrestigeClassInfo(c, alias)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}

		if prestigeClassInfo.SkillsJson.Status == pgtype.Present {
			json.Unmarshal(prestigeClassInfo.SkillsJson.Bytes, &prestigeClassInfo.Skills)
		}
		if prestigeClassInfo.BookJson.Status == pgtype.Present {
			json.Unmarshal(prestigeClassInfo.BookJson.Bytes, &prestigeClassInfo.Book)
		}
		if prestigeClassInfo.HelpersJson.Status == pgtype.Present {
			json.Unmarshal(prestigeClassInfo.HelpersJson.Bytes, &prestigeClassInfo.Helpers)
		}

		c.JSON(http.StatusOK, prestigeClassInfo)
	}
}
