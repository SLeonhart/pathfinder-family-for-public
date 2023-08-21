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

// UsingItems godoc
// @Tags Site
// @Summary Применение волшебных предметов
// @Description Применение волшебных предметов
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /magicItems/usingItems [get]
func (h *Handler) UsingItems(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "UsingItems")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "usingItems.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Применение волшебных предметов",
				Description: "Применение волшебных предметов",
			},
		})
	}
}

// MagicItemsOnTheBody godoc
// @Tags Site
// @Summary Ношение волшебных предметов
// @Description Ношение волшебных предметов
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /magicItems/magicItemsOnTheBody [get]
func (h *Handler) MagicItemsOnTheBody(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "MagicItemsOnTheBody")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "magicItemsOnTheBody.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Ношение волшебных предметов",
				Description: "Ношение волшебных предметов",
			},
		})
	}
}

// SavingThrowsAgainstMagicItemPowers godoc
// @Tags Site
// @Summary Испытания против свойств волшебных предметов
// @Description Испытания против свойств волшебных предметов
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /magicItems/savingThrowsAgainstMagicItemPowers [get]
func (h *Handler) SavingThrowsAgainstMagicItemPowers(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "SavingThrowsAgainstMagicItemPowers")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "savingThrowsAgainstMagicItemPowers.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Испытания против свойств волшебных предметов",
				Description: "Испытания против свойств волшебных предметов",
			},
		})
	}
}

// DamagingMagicItems godoc
// @Tags Site
// @Summary Повреждение волшебных предметов
// @Description Повреждение волшебных предметов
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /magicItems/damagingMagicItems [get]
func (h *Handler) DamagingMagicItems(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "DamagingMagicItems")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "damagingMagicItems.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Повреждение волшебных предметов",
				Description: "Повреждение волшебных предметов",
			},
		})
	}
}

// PurchasingMagicItems godoc
// @Tags Site
// @Summary Покупка волшебных предметов
// @Description Покупка волшебных предметов
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /magicItems/purchasingMagicItems [get]
func (h *Handler) PurchasingMagicItems(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "PurchasingMagicItems")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "purchasingMagicItems.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Покупка волшебных предметов",
				Description: "Покупка волшебных предметов",
			},
		})
	}
}

// MagicItemDescriptions godoc
// @Tags Site
// @Summary Описание волшебных предметов
// @Description Описание волшебных предметов
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /magicItems/magicItemDescriptions [get]
func (h *Handler) MagicItemDescriptions(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "MagicItemDescriptions")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "magicItemDescriptions.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Описание волшебных предметов",
				Description: "Описание волшебных предметов",
			},
		})
	}
}

// MagicItemCreation godoc
// @Tags Site
// @Summary Создание волшебных предметов
// @Description Создание волшебных предметов
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /magicItems/magicItemCreation [get]
func (h *Handler) MagicItemCreation(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "MagicItemCreation")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "magicItemCreation.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Создание волшебных предметов",
				Description: "Создание волшебных предметов",
			},
		})
	}
}

// MagicItemInfo godoc
// @Tags Site
// @Summary Информация по волшебному предмету
// @Description Информация по волшебному предмету
// @Produce html
// @Param alias path string true "Alias волшебного предмета"
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /magicItem/{alias} [get]
func (h *Handler) MagicItemInfo(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "MagicItemInfo")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/

		alias := c.Param("alias")

		title, err := h.postgres.GetNameByAlias(c, "magic_item", alias, true)
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
				"magicItem/main",
				"magicItem/info",
			},
		})
	}
}

// MagicItems godoc
// @Tags Site
// @Summary Все волшебные предметы
// @Description Все волшебные предметы
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /magicItems [get]
func (h *Handler) MagicItems(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "MagicItems")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "magicItems.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Все волшебные предметы",
				Description: "Все волшебные предметы",
			},
		})
	}
}

// MagicArmor godoc
// @Tags Site
// @Summary Броня
// @Description Броня
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /magicItems/armor [get]
func (h *Handler) MagicArmor(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "MagicArmor")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "magicArmor.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Броня",
				Description: "Броня",
			},
		})
	}
}

// MagicWeapons godoc
// @Tags Site
// @Summary Оружие
// @Description Оружие
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /magicItems/weapons [get]
func (h *Handler) MagicWeapons(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "MagicWeapons")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "magicWeapons.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Оружие",
				Description: "Оружие",
			},
		})
	}
}

// RuneforgedWeapon godoc
// @Tags Site
// @Summary Рунное оружие
// @Description Рунное оружие
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /magicItems/runeforgedWeapon [get]
func (h *Handler) RuneforgedWeapon(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "RuneforgedWeapon")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "runeforgedWeapon.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Рунное оружие",
				Description: "Рунное оружие",
			},
		})
	}
}

// Potions godoc
// @Tags Site
// @Summary Зелья
// @Description Зелья
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /magicItems/potions [get]
func (h *Handler) Potions(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "Potions")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "potions.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Зелья",
				Description: "Зелья",
			},
		})
	}
}

// Rings godoc
// @Tags Site
// @Summary Кольца
// @Description Кольца
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /magicItems/rings [get]
func (h *Handler) Rings(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "Rings")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "rings.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Кольца",
				Description: "Кольца",
			},
		})
	}
}

// Rods godoc
// @Tags Site
// @Summary Скипетры
// @Description Скипетры
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /magicItems/rods [get]
func (h *Handler) Rods(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "Rods")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "rods.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Скипетры",
				Description: "Скипетры",
			},
		})
	}
}

// Scrolls godoc
// @Tags Site
// @Summary Свитки
// @Description Свитки
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /magicItems/scrolls [get]
func (h *Handler) Scrolls(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "Scrolls")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "scrolls.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Свитки",
				Description: "Свитки",
			},
		})
	}
}

// Staves godoc
// @Tags Site
// @Summary Посохи
// @Description Посохи
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /magicItems/staves [get]
func (h *Handler) Staves(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "Staves")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "staves.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Посохи",
				Description: "Посохи",
			},
		})
	}
}

// Wands godoc
// @Tags Site
// @Summary Жезлы
// @Description Жезлы
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /magicItems/wands [get]
func (h *Handler) Wands(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "Wands")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "wands.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Жезлы",
				Description: "Жезлы",
			},
		})
	}
}

// WondrousItems godoc
// @Tags Site
// @Summary Волшебные вещицы
// @Description Волшебные вещицы
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /magicItems/wondrousItems [get]
func (h *Handler) WondrousItems(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "WondrousItems")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "wondrousItems.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Волшебные вещицы",
				Description: "Волшебные вещицы",
			},
		})
	}
}

// TattooMagic godoc
// @Tags Site
// @Summary Магические татуировки
// @Description Магические татуировки
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /magicItems/tattooMagic [get]
func (h *Handler) TattooMagic(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "TattooMagic")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "tattooMagic.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Магические татуировки",
				Description: "Магические татуировки",
			},
		})
	}
}

// IntelligentItems godoc
// @Tags Site
// @Summary Разумные предметы
// @Description Разумные предметы
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /magicItems/intelligentItems [get]
func (h *Handler) IntelligentItems(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "IntelligentItems")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "intelligentItems.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Разумные предметы",
				Description: "Разумные предметы",
			},
		})
	}
}

// CursedItems godoc
// @Tags Site
// @Summary Проклятые предметы
// @Description Проклятые предметы
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /magicItems/cursedItems [get]
func (h *Handler) CursedItems(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "CursedItems")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "cursedItems.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Проклятые предметы",
				Description: "Проклятые предметы",
			},
		})
	}
}

// SpecificCursedItems godoc
// @Tags Site
// @Summary Особые проклятые предметы
// @Description Особые проклятые предметы
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /magicItems/specificCursedItems [get]
func (h *Handler) SpecificCursedItems(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "SpecificCursedItems")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "specificCursedItems.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Особые проклятые предметы",
				Description: "Особые проклятые предметы",
			},
		})
	}
}

// Artifacts godoc
// @Tags Site
// @Summary Артефакты
// @Description Артефакты
// @Produce html
// @Success 200 {object} string "OK"
// @Success 500 {object} string "INTERNAL_ERROR"
// @Router /magicItems/artifacts [get]
func (h *Handler) Artifacts(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		/*		span := jaeger.GetSpan(ctx, "Artifacts")
				defer span.Finish()
				ctx = jaeger.SetParentSpan(ctx, span)*/
		c.HTML(http.StatusOK, "artifacts.html", model.CommonPage{
			Page: model.Page{
				Cfg:         h.cfg.App,
				Url:         utils.CurrentUrl(c, h.cfg.App.HostUrl),
				Title:       "Артефакты",
				Description: "Артефакты",
			},
		})
	}
}
