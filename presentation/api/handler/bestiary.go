package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"pathfinder-family/model"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/pgtype"
)

// GetBeasts godoc
// @Tags API
// @Summary Список монстров
// @Description Список монстров
// @Produce json
// @Success 200 {object} []model.Beast "OK"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/beasts [get]
func (h *Handler) GetBeasts(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "GetBeasts")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		beasts, err := h.postgres.GetBeasts(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}
		for i := range beasts {
			if beasts[i].ChildsArray.Status == pgtype.Present {
				beasts[i].Childs = make([]int, 0, len(beasts[i].ChildsArray.Elements))
				for _, element := range beasts[i].ChildsArray.Elements {
					beasts[i].Childs = append(beasts[i].Childs, int(element.Int))
				}
			}
			if beasts[i].BookJson.Status == pgtype.Present {
				json.Unmarshal(beasts[i].BookJson.Bytes, &beasts[i].Book)
			}
			if beasts[i].ClimateJson.Status == pgtype.Present {
				json.Unmarshal(beasts[i].ClimateJson.Bytes, &beasts[i].Climate)
			}
			if beasts[i].CreatureTypeJson.Status == pgtype.Present {
				json.Unmarshal(beasts[i].CreatureTypeJson.Bytes, &beasts[i].CreatureType)
			}
			if beasts[i].TerrainJson.Status == pgtype.Present {
				json.Unmarshal(beasts[i].TerrainJson.Bytes, &beasts[i].Terrain)
			}
			if beasts[i].RolesJson.Status == pgtype.Present {
				json.Unmarshal(beasts[i].RolesJson.Bytes, &beasts[i].Roles)
			}
		}
		c.JSON(http.StatusOK, beasts)
	}
}

// GetBeastInfo godoc
// @Tags API
// @Summary Информация по монстру
// @Description Информация по монстру
// @Produce json
// @Param alias query string true "Alias монстра"
// @Success 200 {object} model.BeastInfo "OK"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/beastInfo [get]
func (h *Handler) GetBeastInfo(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "GetBeastInfo")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		alias := c.Query("alias")

		beastInfo, err := h.postgres.GetBeastInfo(c, alias)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}

		if beastInfo.BookJson.Status == pgtype.Present {
			json.Unmarshal(beastInfo.BookJson.Bytes, &beastInfo.Book)
		}
		if beastInfo.ClimateJson.Status == pgtype.Present {
			json.Unmarshal(beastInfo.ClimateJson.Bytes, &beastInfo.Climate)
		}
		if beastInfo.CreatureTypeJson.Status == pgtype.Present {
			json.Unmarshal(beastInfo.CreatureTypeJson.Bytes, &beastInfo.CreatureType)
		}
		if beastInfo.TerrainJson.Status == pgtype.Present {
			json.Unmarshal(beastInfo.TerrainJson.Bytes, &beastInfo.Terrain)
		}
		if beastInfo.SizeTypeJson.Status == pgtype.Present {
			json.Unmarshal(beastInfo.SizeTypeJson.Bytes, &beastInfo.SizeType)
		}
		if beastInfo.AnimalCompanionJson.Status == pgtype.Present {
			json.Unmarshal(beastInfo.AnimalCompanionJson.Bytes, &beastInfo.AnimalCompanion)
		}
		if beastInfo.ParentJson.Status == pgtype.Present {
			json.Unmarshal(beastInfo.ParentJson.Bytes, &beastInfo.Parent)
		}
		if beastInfo.RolesJson.Status == pgtype.Present {
			json.Unmarshal(beastInfo.RolesJson.Bytes, &beastInfo.Roles)
		}
		if beastInfo.HelpersJson.Status == pgtype.Present {
			json.Unmarshal(beastInfo.HelpersJson.Bytes, &beastInfo.Helpers)
		}
		if beastInfo.ChildsJson.Status == pgtype.Present {
			json.Unmarshal(beastInfo.ChildsJson.Bytes, &beastInfo.Childs)
		}

		c.JSON(http.StatusOK, beastInfo)
	}
}

// GetMonsterAbilities godoc
// @Tags API
// @Summary Список способностей монстров
// @Description Список способностей монстров
// @Produce json
// @Success 200 {object} []model.MonsterAbility "OK"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/monsterAbilities [get]
func (h *Handler) GetMonsterAbilities(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "GetMonsterAbilities")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		beasts, err := h.postgres.GetMonsterAbilities(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}
		for i := range beasts {
			if beasts[i].BookJson.Status == pgtype.Present {
				json.Unmarshal(beasts[i].BookJson.Bytes, &beasts[i].Book)
			}
		}
		c.JSON(http.StatusOK, beasts)
	}
}

// GetCreatureTypes godoc
// @Tags API
// @Summary Список типов существ
// @Description Список типов существ
// @Produce json
// @Success 200 {object} []model.CreatureType "OK"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/creatureTypes [get]
func (h *Handler) GetCreatureTypes(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "GetCreatureTypes")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		beasts, err := h.postgres.GetCreatureTypes(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, beasts)
	}
}

// GetAnimalCompanions godoc
// @Tags API
// @Summary Список верных зверей
// @Description Список верных зверей
// @Produce json
// @Success 200 {object} []model.AnimalCompanion "OK"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/animalCompanions [get]
func (h *Handler) GetAnimalCompanions(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "GetAnimalCompanions")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		animalCompanions, err := h.postgres.GetAnimalCompanions(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}
		for i := range animalCompanions {
			if animalCompanions[i].BookJson.Status == pgtype.Present {
				json.Unmarshal(animalCompanions[i].BookJson.Bytes, &animalCompanions[i].Book)
			}
			if animalCompanions[i].BeastJson.Status == pgtype.Present {
				json.Unmarshal(animalCompanions[i].BeastJson.Bytes, &animalCompanions[i].Beast)
			}
		}
		c.JSON(http.StatusOK, animalCompanions)
	}
}
