package handler

import (
	"context"
	"net/http"

	"pathfinder-family/model"
	"pathfinder-family/utils"

	"github.com/gin-gonic/gin"
)

// GetGoods godoc
// @Tags API
// @Summary Список товаров
// @Description Список товаров
// @Produce json
// @Success 200 {object} []model.Good "OK"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/goods [get]
func (h *Handler) GetGoods(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "GetGoods")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		var token *string
		userInfo := utils.GetUserInfoFromGinContext(c)
		if userInfo != nil {
			token = userInfo.Token
		}

		goods, err := h.postgres.GetGoods(c, token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, goods)
	}
}

// AddGoodInWaitingList godoc
// @Tags API
// @Summary Добавление товара в список ожидания
// @Description Добавление товара в список ожидания
// @Produce json
// @Param body body model.AddGoodInWaitingListRequest true "Модель обратной связи"
// @Success 200 {object} nil "OK"
// @Success 401 {object} model.ErrorResponse "Unauthorized"
// @Success 400 {object} model.ErrorResponse "Bad Request"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/goods/addWaitingList [post]
func (h *Handler) AddGoodInWaitingList(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "AddGoodInWaitingList")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		userInfo := utils.GetUserInfoFromGinContext(c)
		if userInfo == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.ErrorResponse{
				Message: "Unauthorized",
			})
			return
		}

		request := model.AddGoodInWaitingListRequest{}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}

		err := h.postgres.AddGoodInWaitingList(c, userInfo.Id, request.Id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}
		c.Status(http.StatusOK)
	}
}
