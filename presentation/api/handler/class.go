package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"pathfinder-family/model"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/pgtype"
)

// GetClasses godoc
// @Tags API
// @Summary Список классов
// @Description Список классов
// @Produce json
// @Success 200 {object} []model.Class "OK"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/classes [get]
func (h *Handler) GetClasses(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "GetClasses")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		races, err := h.postgres.GetClasses(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}
		for i := range races {
			json.Unmarshal(races[i].BookJson.Bytes, &races[i].Book)
		}
		c.JSON(http.StatusOK, races)
	}
}

// GetClassInfo godoc
// @Tags API
// @Summary Информация по классу
// @Description Информация по классу
// @Produce json
// @Param alias query string true "Alias народа"
// @Success 200 {object} model.ClassInfo "OK"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/classInfo [get]
func (h *Handler) GetClassInfo(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "GetClassInfo")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		alias := c.Query("alias")

		classInfo, err := h.postgres.GetClassInfo(c, alias)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}

		if classInfo.SkillsJson.Status == pgtype.Present {
			json.Unmarshal(classInfo.SkillsJson.Bytes, &classInfo.Skills)
		}
		if classInfo.ParentClassesJson.Status == pgtype.Present {
			json.Unmarshal(classInfo.ParentClassesJson.Bytes, &classInfo.ParentClasses)
		}
		if classInfo.BookJson.Status == pgtype.Present {
			json.Unmarshal(classInfo.BookJson.Bytes, &classInfo.Book)
		}
		if classInfo.HelpersJson.Status == pgtype.Present {
			json.Unmarshal(classInfo.HelpersJson.Bytes, &classInfo.Helpers)
		}

		c.JSON(http.StatusOK, classInfo)
	}
}
