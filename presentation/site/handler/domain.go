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

// Domains godoc
// @Tags Site
// @Summary Сферы/Инквизиции
// @Description Сферы/Инквизиции
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /god/domain [get]
func (h *Handler) Domains(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "Domains")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "domain.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Сферы/Инквизиции",
				Description: "Сферы/Инквизиции",
			},
		})
	}
}

// DomainInfo godoc
// @Tags Site
// @Summary Информация по сфере
// @Description Информация по сфере
// @Produce html
// @Param alias path string true "Alias сферы"
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /god/domain/{alias} [get]
func (h *Handler) DomainInfo(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		h.domainInfo(ctx, c, "domain", "domainInfo.html", []string{"domain/info"})
	}
}

// SubdomainInfo godoc
// @Tags Site
// @Summary Информация по подсфере
// @Description Информация по подсфере
// @Produce html
// @Param alias path string true "Alias подсферы"
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /god/subdomain/{alias} [get]
func (h *Handler) SubdomainInfo(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		h.domainInfo(ctx, c, "subdomain", "subdomainInfo.html", []string{"domain/subDomainInfo"})
	}
}

// InquisitionInfo godoc
// @Tags Site
// @Summary Информация по инквизиции
// @Description Информация по инквизиции
// @Produce html
// @Param alias path string true "Alias инквизиции"
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /god/inquisition/{alias} [get]
func (h *Handler) InquisitionInfo(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		h.domainInfo(ctx, c, "inquisition", "domainInfo.html", []string{"domain/info"})
	}
}

func (h *Handler) domainInfo(ctx context.Context, c *gin.Context, domainType string, templateName string, jsPaths []string) {
	/*	span := jaeger.GetSpan(ctx, "domainInfo")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	alias := c.Param("alias")

	title, err := h.postgres.GetDomainName(c, domainType, alias)
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
		Alias2:  &domainType,
		JsPaths: jsPaths,
	})
}
