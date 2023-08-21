package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"pathfinder-family/model"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/pgtype"
)

// GetNpcs godoc
// @Tags API
// @Summary Список персонажей ведущего
// @Description Список персонажей ведущего
// @Produce json
// @Success 200 {object} []model.NameAlias "OK"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/npcs [get]
func (h *Handler) GetNpcs(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "GetNpcs")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		npcs, err := h.postgres.GetNpcs(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}
		for i := range npcs {
			json.Unmarshal(npcs[i].BookJson.Bytes, &npcs[i].Book)
		}
		c.JSON(http.StatusOK, npcs)
	}
}

// GetNpcInfo godoc
// @Tags API
// @Summary Информация по персонажу ведущего
// @Description Информация по персонажу ведущего
// @Produce json
// @Param alias query string true "Alias персонажа ведущего"
// @Success 200 {object} model.NpcInfo "OK"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/npcInfo [get]
func (h *Handler) GetNpcInfo(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "GetNpcInfo")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		alias := c.Query("alias")

		npcInfo, err := h.postgres.GetNpcInfo(c, alias)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}

		if npcInfo.SkillsJson.Status == pgtype.Present {
			json.Unmarshal(npcInfo.SkillsJson.Bytes, &npcInfo.Skills)
		}
		if npcInfo.BookJson.Status == pgtype.Present {
			json.Unmarshal(npcInfo.BookJson.Bytes, &npcInfo.Book)
		}
		if npcInfo.HelpersJson.Status == pgtype.Present {
			json.Unmarshal(npcInfo.HelpersJson.Bytes, &npcInfo.Helpers)
		}

		c.JSON(http.StatusOK, npcInfo)
	}
}
