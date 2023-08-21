package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"pathfinder-family/model"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/pgtype"
)

// GetDomains godoc
// @Tags API
// @Summary Список Сфер/Инквизиций
// @Description Список Сфер/Инквизиций
// @Produce json
// @Success 200 {object} []model.Domain "OK"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/domains [get]
func (h *Handler) GetDomains(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "GetDomains")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		domains, err := h.postgres.GetDomains(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}
		for i := range domains {
			if domains[i].BookJson.Status == pgtype.Present {
				json.Unmarshal(domains[i].BookJson.Bytes, &domains[i].Book)
			}
			if domains[i].ChildsJson.Status == pgtype.Present {
				json.Unmarshal(domains[i].ChildsJson.Bytes, &domains[i].Childs)
			}
		}
		c.JSON(http.StatusOK, domains)
	}
}

// GetDomainInfo godoc
// @Tags API
// @Summary Информация по сфере/инквизиции
// @Description Информация по сфере/инквизиции
// @Produce json
// @Param alias query string true "Alias сферы/инквизиции"
// @Success 200 {object} model.DomainInfo "OK"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/domainInfo [get]
func (h *Handler) GetDomainInfo(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "GetDomainInfo")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		alias := c.Query("alias")
		domainType := c.Query("type")

		domainInfo, err := h.postgres.GetDomainInfo(c, domainType, alias)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}

		if domainInfo.GodsArray.Status == pgtype.Present {
			domainInfo.Gods = make([]string, 0, len(domainInfo.GodsArray.Elements))
			for _, element := range domainInfo.GodsArray.Elements {
				domainInfo.Gods = append(domainInfo.Gods, element.String)
			}
		}
		if domainInfo.ChildsJson.Status == pgtype.Present {
			json.Unmarshal(domainInfo.ChildsJson.Bytes, &domainInfo.Childs)
		}
		if domainInfo.BookJson.Status == pgtype.Present {
			json.Unmarshal(domainInfo.BookJson.Bytes, &domainInfo.Book)
		}
		if domainInfo.SpellsJson.Status == pgtype.Present {
			json.Unmarshal(domainInfo.SpellsJson.Bytes, &domainInfo.Spells)
		}
		if domainInfo.ParentsJson.Status == pgtype.Present {
			json.Unmarshal(domainInfo.ParentsJson.Bytes, &domainInfo.Parents)
		}

		c.JSON(http.StatusOK, domainInfo)
	}
}
