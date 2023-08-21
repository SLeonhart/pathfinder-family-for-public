package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"pathfinder-family/model"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/pgtype"
)

// GetMagicItemInfo godoc
// @Tags API
// @Summary Информация по волшебному предмету
// @Description Информация по волшебному предмету
// @Produce json
// @Param alias query string true "Alias волшебного предмета"
// @Success 200 {object} model.MagicItemInfo "OK"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/magicItemInfo [get]
func (h *Handler) GetMagicItemInfo(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "GetMagicItemInfo")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		alias := c.Query("alias")

		magicItemInfo, err := h.postgres.GetMagicItemInfo(c, alias)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}

		if magicItemInfo.SlotJson.Status == pgtype.Present {
			json.Unmarshal(magicItemInfo.SlotJson.Bytes, &magicItemInfo.Slot)
		}
		if magicItemInfo.TypeJson.Status == pgtype.Present {
			json.Unmarshal(magicItemInfo.TypeJson.Bytes, &magicItemInfo.Type)
		}
		if magicItemInfo.BookJson.Status == pgtype.Present {
			json.Unmarshal(magicItemInfo.BookJson.Bytes, &magicItemInfo.Book)
		}
		if magicItemInfo.HelpersJson.Status == pgtype.Present {
			json.Unmarshal(magicItemInfo.HelpersJson.Bytes, &magicItemInfo.Helpers)
		}

		c.JSON(http.StatusOK, magicItemInfo)
	}
}

// GetAllMagicItems godoc
// @Tags API
// @Summary Список волшебных предметов
// @Description Список волшебных предметов
// @Produce json
// @Success 200 {object} []model.MagicItemForList "OK"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/allMagicItems [get]
func (h *Handler) GetAllMagicItems(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "GetAllMagicItems")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		magicItems, err := h.postgres.GetAllMagicItems(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}
		for i := range magicItems {
			if magicItems[i].SlotJson.Status == pgtype.Present {
				json.Unmarshal(magicItems[i].SlotJson.Bytes, &magicItems[i].Slot)
			}
			if magicItems[i].TypeJson.Status == pgtype.Present {
				json.Unmarshal(magicItems[i].TypeJson.Bytes, &magicItems[i].Type)
			}
			if magicItems[i].BookJson.Status == pgtype.Present {
				json.Unmarshal(magicItems[i].BookJson.Bytes, &magicItems[i].Book)
			}
		}
		c.JSON(http.StatusOK, magicItems)
	}
}

// GetMagicItemsByTypes godoc
// @Tags API
// @Summary Список волшебных предметов
// @Description Список волшебных предметов
// @Produce json
// @Param types path string true "Тип волшебных предметов"
// @Success 200 {object} []model.MagicItemInfo "OK"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/magicItems [get]
func (h *Handler) GetMagicItemsByTypes(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "GetMagicItemsByTypes")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		magicItemTypes := c.Query("types")
		if len(magicItemTypes) == 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
				Message: "icorrect types",
			})
			return
		}

		magicItems, err := h.postgres.GetMagicItemsByTypes(c, strings.Split(magicItemTypes, ";"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}
		for i := range magicItems {
			if magicItems[i].SlotJson.Status == pgtype.Present {
				json.Unmarshal(magicItems[i].SlotJson.Bytes, &magicItems[i].Slot)
			}
			if magicItems[i].TypeJson.Status == pgtype.Present {
				json.Unmarshal(magicItems[i].TypeJson.Bytes, &magicItems[i].Type)
			}
			if magicItems[i].BookJson.Status == pgtype.Present {
				json.Unmarshal(magicItems[i].BookJson.Bytes, &magicItems[i].Book)
			}
		}
		c.JSON(http.StatusOK, magicItems)
	}
}

// GetMagicItemAbilitiesByTypes godoc
// @Tags API
// @Summary Список волшебных предметов
// @Description Список волшебных предметов
// @Produce json
// @Param types path string true "Тип волшебных предметов"
// @Success 200 {object} []model.MagicItemInfo "OK"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/magicItemAbilities [get]
func (h *Handler) GetMagicItemAbilitiesByTypes(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "GetMagicItemAbilitiesByTypes")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		magicItemTypes := c.Query("types")
		if len(magicItemTypes) == 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
				Message: "icorrect types",
			})
			return
		}

		magicItemAbilities, err := h.postgres.GetMagicItemAbilitiesByTypes(c, strings.Split(magicItemTypes, ";"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}
		for i := range magicItemAbilities {
			if magicItemAbilities[i].BookJson.Status == pgtype.Present {
				json.Unmarshal(magicItemAbilities[i].BookJson.Bytes, &magicItemAbilities[i].Book)
			}
		}
		c.JSON(http.StatusOK, magicItemAbilities)
	}
}
