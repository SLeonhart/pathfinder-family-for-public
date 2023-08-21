package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"pathfinder-family/model"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/pgtype"
)

// GetRaces godoc
// @Tags API
// @Summary Список народов
// @Description Список народов
// @Produce json
// @Success 200 {object} []model.Race "OK"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/races [get]
func (h *Handler) GetRaces(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "GetRaces")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		races, err := h.postgres.GetRaces(c)
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

// GetRaceInfo godoc
// @Tags API
// @Summary Информация по народу
// @Description Информация по народу
// @Produce json
// @Param alias query string true "Alias народа"
// @Success 200 {object} model.RaceInfo "OK"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/raceInfo [get]
func (h *Handler) GetRaceInfo(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "GetRaceInfo")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		alias := c.Query("alias")

		raceInfo, err := h.postgres.GetRaceInfo(c, alias)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}

		if raceInfo.NamesJson.Status == pgtype.Present {
			json.Unmarshal(raceInfo.NamesJson.Bytes, &raceInfo.Names)
		}
		if raceInfo.BaseRaceTraitsJson.Status == pgtype.Present {
			json.Unmarshal(raceInfo.BaseRaceTraitsJson.Bytes, &raceInfo.BaseRaceTraits)
		}
		if raceInfo.AlterRaceTraitsJson.Status == pgtype.Present {
			json.Unmarshal(raceInfo.AlterRaceTraitsJson.Bytes, &raceInfo.AlterRaceTraits)
		}
		if raceInfo.FavoredClassJson.Status == pgtype.Present {
			json.Unmarshal(raceInfo.FavoredClassJson.Bytes, &raceInfo.FavoredClass)
		}
		if raceInfo.AdventurerClassJson.Status == pgtype.Present {
			json.Unmarshal(raceInfo.AdventurerClassJson.Bytes, &raceInfo.AdventurerClass)
		}
		if raceInfo.BookJson.Status == pgtype.Present {
			json.Unmarshal(raceInfo.BookJson.Bytes, &raceInfo.Book)
		}
		if raceInfo.HelpersJson.Status == pgtype.Present {
			json.Unmarshal(raceInfo.HelpersJson.Bytes, &raceInfo.Helpers)
		}

		c.JSON(http.StatusOK, raceInfo)
	}
}
