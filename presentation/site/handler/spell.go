package handler

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"net/url"

	"pathfinder-family/model"
	"pathfinder-family/utils"

	"github.com/gin-gonic/gin"
)

// SpellMainDescription godoc
// @Tags Site
// @Summary Общее описание заклинаний
// @Description Общее описание заклинаний
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /rules/spell/description [get]
func (h *Handler) SpellMainDescription(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "SpellMainDescription")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "spellMainDescription.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Описание заклинаний",
				Description: "Описание заклинаний",
			},
		})
	}
}

// CastingSpells godoc
// @Tags Site
// @Summary Сотворение заклинаний
// @Description Сотворение заклинаний
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /rules/spell/castingSpells [get]
func (h *Handler) CastingSpells(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "CastingSpells")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "castingSpells.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Сотворение заклинаний",
				Description: "Сотворение заклинаний",
			},
		})
	}
}

// SpellDescriptions godoc
// @Tags Site
// @Summary Описание заклинаний
// @Description Описание заклинаний
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /rules/spell/spellDescriptions [get]
func (h *Handler) SpellDescriptions(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "SpellDescriptions")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "spellDescriptions.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Описание заклинаний",
				Description: "Описание заклинаний",
			},
		})
	}
}

// SpellSchools godoc
// @Tags Site
// @Summary Школы заклинаний
// @Description Школы заклинаний
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /rules/spell/schools [get]
func (h *Handler) SpellSchools(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "SpellSchools")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "spellSchools.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Школы заклинаний",
				Description: "Школы заклинаний",
			},
		})
	}
}

// ArcaneSpells godoc
// @Tags Site
// @Summary Мистические заклинания
// @Description Мистические заклинания
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /rules/spell/arcaneSpells [get]
func (h *Handler) ArcaneSpells(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "ArcaneSpells")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "arcaneSpells.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Мистические заклинания",
				Description: "Мистические заклинания",
			},
		})
	}
}

// DivineSpells godoc
// @Tags Site
// @Summary Сакральные заклинания
// @Description Сакральные заклинания
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /rules/spell/divineSpells [get]
func (h *Handler) DivineSpells(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "DivineSpells")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "divineSpells.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Сакральные заклинания",
				Description: "Сакральные заклинания",
			},
		})
	}
}

// SpellSpecialAbilities godoc
// @Tags Site
// @Summary Особые способности
// @Description Особые способности
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /rules/spell/specialAbilities [get]
func (h *Handler) SpellSpecialAbilities(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "SpellSpecialAbilities")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "spellSpecialAbilities.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Особые способности",
				Description: "Особые способности",
			},
		})
	}
}

// WizzardSpellSchools godoc
// @Tags Site
// @Summary Школы магии волшебника
// @Description Школы магии волшебника
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /spell/school [get]
func (h *Handler) WizzardSpellSchools(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "WizzardSpellSchools")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "spellSchool.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Школы магии волшебника",
				Description: "Школы магии волшебника",
			},
		})
	}
}

// WizzardSpellSchoolInfo godoc
// @Tags Site
// @Summary Информация по школе
// @Description Информация по школе
// @Produce html
// @Param alias path string true "Alias школы"
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /spell/school/{alias} [get]
func (h *Handler) WizzardSpellSchoolInfo(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "WizzardSpellSchoolInfo")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		alias := c.Param("alias")

		title, err := h.postgres.GetNameByAlias(c, "school", alias, true)
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
			Alias:   alias,
			JsPaths: []string{"spell/schoolInfo"},
		})
	}
}

// SpellList godoc
// @Tags Site
// @Summary Все заклинания
// @Description Все заклинания
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /spell/list [get]
func (h *Handler) SpellList(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "SpellList")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "spellList.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Все заклинания",
				Description: "Все заклинания",
			},
		})
	}
}

// ClassSpellList godoc
// @Tags Site
// @Summary Список заклинаний для класса
// @Description Список заклинаний для класса
// @Produce html
// @Param classAlias path string true "Alias класса"
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /spell/list/{classAlias} [get]
func (h *Handler) ClassSpellList(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "ClassSpellList")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		classAlias := c.Param("classAlias")

		redirectUrlQuery := url.Values{}
		redirectUrlQuery.Set("menu", "level")
		redirectUrlQuery.Set("menu2", classAlias)
		redirectUrl := url.URL{Path: "/spell/list", RawQuery: redirectUrlQuery.Encode()}
		c.Redirect(http.StatusPermanentRedirect, redirectUrl.RequestURI())
	}
}

// SpellInfo godoc
// @Tags Site
// @Summary Информация по Заклинанию
// @Description Информация по Заклинанию
// @Produce html
// @Param alias path string true "Alias заклинания"
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /spell/{alias} [get]
func (h *Handler) SpellInfo(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "SpellInfo")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		alias := c.Param("alias")

		title, err := h.postgres.GetNameByAlias(c, "spell", alias, true)
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
			Alias:   alias,
			JsPaths: []string{"spell/info"},
		})
	}
}
