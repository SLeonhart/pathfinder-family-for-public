package handler

import (
	"context"
	"net/http"

	"pathfinder-family/model"

	"github.com/gin-gonic/gin"
)

// GetAbilities godoc
// @Tags API
// @Summary Список характеристик
// @Description Список характеристик
// @Produce json
// @Success 200 {object} []model.Class "OK"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/abilities [get]
func (h *Handler) GetAbilities(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "GetAbilities")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		races, err := h.postgres.GetAbilities(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, races)
	}
}

// GetAbilityInfo godoc
// @Tags API
// @Summary Информация по характеристике
// @Description Информация по характеристике
// @Produce json
// @Param alias query string true "Alias народа"
// @Success 200 {object} model.ClassInfo "OK"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/abilityInfo [get]
func (h *Handler) GetAbilityInfo(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "GetAbilityInfo")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		alias := c.Query("alias")

		classInfo, err := h.postgres.GetAbilityInfo(c, alias)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, classInfo)
	}
}
