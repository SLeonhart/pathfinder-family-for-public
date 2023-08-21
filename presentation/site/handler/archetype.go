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

// AllArchetypes godoc
// @Tags Site
// @Summary Все Архетипы
// @Description Все Архетипы
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /class/archetype [get]
func (h *Handler) AllArchetypes(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "AllArchetypes")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		c.HTML(http.StatusOK, "defaultPage.html", model.DeafultPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Архетипы",
				Description: "Архетипы",
			},
			JsPaths: []string{"archetype/all"},
		})
	}
}

// ClassArchetypes godoc
// @Tags Site
// @Summary Архетипы Класса
// @Description Архетипы Класса
// @Produce html
// @Param classAlias path string true "Alias класса"
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /class/archetype/{classAlias} [get]
func (h *Handler) ClassArchetypes(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "ClassArchetypes")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		classAlias := c.Param("classAlias")

		title, err := h.postgres.GetNameByAlias(c, "class", classAlias, false)
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
			Alias:   classAlias,
			JsPaths: []string{"archetype/class"},
		})
	}
}

// ArchetypeInfo godoc
// @Tags Site
// @Summary Информация по классу
// @Description Информация по классу
// @Produce html
// @Param alias path string true "Alias народа"
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /class/archetype/{classAlias}/{alias} [get]
func (h *Handler) ArchetypeInfo(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "ArchetypeInfo")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		alias := c.Param("alias")

		title, err := h.postgres.GetNameByAlias(c, "class", alias, true)
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
			JsPaths: []string{"archetype/info"},
		})
	}
}
