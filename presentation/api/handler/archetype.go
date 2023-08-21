package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"pathfinder-family/model"
	"pathfinder-family/utils"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/pgtype"
)

// GetArchetypes godoc
// @Tags API
// @Summary Список архетипов
// @Description Список архетипов
// @Produce json
// @Param alias query string false "Alias класса"
// @Success 200 {object} []model.ClassWithArchetypes "OK"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/archetypes [get]
func (h *Handler) GetArchetypes(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "GetArchetypes")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		var alias *string
		if len(c.Query("alias")) > 0 {
			alias = utils.Ptr(c.Query("alias"))
		}

		classes, err := h.postgres.GetArchetypes(c, alias)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}
		for i := range classes {
			json.Unmarshal(classes[i].ArchetypesJson.Bytes, &classes[i].Archetypes)
		}
		c.JSON(http.StatusOK, classes)
	}
}

// GetArchetypeInfo godoc
// @Tags API
// @Summary Информация по классу
// @Description Информация по классу
// @Produce json
// @Param alias query string true "Alias народа"
// @Success 200 {object} model.ArchetypeInfo "OK"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/archetypeInfo [get]
func (h *Handler) GetArchetypeInfo(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "GetArchetypeInfo")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		alias := c.Query("alias")

		archetypeInfo, err := h.postgres.GetArchetypeInfo(c, alias)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}

		if archetypeInfo.SkillsJson.Status == pgtype.Present {
			json.Unmarshal(archetypeInfo.SkillsJson.Bytes, &archetypeInfo.Skills)
		}
		if archetypeInfo.ParentClassJson.Status == pgtype.Present {
			json.Unmarshal(archetypeInfo.ParentClassJson.Bytes, &archetypeInfo.ParentClass)
		}
		if archetypeInfo.BookJson.Status == pgtype.Present {
			json.Unmarshal(archetypeInfo.BookJson.Bytes, &archetypeInfo.Book)
		}
		if archetypeInfo.HelpersJson.Status == pgtype.Present {
			json.Unmarshal(archetypeInfo.HelpersJson.Bytes, &archetypeInfo.Helpers)
		}

		c.JSON(http.StatusOK, archetypeInfo)
	}
}
