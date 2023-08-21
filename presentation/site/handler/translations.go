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

// Translations godoc
// @Tags Site
// @Summary Что нового
// @Description Что нового
// @Produce html
// @Param translationType path string true "Тип переводов"
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /translations/{translationType} [get]
func (h *Handler) Translations(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "Translations")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		translationType := c.Param("translationType")

		alias := ""
		switch translationType {
		case "pathfinderAdventurePath":
			alias = "adventurePath"
		case "pathfinderModules":
			alias = "module"
		case "pathfinderPlayerCompanion":
			alias = "playerCompanion"
		case "pathfinderCampaignSetting":
			alias = "campaignSetting"
		case "pathfinderSocietyOrganizedPlay":
			alias = "societyOrganizedPlay"
		case "pathfinderRoleplayingGame":
			alias = "roleplayingGame"
		}

		title, err := h.postgres.GetNameByAlias(c, "translation_type", alias, false)
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
			JsPaths: []string{"translations/main"},
		})
	}
}
