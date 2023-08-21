package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"pathfinder-family/model"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/pgtype"
)

// GetGoodsAndServices godoc
// @Tags API
// @Summary Список товаров и услуг
// @Description Список товаров и услуг
// @Produce json
// @Success 200 {object} []model.GoodAndService "OK"
// @Success 500 {object} model.ErrorResponse "INTERNAL_ERROR"
// @Router /api/goodsAndServices [get]
func (h *Handler) GetGoodsAndServices(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "GetGoodsAndServices")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		goodsAndServices, err := h.postgres.GetGoodsAndServices(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}
		for i := range goodsAndServices {
			if goodsAndServices[i].ParentsArray.Status == pgtype.Present {
				goodsAndServices[i].Parents = make([]string, 0, len(goodsAndServices[i].ParentsArray.Elements))
				for _, element := range goodsAndServices[i].ParentsArray.Elements {
					goodsAndServices[i].Parents = append(goodsAndServices[i].Parents, element.String)
				}
			}
			if goodsAndServices[i].TypeJson.Status == pgtype.Present {
				json.Unmarshal(goodsAndServices[i].TypeJson.Bytes, &goodsAndServices[i].Type)
			}
			if goodsAndServices[i].BookJson.Status == pgtype.Present {
				json.Unmarshal(goodsAndServices[i].BookJson.Bytes, &goodsAndServices[i].Book)
			}
			if goodsAndServices[i].ChildsJson.Status == pgtype.Present {
				json.Unmarshal(goodsAndServices[i].ChildsJson.Bytes, &goodsAndServices[i].Childs)
			}
		}
		c.JSON(http.StatusOK, goodsAndServices)
	}
}
