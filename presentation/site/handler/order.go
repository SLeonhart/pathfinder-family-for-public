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

// Orders godoc
// @Tags Site
// @Summary Наследия
// @Description Наследия
// @Produce html
// @Param alias path string true "Alias класса"
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /class/{alias}/order [get]
func (h *Handler) Orders(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "Orders")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		alias := c.Param("alias")
		if alias != "samurai_cavalier" {
			c.Status(http.StatusNotFound)
			return
		}

		c.HTML(http.StatusOK, "order.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Ордена самураев и кавалеристов",
				Description: "Ордена самураев и кавалеристов",
			},
		})
	}
}

// OrderInfo godoc
// @Tags Site
// @Summary Информация по ордену
// @Description Информация по ордену
// @Produce html
// @Param alias path string true "Alias класса"
// @Param orderAlias path string true "Alias ордена"
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /class/{alias}/order/{orderAlias} [get]
func (h *Handler) OrderInfo(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "OrderInfo")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		alias := c.Param("alias")
		orderAlias := c.Param("orderAlias")

		if alias != "samurai_cavalier" {
			c.Status(http.StatusNotFound)
			return
		}

		title, err := h.postgres.GetNameByAlias(c, "orders", orderAlias, true)
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
			Alias:   orderAlias,
			JsPaths: []string{"order/info"},
		})
	}
}
