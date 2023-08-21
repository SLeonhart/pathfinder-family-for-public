package handler

import (
	"context"
	"net/http"

	"pathfinder-family/model"
	"pathfinder-family/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/pgtype"
)

// GetCommonPage godoc
// @Tags API
// @Summary Получение статической страницы по id
// @Description Получение статической страницы по id
// @Produce json
// @Param id query int true "Идентификатор страницы"
// @Success 200 {object} model.StaticPage "OK"
// @Success 400 {object} model.ErrorResponse "Bad Request"
// @Router /api/gui/commonPage [get]
func (h *Handler) GetCommonPage(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "GetCommonPage")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		req := c.Query("id")
		id, err := strconv.Atoi(req)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
				Title:   utils.Ptr("id не передан или задан некорректно"),
				Message: err.Error(),
			})
			return
		}

		page, err := h.postgres.GetCommonPage(ctx, id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Title:   utils.Ptr("Ошибка при запросе страницы"),
				Message: err.Error(),
			})
			return
		}

		if page == nil {
			c.AbortWithStatusJSON(http.StatusNotFound, model.ErrorResponse{
				Message: "Страница не найдена",
			})
			return
		}

		if page.ContentJson.Status == pgtype.Present {
			page.Content = utils.Ptr(string(page.ContentJson.Bytes))
		}

		c.JSON(http.StatusOK, page)
	}
}

// UpsertCommonPage godoc
// @Tags API
// @Summary Добавление или изменение статической страницы
// @Description Добавление или изменение статической страницы
// @Produce json
// @Param body body model.StaticPage true "Модель статической страницы"
// @Success 200 {object} model.UpsertResponse "OK"
// @Success 400 {object} model.ErrorResponse "Bad Request"
// @Router /api/gui/commonPage [post]
func (h *Handler) UpsertCommonPage(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "AddOrUpdateCommonPage")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		request := model.StaticPage{}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}

		if request.Id == nil {
			id, err := h.postgres.AddCommonPage(ctx, request)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
					Title:   utils.Ptr("Ошибка при добавлении страницы"),
					Message: err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, model.UpsertResponse{Id: *id})
		} else {
			err := h.postgres.UpdateCommonPage(ctx, request)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
					Title:   utils.Ptr("Ошибка при обновлении страницы"),
					Message: err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, model.UpsertResponse{Id: *request.Id})
		}
	}
}
