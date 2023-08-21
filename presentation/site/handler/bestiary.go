package handler

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"pathfinder-family/model"
	"pathfinder-family/utils"

	"github.com/gin-gonic/gin"
)

// BestiaryMainDescription godoc
// @Tags Site
// @Summary Общее описание существ
// @Description Общее описание существ
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /bestiary/description [get]
func (h *Handler) BestiaryMainDescription(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "BestiaryMainDescription")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "bestiaryMainDescription.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Описание существ",
				Description: "Описание существ",
			},
		})
	}
}

// MonsterCreation godoc
// @Tags Site
// @Summary Создание чудовища
// @Description Создание чудовища
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /bestiary/appendix/monsterCreation [get]
func (h *Handler) MonsterCreation(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "MonsterCreation")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "bestiaryMonsterCreation.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Создание чудовища",
				Description: "Создание чудовища",
			},
		})
	}
}

// SimpleTemplates godoc
// @Tags Site
// @Summary Простые шаблоны
// @Description Простые шаблоны
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /bestiary/appendix/simpleTemplates [get]
func (h *Handler) SimpleTemplates(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "SimpleTemplates")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "bestiarySimpleTemplates.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Простые шаблоны",
				Description: "Простые шаблоны",
			},
		})
	}
}

// AcquiredTemplates godoc
// @Tags Site
// @Summary Приобретённые шаблоны
// @Description Приобретённые шаблоны
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /bestiary/appendix/acquiredTemplates [get]
func (h *Handler) AcquiredTemplates(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "AcquiredTemplates")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "bestiaryAcquiredTemplates.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Приобретённые шаблоны",
				Description: "Приобретённые шаблоны",
			},
		})
	}
}

// AddingRacialHitDice godoc
// @Tags Site
// @Summary Добавление КЗ народа
// @Description Добавление КЗ народа
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /bestiary/appendix/addingRacialHitDice [get]
func (h *Handler) AddingRacialHitDice(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "AddingRacialHitDice")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "bestiaryAddingRacialHitDice.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Добавление КЗ народа",
				Description: "Добавление КЗ народа",
			},
		})
	}
}

// UniversalMonsterRules godoc
// @Tags Site
// @Summary Общие правила для чудовищ
// @Description Общие правила для чудовищ
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /bestiary/appendix/universalMonsterRules [get]
func (h *Handler) UniversalMonsterRules(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "UniversalMonsterRules")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "bestiaryUniversalMonsterRules.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Общие правила для чудовищ",
				Description: "Общие правила для чудовищ",
			},
		})
	}
}

// CreatureTypes godoc
// @Tags Site
// @Summary Типы существ
// @Description Типы существ
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /bestiary/appendix/creatureTypes [get]
func (h *Handler) CreatureTypes(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "CreatureTypes")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		c.HTML(http.StatusOK, "defaultPage.html", model.DeafultPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Типы существ",
				Description: "Типы существ",
			},
			JsPaths: []string{
				"bestiary/type",
			},
		})
	}
}

// MonstersAsPCs godoc
// @Tags Site
// @Summary Чудовища как персонажи игроков
// @Description Чудовища как персонажи игроков
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /bestiary/appendix/monstersAsPCs [get]
func (h *Handler) MonstersAsPCs(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "MonstersAsPCs")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "bestiaryMonstersAsPCs.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Чудовища как персонажи игроков",
				Description: "Чудовища как персонажи игроков",
			},
		})
	}
}

// MonsterRoles godoc
// @Tags Site
// @Summary Роли чудовищ
// @Description Роли чудовищ
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /bestiary/appendix/monsterRoles [get]
func (h *Handler) MonsterRoles(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "MonsterRoles")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "bestiaryMonsterRoles.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Роли чудовищ",
				Description: "Роли чудовищ",
			},
		})
	}
}

// EncounterTables godoc
// @Tags Site
// @Summary Таблицы случайных сцен
// @Description Таблицы случайных сцен
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /bestiary/appendix/encounterTables [get]
func (h *Handler) EncounterTables(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "EncounterTables")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "bestiaryEncounterTables.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Таблицы случайных сцен",
				Description: "Таблицы случайных сцен",
			},
		})
	}
}

// Bestiary godoc
// @Tags Site
// @Summary Бестиарий
// @Description Бестиарий
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /bestiary [get]
func (h *Handler) Bestiary(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "Bestiary")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "beastList.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Бестиарий",
				Description: "Бестиарий",
			},
		})
	}
}

// BeastInfo godoc
// @Tags Site
// @Summary Информация по монстру
// @Description Информация по монстру
// @Produce html
// @Param alias path string true "Alias монстра"
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /bestiary/beast/{alias} [get]
func (h *Handler) BeastInfo(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "BeastInfo")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		alias := c.Param("alias")

		title, err := h.postgres.GetNameByAlias(c, "beast", alias, true)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				c.Status(http.StatusNotFound)
			} else {
				c.Status(http.StatusInternalServerError)
			}
			return
		}

		c.HTML(http.StatusOK, "defaultPage.html", model.DeafultPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       *title,
				Description: *title,
			},
			H1Class: utils.Ptr("colorPage"),
			Alias:   alias,
			JsPaths: []string{
				"bestiary/main",
				"bestiary/info",
			},
		})
	}
}

// AnimalCompanion godoc
// @Tags Site
// @Summary Верные звери
// @Description Верные звери
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /bestiary/animalCompanion [get]
func (h *Handler) AnimalCompanion(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "AnimalCompanion")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "animalCompanion.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Верные звери",
				Description: "Верные звери",
			},
		})
	}
}

// Familiar godoc
// @Tags Site
// @Summary Фамильяры
// @Description Фамильяры
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /bestiary/familiar [get]
func (h *Handler) Familiar(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "Familiar")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "familiar.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Фамильяры",
				Description: "Фамильяры",
			},
		})
	}
}

// Eidolon godoc
// @Tags Site
// @Summary Эйдолоны
// @Description Эйдолоны
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /bestiary/eidolon [get]
func (h *Handler) Eidolon(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "Eidolon")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "eidolon.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Эйдолоны",
				Description: "Эйдолоны",
			},
		})
	}
}

// Phantom godoc
// @Tags Site
// @Summary Фантомы
// @Description Фантомы
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /bestiary/phantom [get]
func (h *Handler) Phantom(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "Phantom")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "phantom.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Фантомы",
				Description: "Фантомы",
			},
		})
	}
}
