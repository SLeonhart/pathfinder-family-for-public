package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"pathfinder-family/model"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/pgtype"
)

// GetSkills godoc
// @Tags API
// @Summary Список навыков
// @Description Список навыков
// @Produce json
// @Success 200 {object} model.GetSkillsResponse "OK"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/skills [get]
func (h *Handler) GetSkills(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "GetSkills")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		skills, err := h.postgres.GetSkills(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}
		for i := range skills {
			if skills[i].ClassesJson.Status == pgtype.Present {
				json.Unmarshal(skills[i].ClassesJson.Bytes, &skills[i].Classes)
			}
			if skills[i].PrestigeClassesJson.Status == pgtype.Present {
				json.Unmarshal(skills[i].PrestigeClassesJson.Bytes, &skills[i].PrestigeClasses)
			}
		}
		skillsPerLvl, err := h.postgres.GetSkillsPerLvl(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, model.GetSkillsResponse{
			SkillsWithClasses: skills,
			SkillsPerLvl:      skillsPerLvl,
		})
	}
}

// GetSkillInfo godoc
// @Tags API
// @Summary Информация по навыку
// @Description Информация по навыку
// @Produce json
// @Param alias query string true "Alias навыка"
// @Success 200 {object} model.SkillInfo "OK"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/skillInfo [get]
func (h *Handler) GetSkillInfo(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "GetSkillInfo")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		alias := c.Query("alias")

		skillInfo, err := h.postgres.GetSkillInfo(c, alias)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}

		if skillInfo.AbilityJson.Status == pgtype.Present {
			json.Unmarshal(skillInfo.AbilityJson.Bytes, &skillInfo.Ability)
		}

		c.JSON(http.StatusOK, skillInfo)
	}
}
