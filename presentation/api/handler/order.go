package handler

import (
	"context"
	"net/http"

	"pathfinder-family/model"

	"github.com/gin-gonic/gin"
)

// GetOrders godoc
// @Tags API
// @Summary Список орденов
// @Description Список орденов
// @Produce json
// @Success 200 {object} []model.Order "OK"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/orders [get]
func (h *Handler) GetOrders(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "GetOrders")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		orders, err := h.postgres.GetOrders(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, orders)
	}
}

// GetOrderInfo godoc
// @Tags API
// @Summary Информация по ордену
// @Description Информация по ордену
// @Produce json
// @Param alias path string true "Alias ордена"
// @Success 200 {object} model.OrderInfo "OK"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/orderInfo [get]
func (h *Handler) GetOrderInfo(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "GetOrderInfo")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		alias := c.Query("alias")

		orderInfo, err := h.postgres.GetOrderInfo(c, alias)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, orderInfo)
	}
}
