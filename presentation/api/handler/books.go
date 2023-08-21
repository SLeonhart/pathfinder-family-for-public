package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"pathfinder-family/model"

	"github.com/gin-gonic/gin"
)

// GetBooks godoc
// @Tags API
// @Summary Список книг
// @Description Список книг
// @Produce json
// @Success 200 {object} []model.Books "OK"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/books [get]
func (h *Handler) GetBooks(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "GetBooks")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		books, err := h.postgres.GetBooks(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, books)
	}
}

// GetBookInfo godoc
// @Tags API
// @Summary Информация по книге
// @Description Информация по книге
// @Produce json
// @Param alias query string true "Alias книги"
// @Success 200 {object} model.BookInfo "OK"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/bookInfo [get]
func (h *Handler) GetBookInfo(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "GetBookInfo")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		alias := c.Query("alias")

		bookInfo, err := h.postgres.GetBookInfo(c, alias)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}

		json.Unmarshal(bookInfo.RacesJson.Bytes, &bookInfo.Races)
		json.Unmarshal(bookInfo.ClassesJson.Bytes, &bookInfo.Classes)
		json.Unmarshal(bookInfo.ArchetypesJson.Bytes, &bookInfo.Archetypes)
		json.Unmarshal(bookInfo.FeatsJson.Bytes, &bookInfo.Feats)
		json.Unmarshal(bookInfo.PrestigeClassesJson.Bytes, &bookInfo.PrestigeClasses)
		json.Unmarshal(bookInfo.TraitsJson.Bytes, &bookInfo.Traits)
		json.Unmarshal(bookInfo.GodsJson.Bytes, &bookInfo.Gods)
		json.Unmarshal(bookInfo.DomainsJson.Bytes, &bookInfo.Domains)
		json.Unmarshal(bookInfo.SubdomainsJson.Bytes, &bookInfo.Subdomains)
		json.Unmarshal(bookInfo.InquisitionsJson.Bytes, &bookInfo.Inquisitions)
		json.Unmarshal(bookInfo.BloodlinesJson.Bytes, &bookInfo.Bloodlines)
		json.Unmarshal(bookInfo.SchoolsJson.Bytes, &bookInfo.Schools)
		json.Unmarshal(bookInfo.SpellsJson.Bytes, &bookInfo.Spells)
		json.Unmarshal(bookInfo.WeaponsJson.Bytes, &bookInfo.Weapons)
		json.Unmarshal(bookInfo.ArmorsJson.Bytes, &bookInfo.Armors)
		json.Unmarshal(bookInfo.EquipmentsJson.Bytes, &bookInfo.Equipments)
		json.Unmarshal(bookInfo.MagicItemAbilitiesJson.Bytes, &bookInfo.MagicItemAbilities)
		json.Unmarshal(bookInfo.MagicItemsJson.Bytes, &bookInfo.MagicItems)
		json.Unmarshal(bookInfo.MonsterAbilitiesJson.Bytes, &bookInfo.MonsterAbilities)
		json.Unmarshal(bookInfo.BeastsJson.Bytes, &bookInfo.Beasts)
		json.Unmarshal(bookInfo.AfflictionsJson.Bytes, &bookInfo.Afflictions)
		json.Unmarshal(bookInfo.TrapsJson.Bytes, &bookInfo.Traps)
		json.Unmarshal(bookInfo.HauntsJson.Bytes, &bookInfo.Haunts)

		c.JSON(http.StatusOK, bookInfo)
	}
}
