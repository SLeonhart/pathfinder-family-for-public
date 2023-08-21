package handler

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"pathfinder-family/model"
	"pathfinder-family/utils"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/pgtype"
)

// UserAuth godoc
// @Tags API
// @Summary Авторизация пользователя
// @Description Авторизация пользователя
// @Produce json
// @Param body body model.UserAuthRequest true "Модель авторизации"
// @Success 200 {object} model.UserAuthResponse "OK"
// @Success 400 {object} model.ErrorResponse "Bad Request"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/user/auth [post]
func (h *Handler) UserAuth(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "UserAuth")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		request := model.UserAuthRequest{}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}
		if len(request.Login) == 0 || len(request.Password) == 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
				Message: "Логин и пароль - обязательные поля",
			})
			return
		}

		userInfo, err := h.postgres.UserAuth(c, request)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
					Message: "Неверный логин и/или Email",
				})
				return
			}
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}

		if userInfo.Token == nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: "Token is null",
			})
			return
		}
		utils.SetAuthCookie(c, h.cfg.App.HostUrl, *userInfo.Token)

		c.JSON(http.StatusOK, model.UserAuthResponse{
			Url: "/user/profile",
		})
	}
}

// UserRegister godoc
// @Tags API
// @Summary Регистрация пользователя
// @Description Регистрация пользователя
// @Produce json
// @Param body body model.UserRegisterRequest true "Модель регистрации"
// @Success 200 {object} model.UserAuthResponse "OK"
// @Success 400 {object} model.ErrorResponse "Bad Request"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/user/register [post]
func (h *Handler) UserRegister(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "UserRegister")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		request := model.UserRegisterRequest{}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}
		if len(request.Login) == 0 || len(request.Password) == 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
				Message: "Логин и пароль - обязательные поля",
			})
			return
		}

		userInfo, err := h.postgres.UserRegister(c, request)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}
		if userInfo == nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
				Message: "Логин занят",
			})
			return
		}

		if userInfo.Token == nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: "Token is null",
			})
			return
		}

		utils.SetAuthCookie(c, h.cfg.App.HostUrl, *userInfo.Token)

		c.JSON(http.StatusOK, model.UserAuthResponse{
			Url: "/user/profile",
		})
	}
}

// UserResetPassword godoc
// @Tags API
// @Summary Сброс пароля пользователя
// @Description Сброс пароля пользователя
// @Produce json
// @Param body body model.UserResetPasswordRequest true "Модель регистрации"
// @Success 200 {object} nil "OK"
// @Success 400 {object} model.ErrorResponse "Bad Request"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/user/resetPassword [post]
func (h *Handler) UserResetPassword(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "UserResetPassword")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		request := model.UserResetPasswordRequest{}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}
		if len(request.Login) == 0 || len(request.Email) == 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
				Message: "Логин и Email - обязательные поля",
			})
			return
		}

		newPassword := utils.GeneratePassword(10)
		userInfo, err := h.postgres.UserResetPassword(c, request, newPassword)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
					Message: "Неверный логин и/или Email",
				})
				return
			}
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}

		err = h.emailService.SendEmail(c, *userInfo.Email, fmt.Sprintf("Новый пароль от логина %v - %v", userInfo.Login, newPassword))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}

		utils.RemoveAuthCookie(c, h.cfg.App.HostUrl)

		c.Status(http.StatusOK)
	}
}

// UserChangeData godoc
// @Tags API
// @Summary Изменить данные пользователя
// @Description Изменить данные пользователя
// @Produce json
// @Param body body model.UserChangeDataRequest true "Модель регистрации"
// @Success 200 {object} nil "OK"
// @Success 400 {object} model.ErrorResponse "Bad Request"
// @Success 401 {object} model.ErrorResponse "Unauthorized"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/user/changeData [post]
func (h *Handler) UserChangeData(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "UserChangeData")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		userInfo := utils.GetUserInfoFromGinContext(c)
		if userInfo == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.ErrorResponse{
				Message: "Unauthorized",
			})
			return
		}

		request := model.UserChangeDataRequest{}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}
		if (request.Email == nil || len(*request.Email) == 0) && (request.Password == nil || len(*request.Password) == 0) {
			c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
				Message: "Для обновления данных необходимо заполнить обновляемые данные",
			})
			return
		}

		userInfo, err := h.postgres.UserChangeData(c, userInfo.Id, request)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}

		utils.SetAuthCookie(c, h.cfg.App.HostUrl, *userInfo.Token)

		c.Status(http.StatusOK)
	}
}

// UserFavourites godoc
// @Tags API
// @Summary Список избранного
// @Description Список избранного
// @Produce json
// @Success 200 {object} []model.Favourites "OK"
// @Success 400 {object} model.ErrorResponse "Bad Request"
// @Success 401 {object} model.ErrorResponse "Unauthorized"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/user/favourites [get]
func (h *Handler) UserFavourites(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "UserFavourites")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		userInfo := utils.GetUserInfoFromGinContext(c)
		if userInfo == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.ErrorResponse{
				Message: "Unauthorized",
			})
			return
		}

		favourites, err := h.postgres.GetUserFavourites(c, userInfo.Id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}

		for i := range favourites {
			if favourites[i].FavouritesJson.Status == pgtype.Present {
				json.Unmarshal(favourites[i].FavouritesJson.Bytes, &favourites[i].Favourites)
			}
		}

		c.JSON(http.StatusOK, favourites)
	}
}

// AddUserFavourites godoc
// @Tags API
// @Summary Добавление списка избранного
// @Description Добавление списка избранного
// @Produce json
// @Param body body model.UserFavouritesRequest true "Модель добавления/удаления списка избранного"
// @Success 200 {object} []model.Favourites "OK"
// @Success 400 {object} model.ErrorResponse "Bad Request"
// @Success 401 {object} model.ErrorResponse "Unauthorized"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/user/favourites [post]
func (h *Handler) AddUserFavourites(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "AddUserFavourites")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		userInfo := utils.GetUserInfoFromGinContext(c)
		if userInfo == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.ErrorResponse{
				Message: "Unauthorized",
			})
			return
		}

		request := model.UserFavouritesRequest{}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}
		if len(request.Guid) == 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
				Message: "Идентификатор списка избранного обязателен",
			})
			return
		}

		err := h.postgres.AddUserFavourites(c, userInfo.Id, request.Guid)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}

		favourites, err := h.postgres.GetUserFavourites(c, userInfo.Id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}

		for i := range favourites {
			if favourites[i].FavouritesJson.Status == pgtype.Present {
				json.Unmarshal(favourites[i].FavouritesJson.Bytes, &favourites[i].Favourites)
			}
		}

		c.JSON(http.StatusOK, favourites)
	}
}

// DeleteUserFavourites godoc
// @Tags API
// @Summary Удаление списка избранного
// @Description Удаление списка избранного
// @Produce json
// @Param body body model.UserFavouritesRequest true "Модель добавления/удаления списка избранного"
// @Success 200 {object} []model.Favourites "OK"
// @Success 400 {object} model.ErrorResponse "Bad Request"
// @Success 401 {object} model.ErrorResponse "Unauthorized"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/user/favourites [delete]
func (h *Handler) DeleteUserFavourites(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "DeleteUserFavourites")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		userInfo := utils.GetUserInfoFromGinContext(c)
		if userInfo == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.ErrorResponse{
				Message: "Unauthorized",
			})
			return
		}

		request := model.UserFavouritesRequest{}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}
		if len(request.Guid) == 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
				Message: "Идентификатор списка избранного обязателен",
			})
			return
		}

		err := h.postgres.DeleteUserFavourites(c, userInfo.Id, request.Guid)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}

		favourites, err := h.postgres.GetUserFavourites(c, userInfo.Id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}

		for i := range favourites {
			if favourites[i].FavouritesJson.Status == pgtype.Present {
				json.Unmarshal(favourites[i].FavouritesJson.Bytes, &favourites[i].Favourites)
			}
		}

		c.JSON(http.StatusOK, favourites)
	}
}

// ChangeUserFavouritesItems godoc
// @Tags API
// @Summary Изменение списка избранного
// @Description Изменение списка избранного
// @Produce json
// @Param body body model.ChangeUserFavouritesItemsRequest true "Модель изменения списка избранного"
// @Success 200 {object} []model.Favourites "OK"
// @Success 400 {object} model.ErrorResponse "Bad Request"
// @Success 401 {object} model.ErrorResponse "Unauthorized"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/user/favourites/items [patch]
func (h *Handler) ChangeUserFavouritesItems(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "ChangeUserFavouritesItems")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		userInfo := utils.GetUserInfoFromGinContext(c)
		if userInfo == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.ErrorResponse{
				Message: "Unauthorized",
			})
			return
		}

		request := model.ChangeUserFavouritesItemsRequest{}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}
		if len(request.PageName) == 0 || len(request.Url) == 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
				Message: "Наименование страницы и ссылка на страницу обязательны",
			})
			return
		}

		err := h.postgres.ChangeUserFavouritesItems(c, userInfo.Id, request)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}

		favourites, err := h.postgres.GetUserFavourites(c, userInfo.Id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}

		for i := range favourites {
			if favourites[i].FavouritesJson.Status == pgtype.Present {
				json.Unmarshal(favourites[i].FavouritesJson.Bytes, &favourites[i].Favourites)
			}
		}

		c.JSON(http.StatusOK, favourites)
	}
}

// RenameUserFavourites godoc
// @Tags API
// @Summary Изменение списка избранного
// @Description Изменение списка избранного
// @Produce json
// @Param body body model.RenameUserFavouritesRequest true "Модель изменения списка избранного"
// @Success 200 {object} []model.Favourites "OK"
// @Success 400 {object} model.ErrorResponse "Bad Request"
// @Success 401 {object} model.ErrorResponse "Unauthorized"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/user/favourites [patch]
func (h *Handler) RenameUserFavourites(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "RenameUserFavourites")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		userInfo := utils.GetUserInfoFromGinContext(c)
		if userInfo == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.ErrorResponse{
				Message: "Unauthorized",
			})
			return
		}

		request := model.RenameUserFavouritesRequest{}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}
		if len(request.Name) == 0 || len(request.Guid) == 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
				Message: "Наименование и идентификатор списка обязательны",
			})
			return
		}

		err := h.postgres.RenameUserFavourites(c, userInfo.Id, request)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}

		favourites, err := h.postgres.GetUserFavourites(c, userInfo.Id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}

		for i := range favourites {
			if favourites[i].FavouritesJson.Status == pgtype.Present {
				json.Unmarshal(favourites[i].FavouritesJson.Bytes, &favourites[i].Favourites)
			}
		}

		c.JSON(http.StatusOK, favourites)
	}
}
