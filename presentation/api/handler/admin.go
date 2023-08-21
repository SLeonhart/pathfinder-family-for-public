package handler

import (
	"context"
	"fmt"
	"net/http"
	"path"

	"pathfinder-family/model"
	"pathfinder-family/utils"
	"time"

	"github.com/gin-gonic/gin"
)

// AddDonate godoc
// @Tags API
// @Summary Добавление доната
// @Description Добавление доната
// @Produce json
// @Param body body model.AddDonateRequest true "Модель добавления доната"
// @Success 200 {object} nil "OK"
// @Success 400 {object} model.ErrorResponse "Bad Request"
// @Success 401 {object} model.ErrorResponse "Unauthorized"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/admin/donate [post]
func (h *Handler) AddDonate(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "AddDonate")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		request := model.AddDonateRequest{}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}
		if len(request.Helper) == 0 || request.Sum <= 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
				Message: "Помощник и сумма - обязательные поля",
			})
			return
		}

		err := h.postgres.AddDonate(c, request)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}

		c.Status(http.StatusOK)
	}
}

// RemovePushTokens godoc
// @Tags API
// @Summary Добавление доната
// @Description Добавление доната
// @Produce json
// @Success 200 {object} nil "OK"
// @Success 400 {object} model.ErrorResponse "Bad Request"
// @Success 401 {object} model.ErrorResponse "Unauthorized"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/admin/push [delete]
func (h *Handler) RemovePushTokens(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "RemovePushTokens")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		_, err := h.removePushTokens(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}

		c.Status(http.StatusOK)
	}
}

// SendPush godoc
// @Tags API
// @Summary Отправка пушей
// @Description Отправка пушей
// @Produce json
// @Param body body model.SendPushRequest true "Модель сообщения пуша"
// @Success 200 {object} nil "OK"
// @Success 400 {object} model.ErrorResponse "Bad Request"
// @Success 401 {object} model.ErrorResponse "Unauthorized"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/admin/push [post]
func (h *Handler) SendPush(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "SendPush")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		request := model.SendPushRequest{}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}

		var tokens []model.PushToken
		if request.Token != nil {
			time.Sleep(2 * time.Second)
			tokens = []model.PushToken{{
				Id:    0,
				Token: *request.Token,
			}}
		} else {
			var err error
			tokens, err = h.removePushTokens(c)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
					Message: err.Error(),
				})
				return
			}
		}

		for _, token := range tokens {
			_, err := h.fcmApi.SendPush(c, request, token.Token)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
					Message: err.Error(),
				})
				return
			}
			// if res.Failure > 0 || res.Success == 0 {
			// 	c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
			// 		Message: "Ошибка при отправке пуша",
			// 	})
			// 	return
			// }
		}

		c.Status(http.StatusOK)
	}
}

// SendTelegram godoc
// @Tags API
// @Summary Отправка сообщения в телеграм
// @Description Отправка сообщения в телеграм
// @Produce json
// @Param body body model.SendTelegramRequest true "Модель сообщения телеграма"
// @Success 200 {object} nil "OK"
// @Success 400 {object} model.ErrorResponse "Bad Request"
// @Success 401 {object} model.ErrorResponse "Unauthorized"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/admin/telegram [post]
func (h *Handler) SendTelegram(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "SendTelegram")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		request := model.SendTelegramRequest{}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}

		res, err := h.telegramApi.Send(c, request)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}
		if !res.Ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: "Ошибка при отправке сообщения в телеграм",
			})
			return
		}

		c.Status(http.StatusOK)
	}
}

// SendVk godoc
// @Tags API
// @Summary Отправка сообщения в ВК
// @Description Отправка сообщения в ВК
// @Produce json
// @Param body body model.SendVkRequest true "Модель сообщения ВК"
// @Success 200 {object} nil "OK"
// @Success 400 {object} model.ErrorResponse "Bad Request"
// @Success 401 {object} model.ErrorResponse "Unauthorized"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/admin/vk [post]
func (h *Handler) SendVk(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "SendVk")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		request := model.SendVkRequest{}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}

		attachments := make([]string, 0, len(request.Images))

		if len(request.Images) > 0 {
			resPhotoServer, err := h.vkApi.GetPhotoServer(c, request)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
					Message: err.Error(),
				})
				return
			}
			if resPhotoServer.Error != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
					Message: resPhotoServer.Error.ErrorMsg,
				})
				return
			}
			for _, image := range request.Images {
				photoByteArray, err := h.vkApi.GetPhoto(c, image)
				if err != nil {
					c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
						Message: err.Error(),
					})
					return
				}

				loadPhotoRes, err := h.vkApi.LoadPhotoIntoServer(c, resPhotoServer.Response.UploadUrl, path.Base(image), photoByteArray)
				if err != nil {
					c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
						Message: err.Error(),
					})
					return
				}
				if loadPhotoRes.Hash == nil || loadPhotoRes.Photo == nil || loadPhotoRes.Server == nil || *loadPhotoRes.Photo == "[]" {
					c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
						Message: fmt.Sprintf("loadPhotoRes has empty values"),
					})
					return
				}

				savePhotoRes, err := h.vkApi.SavePhoto(c, *loadPhotoRes, request.IsTest)
				if err != nil {
					c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
						Message: err.Error(),
					})
					return
				}
				if savePhotoRes.Response[0].Id == nil || savePhotoRes.Response[0].OwnerId == nil {
					c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
						Message: fmt.Sprintf("savePhotoRes has empty values"),
					})
					return
				}

				attachments = append(attachments, fmt.Sprintf("photo%v_%v", *savePhotoRes.Response[0].OwnerId, *savePhotoRes.Response[0].Id))
			}
		}

		res, err := h.vkApi.Send(c, request, attachments)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}
		if res.Error != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: res.Error.ErrorMsg,
			})
			return
		}
		if res.Response == nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: "Ошибка при отправке сообщения в ВК",
			})
			return
		}

		c.Status(http.StatusOK)
	}
}

// AddNews godoc
// @Tags API
// @Summary Добавление новости на сайт
// @Description Добавление новости на сайт
// @Produce json
// @Param body body model.AddNewsRequest true "Модель новости"
// @Success 200 {object} nil "OK"
// @Success 400 {object} model.ErrorResponse "Bad Request"
// @Success 401 {object} model.ErrorResponse "Unauthorized"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/admin/news [post]
func (h *Handler) AddNews(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "AddNews")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		request := model.AddNewsRequest{}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}
		if len(request.Body) == 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
				Message: "Тело новости обязательно",
			})
			return
		}

		if err := h.postgres.AddNews(c, request); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}

		c.Status(http.StatusOK)
	}
}

func (h *Handler) removePushTokens(c *gin.Context) ([]model.PushToken, error) {
	tokens, err := h.postgres.GetPushTokens(c)
	if err != nil {
		return nil, err
	}

	resTokens := make([]model.PushToken, 0, len(tokens))
	i := 0
	for {
		fcmTokens := make([]string, 0, 1000)
		pageTokens := utils.PaginatePushToken(tokens, i, 1000)
		if len(pageTokens) == 0 {
			break
		}

		for _, token := range pageTokens {
			fcmTokens = append(fcmTokens, token.Token)
		}

		checkRes, err := h.fcmApi.CheckTokens(c, fcmTokens)
		if err != nil {
			return nil, err
		}
		if checkRes.Failure > 0 {
			for j := range pageTokens {
				if checkRes.Results[j].Error != nil {
					h.postgres.DeletePushToken(c, tokens[j].Id)
				} else {
					resTokens = append(resTokens, tokens[j])
				}
			}
		}
		i++
	}

	return resTokens, nil
}

// UpsertSearch godoc
// @Tags API
// @Summary Запустить процесс обновления данных в эластике
// @Description Запустить процесс обновления данных в эластике
// @Produce json
// @Success 200 {object} nil "OK"
// @Success 400 {object} model.ErrorResponse "Bad Request"
// @Success 401 {object} model.ErrorResponse "Unauthorized"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/admin/search [patch]
func (h *Handler) UpsertSearch(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "UpsertSearch")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		err := h.elasticSearchService.Upsert(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}

		c.Status(http.StatusOK)
	}
}

// UpdateSearch godoc
// @Tags API
// @Summary Запустить процесс добавления новых данных в эластике
// @Description Запустить процесс добавления новых данных в эластике
// @Produce json
// @Param body body model.ElasticRequest true "Тело сообщения"
// @Success 200 {object} nil "OK"
// @Success 400 {object} model.ErrorResponse "Bad Request"
// @Success 401 {object} model.ErrorResponse "Unauthorized"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/admin/search [post]
func (h *Handler) UpdateSearch(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "UpdateSearch")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		request := model.ElasticRequest{}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}

		err := h.elasticSearchService.Update(c, request.Dt)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}

		c.Status(http.StatusOK)
	}
}

// ClearOldSearch godoc
// @Tags API
// @Summary Запустить процесс чистки старых данных в эластике
// @Description Запустить процесс чистки старых данных в эластике
// @Produce json
// @Param body body model.ElasticRequest true "Тело сообщения"
// @Success 200 {object} nil "OK"
// @Success 400 {object} model.ErrorResponse "Bad Request"
// @Success 401 {object} model.ErrorResponse "Unauthorized"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/admin/search [delete]
func (h *Handler) ClearOldSearch(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "ClearOldSearch")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		request := model.ElasticRequest{}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}

		err := h.elasticSearchService.ClearOld(c, request.Dt)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}

		c.Status(http.StatusOK)
	}
}
