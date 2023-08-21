package handler

import (
	"context"
	"net/http"

	"pathfinder-family/model"

	"github.com/gin-gonic/gin"
)

// SendFeedback godoc
// @Tags API
// @Summary Получение всех помощников по группам
// @Description Получение всех помощников по группам
// @Produce json
// @Param body body model.FeedbackRequest true "Модель обратной связи"
// @Success 200 {object} nil "OK"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/sendFeedback [post]
func (h *Handler) SendFeedback(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "SendFeedback")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		request := model.FeedbackRequest{}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}

		if err := h.postgres.SendFeedback(c, request.Theme, request.Email, request.Message); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}
		c.Status(http.StatusOK)
	}
}
