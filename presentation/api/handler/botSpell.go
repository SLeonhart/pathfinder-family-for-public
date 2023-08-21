package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"pathfinder-family/model"
	"pathfinder-family/utils"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/pgtype"
)

// GetBotSpells godoc
// @Tags Bot Spell API
// @Summary Список заклинаний
// @Description Список заклинаний
// @Produce json
// @Param id query int false "Id заклинания"
// @Param name query string false "Наименование"
// @Param engName query string false "Name"
// @Param classId query int false "Идентификатор класса"
// @Param alias query string false "Alias заклинания"
// @Param level query int false "Круг заклинания"
// @Param rulebookIds query []int false "Список идентификаторов книг"
// @Param extended query bool false "Флаг выдачи полного контента по заклинанию"
// @Success 200 {object} []model.BotSpellInfo "OK"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/bot/spells [get]
func (h *Handler) GetBotSpells(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*span := jaeger.GetSpan(ctx, "GetBotSpells")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

		var (
			id          *int
			name        *string
			engName     *string
			classId     *int
			alias       *string
			level       *int
			rulebookIds []int
			extended    bool
		)

		idStr := c.Query("id")
		if len(idStr) > 0 {
			res, err := strconv.Atoi(idStr)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
					Title:   utils.Ptr("id задан некорректно"),
					Message: err.Error(),
				})
				return
			}
			id = &res
		}
		nameStr := c.Query("name")
		if len(nameStr) > 0 {
			name = &nameStr
		}
		engNameStr := c.Query("engName")
		if len(engNameStr) > 0 {
			engName = &engNameStr
		}
		classIdStr := c.Query("classId")
		if len(classIdStr) > 0 {
			res, err := strconv.Atoi(classIdStr)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
					Title:   utils.Ptr("classId задан некорректно"),
					Message: err.Error(),
				})
				return
			}
			classId = &res
		}
		aliasStr := c.Query("alias")
		if len(aliasStr) > 0 {
			alias = &aliasStr
		}
		levelStr := c.Query("level")
		if len(levelStr) > 0 {
			res, err := strconv.Atoi(levelStr)
			if err != nil || res < 0 || res > 9 {
				c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
					Title:   utils.Ptr("level задан некорректно"),
					Message: err.Error(),
				})
				return
			}
			level = &res
		}
		rulebookIdsStr := c.Query("rulebookIds")
		if len(rulebookIdsStr) > 0 {
			rulebookIdsArr := strings.Split(rulebookIdsStr, ",")
			for _, rulebookIdStr := range rulebookIdsArr {
				res, err := strconv.Atoi(rulebookIdStr)
				if err != nil {
					c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
						Title:   utils.Ptr("rulebookIds задан некорректно"),
						Message: err.Error(),
					})
					return
				}
				rulebookIds = append(rulebookIds, res)
			}
		}
		extended = c.Query("extended") == "true"

		if name == nil && engName == nil && classId == nil && level == nil && alias == nil && id == nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
				Title:   utils.Ptr("Одно из полей фильтрации обязательно"),
				Message: "Одно из полей фильтрации обязательно",
			})
			return
		}

		spells, err := h.postgres.GetBotSpells(c, id, name, engName, classId, alias, level, rulebookIds)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}

		var res []model.BotSpellInfo
		if extended {
			res = spells
		} else {
			res = make([]model.BotSpellInfo, len(spells))
		}

		for i := range spells {
			if extended {
				if spells[i].SchoolJson.Status == pgtype.Present {
					json.Unmarshal(spells[i].SchoolJson.Bytes, &res[i].School)
				}
				if spells[i].RacesJson.Status == pgtype.Present {
					json.Unmarshal(spells[i].RacesJson.Bytes, &res[i].Races)
				}
				if spells[i].BookJson.Status == pgtype.Present {
					json.Unmarshal(spells[i].BookJson.Bytes, &res[i].Book)
				}
				if spells[i].HelpersJson.Status == pgtype.Present {
					json.Unmarshal(spells[i].HelpersJson.Bytes, &res[i].Helpers)
				}
				if spells[i].GodJson.Status == pgtype.Present {
					json.Unmarshal(spells[i].GodJson.Bytes, &res[i].God)
				}
			} else {
				res[i] = model.BotSpellInfo{
					Id:                         spells[i].Id,
					Alias:                      spells[i].Alias,
					Name:                       spells[i].Name,
					EngName:                    spells[i].EngName,
					ShortDescription:           spells[i].ShortDescription,
					ShortDescriptionComponents: spells[i].ShortDescriptionComponents,
				}
			}
			if spells[i].ClassesJson.Status == pgtype.Present {
				json.Unmarshal(spells[i].ClassesJson.Bytes, &res[i].Classes)
			}
		}
		c.JSON(http.StatusOK, res)
	}
}

// GetBotClasses godoc
// @Tags Bot Spell API
// @Summary Список классов
// @Description Список классов
// @Produce json
// @Param id query int false "Id класса"
// @Param alias query string false "Alias класса"
// @Param magicClass query bool false "Флаг магического класса"
// @Param extended query bool false "Флаг выдачи полного контента по классу"
// @Success 200 {object} []model.BotClassInfo "OK"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/bot/classes [get]
func (h *Handler) GetBotClasses(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*span := jaeger.GetSpan(ctx, "GetBotClasses")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

		var (
			id         *int
			alias      *string
			magicClass *bool
		)

		idStr := c.Query("id")
		if len(idStr) > 0 {
			res, err := strconv.Atoi(idStr)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
					Title:   utils.Ptr("id задан некорректно"),
					Message: err.Error(),
				})
				return
			}
			id = &res
		}
		aliasStr := c.Query("alias")
		if len(aliasStr) > 0 {
			alias = &aliasStr
		}
		magicClassStr := c.Query("magicClass")
		if len(magicClassStr) > 0 {
			magicClass = utils.Ptr(magicClassStr == "true")
		}
		extended := c.Query("extended") == "true"

		classes, err := h.postgres.GetBotClasses(c, id, alias, magicClass)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}

		var res []model.BotClassInfo
		if extended {
			res = classes
		} else {
			res = make([]model.BotClassInfo, len(classes))
		}

		for i := range classes {
			if !extended {
				res[i] = model.BotClassInfo{
					Id:            classes[i].Id,
					Alias:         classes[i].Alias,
					Name:          classes[i].Name,
					Description:   classes[i].Description,
					SpellLevelsDB: classes[i].SpellLevelsDB,
					SpellLevels:   classes[i].SpellLevels,
				}
			}
			if res[i].SpellLevelsDB.Status == pgtype.Present {
				res[i].SpellLevels = make([]int, 0, len(res[i].SpellLevelsDB.Elements))
				for _, element := range res[i].SpellLevelsDB.Elements {
					res[i].SpellLevels = append(res[i].SpellLevels, int(element.Int))
				}
			}
		}
		c.JSON(http.StatusOK, res)
	}
}

// GetBotRulebooks godoc
// @Tags Bot Spell API
// @Summary Список книг
// @Description Список книг
// @Produce json
// @Param withSpells query bool false "Флаг наличия заклинаний"
// @Success 200 {object} []model.BotBook "OK"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/bot/rulebooks [get]
func (h *Handler) GetBotRulebooks(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*span := jaeger.GetSpan(ctx, "GetBotRulebooks")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

		withSpells := c.Query("withSpells") == "true"

		books, err := h.postgres.GetBotBooks(c, withSpells)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, books)
	}
}
