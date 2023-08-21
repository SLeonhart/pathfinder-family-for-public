package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"pathfinder-family/model"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/pgtype"
)

// GetSpellSchools godoc
// @Tags API
// @Summary Список школ магии
// @Description Список школ магии
// @Produce json
// @Success 200 {object} []model.School "OK"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/spellSchools [get]
func (h *Handler) GetSpellSchools(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "GetSpellSchools")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		schools, err := h.postgres.GetSpellSchools(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}
		for i := range schools {
			if schools[i].BookJson.Status == pgtype.Present {
				json.Unmarshal(schools[i].BookJson.Bytes, &schools[i].Book)
			}
			if schools[i].TypeJson.Status == pgtype.Present {
				json.Unmarshal(schools[i].TypeJson.Bytes, &schools[i].Type)
			}
		}
		c.JSON(http.StatusOK, schools)
	}
}

// GetSpellSchoolInfo godoc
// @Tags API
// @Summary Информация по школе
// @Description Информация по школе
// @Produce json
// @Param alias query string true "Alias школы"
// @Success 200 {object} model.FeatInfo "OK"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/spellSchoolInfo [get]
func (h *Handler) GetSpellSchoolInfo(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "GetSpellSchoolInfo")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		alias := c.Query("alias")

		schoolInfo, err := h.postgres.GetSpellSchoolInfo(c, alias)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}

		if schoolInfo.TypeJson.Status == pgtype.Present {
			json.Unmarshal(schoolInfo.TypeJson.Bytes, &schoolInfo.Type)
		}
		if schoolInfo.BookJson.Status == pgtype.Present {
			json.Unmarshal(schoolInfo.BookJson.Bytes, &schoolInfo.Book)
		}

		c.JSON(http.StatusOK, schoolInfo)
	}
}

// GetSpells godoc
// @Tags API
// @Summary Список заклинаний
// @Description Список заклинаний
// @Produce json
// @Param classAlias path string true "Alias класса"
// @Success 200 {object} []model.Spell "OK"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/spells [get]
func (h *Handler) GetSpells(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "GetSpells")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		classAlias := c.Query("classAlias")
		var alias *string
		if len(classAlias) > 0 {
			alias = &classAlias
		}

		spells, err := h.postgres.GetSpells(c, alias)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}
		for i := range spells {
			if spells[i].BookJson.Status == pgtype.Present {
				json.Unmarshal(spells[i].BookJson.Bytes, &spells[i].Book)
			}
			if spells[i].SchoolsJson.Status == pgtype.Present {
				json.Unmarshal(spells[i].SchoolsJson.Bytes, &spells[i].Schools)
			}
			if spells[i].ClassesJson.Status == pgtype.Present {
				json.Unmarshal(spells[i].ClassesJson.Bytes, &spells[i].Classes)
			}
		}
		c.JSON(http.StatusOK, spells)
	}
}

// GetSpellInfo godoc
// @Tags API
// @Summary Информация по заклинанию
// @Description Информация по заклинанию
// @Produce json
// @Param alias query string true "Alias заклинания"
// @Success 200 {object} model.SpellInfo "OK"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/spellInfo [get]
func (h *Handler) GetSpellInfo(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "GetSpellInfo")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		alias := c.Query("alias")

		spellInfo, err := h.postgres.GetSpellInfo(c, alias)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}

		if spellInfo.SchoolJson.Status == pgtype.Present {
			json.Unmarshal(spellInfo.SchoolJson.Bytes, &spellInfo.School)
		}
		if spellInfo.ClassesJson.Status == pgtype.Present {
			json.Unmarshal(spellInfo.ClassesJson.Bytes, &spellInfo.Classes)
		}
		if spellInfo.RacesJson.Status == pgtype.Present {
			json.Unmarshal(spellInfo.RacesJson.Bytes, &spellInfo.Races)
		}
		if spellInfo.BookJson.Status == pgtype.Present {
			json.Unmarshal(spellInfo.BookJson.Bytes, &spellInfo.Book)
		}
		if spellInfo.GodJson.Status == pgtype.Present {
			json.Unmarshal(spellInfo.GodJson.Bytes, &spellInfo.God)
		}
		if spellInfo.HelpersJson.Status == pgtype.Present {
			json.Unmarshal(spellInfo.HelpersJson.Bytes, &spellInfo.Helpers)
		}

		c.JSON(http.StatusOK, spellInfo)
	}
}
