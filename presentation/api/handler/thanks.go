package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"pathfinder-family/model"

	"github.com/gin-gonic/gin"
)

// GetThanks godoc
// @Tags API
// @Summary Получение всех помощников по группам
// @Description Получение всех помощников по группам
// @Produce json
// @Success 200 {object} []model.Thanks "OK"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/thanks [get]
func (h *Handler) GetThanks(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "GetThanks")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		thanks, err := h.postgres.GetThanks(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}
		for i := range thanks {
			json.Unmarshal(thanks[i].ListJson.Bytes, &thanks[i].HelperStat)
		}
		c.JSON(http.StatusOK, thanks)
	}
}
