package server

import (
	"context"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"pathfinder-family/config"
	"pathfinder-family/data/cache/cacheInterface"
	"pathfinder-family/data/db/dbInterface"

	a "pathfinder-family/presentation/api/handler"
	"pathfinder-family/presentation/api/middleware"
	s "pathfinder-family/presentation/site/handler"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type WebServer struct {
	cgf         *config.Config
	engine      *gin.Engine
	apiHandler  *a.Handler
	siteHandler *s.Handler
	inmemory    cacheInterface.IInMemory
	postgres    dbInterface.IPostgres
}

type WebServerOptions func(*WebServer)

func NewWebServer(config *config.Config, inmemory cacheInterface.IInMemory, postgres dbInterface.IPostgres, opts ...WebServerOptions) *WebServer {
	server := &WebServer{
		cgf:      config,
		engine:   gin.New(),
		inmemory: inmemory,
		postgres: postgres,
	}

	for _, opt := range opts {
		if opt == nil {
			continue
		}

		opt(server)
	}

	if server.apiHandler == nil {
		panic("Api handler should be set")
	}

	return server
}

func (s *WebServer) Run(ctx context.Context) error {
	gin.SetMode(s.cgf.App.Mode)

	s.engine.
		Use(gin.Recovery()).
		Use(cors.New(cors.Config{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"*"},
			AllowHeaders:     []string{"*"},
			ExposeHeaders:    []string{"*"},
			AllowCredentials: true,
			// AllowOriginFunc: func(origin string) bool {
			// 	return origin == "http://localhost:8888"
			// },
			MaxAge: 12 * time.Hour,
		}))

	s.setupRoutes(ctx)

	return s.engine.Run(s.cgf.App.ServerAddress())
}

func (s *WebServer) setupRoutes(ctx context.Context) {
	if s.apiHandler != nil {
		s.setupApiRoutes(ctx)
	}

	if s.siteHandler != nil {
		s.setupSiteRoutes(ctx)
	}

	s.setupSwagger()
}

func (s *WebServer) setupSwagger() {
	s.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func (s *WebServer) setupApiRoutes(ctx context.Context) {
	apiGroup := s.engine.Group("/api")

	authGroup := apiGroup.Group("")
	authGroup.Use(middleware.Auth(ctx, s.inmemory, s.postgres))
	{
		authGroup.GET("/goods", s.apiHandler.GetGoods(ctx))
		authGroup.POST("/goods/addWaitingList", s.apiHandler.AddGoodInWaitingList(ctx))
		userAuthGroup := authGroup.Group("/user")
		userAuthGroup.POST("/changeData", s.apiHandler.UserChangeData(ctx))
		userAuthGroup.GET("/favourites", s.apiHandler.UserFavourites(ctx))
		userAuthGroup.POST("/favourites", s.apiHandler.AddUserFavourites(ctx))
		userAuthGroup.PATCH("/favourites", s.apiHandler.RenameUserFavourites(ctx))
		userAuthGroup.DELETE("/favourites", s.apiHandler.DeleteUserFavourites(ctx))
		userAuthGroup.PATCH("/favourites/items", s.apiHandler.ChangeUserFavouritesItems(ctx))

		adminGroup := authGroup.Group("/admin")
		adminGroup.Use(middleware.RequireAdmin())
		{
			adminGroup.POST("/donate", s.apiHandler.AddDonate(ctx))
			adminGroup.POST("/push", s.apiHandler.SendPush(ctx))
			adminGroup.DELETE("/push", s.apiHandler.RemovePushTokens(ctx))
			adminGroup.POST("/telegram", s.apiHandler.SendTelegram(ctx))
			adminGroup.POST("/vk", s.apiHandler.SendVk(ctx))
			adminGroup.POST("/news", s.apiHandler.AddNews(ctx))
			searchGroup := adminGroup.Group("/search")
			searchGroup.PATCH("", s.apiHandler.UpsertSearch(ctx))
			searchGroup.POST("", s.apiHandler.UpdateSearch(ctx))
			searchGroup.DELETE("", s.apiHandler.ClearOldSearch(ctx))
		}

		guiGroup := authGroup.Group("/gui")
		guiGroup.Use(middleware.RequireContent())
		{
			guiGroup.GET("/commonPage", s.apiHandler.GetCommonPage(ctx))
			guiGroup.POST("/commonPage", s.apiHandler.UpsertCommonPage(ctx))
		}
	}
	userGroup := apiGroup.Group("/user")
	userGroup.POST("/auth", s.apiHandler.UserAuth(ctx))
	userGroup.POST("/register", s.apiHandler.UserRegister(ctx))
	userGroup.POST("/resetPassword", s.apiHandler.UserResetPassword(ctx))

	botGroup := apiGroup.Group("/bot")
	botGroup.GET("/spells", s.apiHandler.GetBotSpells(ctx))
	botGroup.GET("/classes", s.apiHandler.GetBotClasses(ctx))
	botGroup.GET("/rulebooks", s.apiHandler.GetBotRulebooks(ctx))

	newsGroup := apiGroup.Group("/news")
	newsGroup.GET("", s.apiHandler.GetNews(ctx))
	newsGroup.GET("/last", s.apiHandler.GetNewsLast(ctx))
	apiGroup.GET("/thanks", s.apiHandler.GetThanks(ctx))
	apiGroup.GET("/translations", s.apiHandler.GetTranslations(ctx))
	apiGroup.GET("/books", s.apiHandler.GetBooks(ctx))
	apiGroup.GET("/bookInfo", s.apiHandler.GetBookInfo(ctx))
	apiGroup.POST("/sendFeedback", s.apiHandler.SendFeedback(ctx))
	apiGroup.GET("/races", s.apiHandler.GetRaces(ctx))
	apiGroup.GET("/raceInfo", s.apiHandler.GetRaceInfo(ctx))
	apiGroup.GET("/classes", s.apiHandler.GetClasses(ctx))
	apiGroup.GET("/classInfo", s.apiHandler.GetClassInfo(ctx))
	apiGroup.GET("/archetypes", s.apiHandler.GetArchetypes(ctx))
	apiGroup.GET("/archetypeInfo", s.apiHandler.GetArchetypeInfo(ctx))
	apiGroup.GET("/prestigeClasses", s.apiHandler.GetPrestigeClasses(ctx))
	apiGroup.GET("/prestigeClassInfo", s.apiHandler.GetPrestigeClassInfo(ctx))
	apiGroup.GET("/npcs", s.apiHandler.GetNpcs(ctx))
	apiGroup.GET("/npcInfo", s.apiHandler.GetNpcInfo(ctx))
	apiGroup.GET("/skills", s.apiHandler.GetSkills(ctx))
	apiGroup.GET("/skillInfo", s.apiHandler.GetSkillInfo(ctx))
	apiGroup.GET("/feats", s.apiHandler.GetFeats(ctx))
	apiGroup.GET("/featInfo", s.apiHandler.GetFeatInfo(ctx))
	apiGroup.GET("/traits", s.apiHandler.GetTraits(ctx))
	apiGroup.GET("/gods", s.apiHandler.GetGods(ctx))
	apiGroup.GET("/domains", s.apiHandler.GetDomains(ctx))
	apiGroup.GET("/domainInfo", s.apiHandler.GetDomainInfo(ctx))
	apiGroup.GET("/bloodlines", s.apiHandler.GetBloodlines(ctx))
	apiGroup.GET("/bloodlineInfo", s.apiHandler.GetBloodlineInfo(ctx))
	apiGroup.GET("/orders", s.apiHandler.GetOrders(ctx))
	apiGroup.GET("/orderInfo", s.apiHandler.GetOrderInfo(ctx))
	apiGroup.GET("/spellSchools", s.apiHandler.GetSpellSchools(ctx))
	apiGroup.GET("/spellSchoolInfo", s.apiHandler.GetSpellSchoolInfo(ctx))
	apiGroup.GET("/spells", s.apiHandler.GetSpells(ctx))
	apiGroup.GET("/spellInfo", s.apiHandler.GetSpellInfo(ctx))
	apiGroup.GET("/weapons", s.apiHandler.GetWeapons(ctx))
	apiGroup.GET("/armors", s.apiHandler.GetArmors(ctx))
	apiGroup.GET("/goodsAndServices", s.apiHandler.GetGoodsAndServices(ctx))
	apiGroup.GET("/magicItemInfo", s.apiHandler.GetMagicItemInfo(ctx))
	apiGroup.GET("/allMagicItems", s.apiHandler.GetAllMagicItems(ctx))
	apiGroup.GET("/magicItems", s.apiHandler.GetMagicItemsByTypes(ctx))
	apiGroup.GET("/magicItemAbilities", s.apiHandler.GetMagicItemAbilitiesByTypes(ctx))
	apiGroup.GET("/beasts", s.apiHandler.GetBeasts(ctx))
	apiGroup.GET("/beastInfo", s.apiHandler.GetBeastInfo(ctx))
	apiGroup.GET("/monsterAbilities", s.apiHandler.GetMonsterAbilities(ctx))
	apiGroup.GET("/creatureTypes", s.apiHandler.GetCreatureTypes(ctx))
	apiGroup.GET("/animalCompanions", s.apiHandler.GetAnimalCompanions(ctx))
	apiGroup.GET("/abilities", s.apiHandler.GetAbilities(ctx))
	apiGroup.GET("/abilityInfo", s.apiHandler.GetAbilityInfo(ctx))
	apiGroup.GET("/afflictions", s.apiHandler.GetAfflictions(ctx))
	apiGroup.GET("/traps", s.apiHandler.GetTraps(ctx))
	apiGroup.GET("/haunts", s.apiHandler.GetHaunts(ctx))
	apiGroup.GET("/search", s.apiHandler.GetSearchResult(ctx))
	apiGroup.GET("/wildTalents", s.apiHandler.GetWildTalents(ctx))
	apiGroup.GET("/aspects", s.apiHandler.GetAspects(ctx))
	spiritsGroup := apiGroup.Group("/spirits")
	spiritsGroup.GET("/shaman", s.apiHandler.GetShamanSpirits(ctx))
	spiritsGroup.GET("/:alias", s.apiHandler.GetMediumSpirits(ctx))
}

func (s *WebServer) setupSiteRoutes(ctx context.Context) {
	s.engine.LoadHTMLGlob("./htmlTemplates/*.html")
	s.engine.Static("/files", "./files")

	userGroup := s.engine.Group("/user")
	userGroup.Use(middleware.Auth(ctx, s.inmemory, s.postgres))
	{
		userGroup.GET("/login", s.siteHandler.Login(ctx))
		userGroup.GET("/logout", s.siteHandler.Logout(ctx))
		userGroup.GET("/profile", s.siteHandler.Profile(ctx))
		userGroup.GET("/favourites", s.siteHandler.Favourites(ctx))
	}

	adminGroup := s.engine.Group("/admin")
	adminGroup.Use(middleware.Auth(ctx, s.inmemory, s.postgres))
	{
		adminGroup.GET("/profile", s.siteHandler.AdminProfile(ctx))
	}

	s.engine.GET("/shop", s.siteHandler.Shop(ctx))

	s.engine.GET("/openGameLicense", s.siteHandler.OpenGameLicense(ctx))
	s.engine.GET("", s.siteHandler.Index(ctx))
	s.engine.GET("/aboutMe", s.siteHandler.AboutMe(ctx))
	s.engine.GET("/thanks", s.siteHandler.Thanks(ctx))
	s.engine.GET("/dice", s.siteHandler.Dice(ctx))
	s.engine.GET("/translations/:translationType", s.siteHandler.Translations(ctx))
	booksGroup := s.engine.Group("/books")
	booksGroup.GET("", s.siteHandler.Books(ctx))
	booksGroup.GET("/:alias", s.siteHandler.Book(ctx))
	s.engine.GET("/contacts", s.siteHandler.Contacts(ctx))
	raceGroup := s.engine.Group("/race")
	raceGroup.GET("", s.siteHandler.Races(ctx))
	raceGroup.GET("/:alias", s.siteHandler.Race(ctx))
	classGroup := s.engine.Group("/class")
	classGroup.GET("", s.siteHandler.Classes(ctx))
	classGroup.GET("/:alias", s.siteHandler.ClassInfo(ctx))
	classGroup.GET("/:alias/wildTalent", s.siteHandler.WildTalents(ctx))
	classGroup.GET("/:alias/spirit", s.siteHandler.Spirits(ctx))
	classGroup.GET("/:alias/aspect", s.siteHandler.Aspect(ctx))
	orderGroup := classGroup.Group("/:alias/order")
	orderGroup.GET("", s.siteHandler.Orders(ctx))
	orderGroup.GET("/:orderAlias", s.siteHandler.OrderInfo(ctx))
	bloodlineGroup := classGroup.Group("/:alias/bloodline")
	bloodlineGroup.GET("", s.siteHandler.Bloodlines(ctx))
	bloodlineGroup.GET("/:bloodlineAlias", s.siteHandler.BloodlineInfo(ctx))
	archetypeGroup := classGroup.Group("/archetype")
	archetypeGroup.GET("", s.siteHandler.AllArchetypes(ctx))
	archetypeGroup.GET("/:classAlias", s.siteHandler.ClassArchetypes(ctx))
	archetypeGroup.GET("/:classAlias/:alias", s.siteHandler.ArchetypeInfo(ctx))
	prestigeClassGroup := classGroup.Group("/prestige")
	prestigeClassGroup.GET("", s.siteHandler.PrestigeClasses(ctx))
	prestigeClassGroup.GET("/:alias", s.siteHandler.PrestigeClassInfo(ctx))
	npcGroup := s.engine.Group("/npc")
	npcGroup.GET("", s.siteHandler.Npcs(ctx))
	npcGroup.GET("/:alias", s.siteHandler.NpcInfo(ctx))
	skillGroup := s.engine.Group("/skill")
	skillGroup.GET("", s.siteHandler.Skills(ctx))
	skillGroup.GET("/:alias", s.siteHandler.SkillInfo(ctx))
	featGroup := s.engine.Group("/feat")
	featGroup.GET("", s.siteHandler.Feats(ctx))
	featGroup.GET("/:alias", s.siteHandler.FeatInfo(ctx))
	traitsGroup := s.engine.Group("/traits")
	traitsGroup.GET("", s.siteHandler.Traits(ctx))
	godGroup := s.engine.Group("/god")
	godGroup.GET("", s.siteHandler.Gods(ctx))
	godGroup.GET("/domain", s.siteHandler.Domains(ctx))
	godGroup.GET("/domain/:alias", s.siteHandler.DomainInfo(ctx))
	godGroup.GET("/subdomain/:alias", s.siteHandler.SubdomainInfo(ctx))
	godGroup.GET("/inquisition/:alias", s.siteHandler.InquisitionInfo(ctx))
	spellGroup := s.engine.Group("/spell")
	spellListGroup := spellGroup.Group("/list")
	spellListGroup.GET("", s.siteHandler.SpellList(ctx))
	spellListGroup.GET("/:classAlias", s.siteHandler.ClassSpellList(ctx))
	spellGroup.GET("/:alias", s.siteHandler.SpellInfo(ctx))
	schoolGroup := spellGroup.Group("/school")
	schoolGroup.GET("", s.siteHandler.WizzardSpellSchools(ctx))
	schoolGroup.GET("/:alias", s.siteHandler.WizzardSpellSchoolInfo(ctx))
	s.engine.GET("/wealthAndMoney", s.siteHandler.WealthAndMoney(ctx))
	weaponsGroup := s.engine.Group("/weapons")
	weaponsGroup.GET("", s.siteHandler.Weapons(ctx))
	weaponsGroup.GET("/description", s.siteHandler.WeaponsMainDescription(ctx))
	armorsGroup := s.engine.Group("/armors")
	armorsGroup.GET("", s.siteHandler.Armors(ctx))
	armorsGroup.GET("/description", s.siteHandler.ArmorsMainDescription(ctx))
	s.engine.GET("/specialMaterials", s.siteHandler.SpecialMaterials(ctx))
	s.engine.GET("/goodsAndServices", s.siteHandler.GoodsAndServices(ctx))
	s.engine.GET("/magicItem/:alias", s.siteHandler.MagicItemInfo(ctx))
	magicItemsGroup := s.engine.Group("/magicItems")
	magicItemsGroup.GET("", s.siteHandler.MagicItems(ctx))
	magicItemsGroup.GET("/usingItems", s.siteHandler.UsingItems(ctx))
	magicItemsGroup.GET("/magicItemsOnTheBody", s.siteHandler.MagicItemsOnTheBody(ctx))
	magicItemsGroup.GET("/savingThrowsAgainstMagicItemPowers", s.siteHandler.SavingThrowsAgainstMagicItemPowers(ctx))
	magicItemsGroup.GET("/damagingMagicItems", s.siteHandler.DamagingMagicItems(ctx))
	magicItemsGroup.GET("/purchasingMagicItems", s.siteHandler.PurchasingMagicItems(ctx))
	magicItemsGroup.GET("/magicItemDescriptions", s.siteHandler.MagicItemDescriptions(ctx))
	magicItemsGroup.GET("/magicItemCreation", s.siteHandler.MagicItemCreation(ctx))
	magicItemsGroup.GET("/armor", s.siteHandler.MagicArmor(ctx))
	magicItemsGroup.GET("/weapons", s.siteHandler.MagicWeapons(ctx))
	magicItemsGroup.GET("/runeforgedWeapon", s.siteHandler.RuneforgedWeapon(ctx))
	magicItemsGroup.GET("/potions", s.siteHandler.Potions(ctx))
	magicItemsGroup.GET("/rings", s.siteHandler.Rings(ctx))
	magicItemsGroup.GET("/rods", s.siteHandler.Rods(ctx))
	magicItemsGroup.GET("/scrolls", s.siteHandler.Scrolls(ctx))
	magicItemsGroup.GET("/staves", s.siteHandler.Staves(ctx))
	magicItemsGroup.GET("/wands", s.siteHandler.Wands(ctx))
	magicItemsGroup.GET("/wondrousItems", s.siteHandler.WondrousItems(ctx))
	magicItemsGroup.GET("/tattooMagic", s.siteHandler.TattooMagic(ctx))
	magicItemsGroup.GET("/intelligentItems", s.siteHandler.IntelligentItems(ctx))
	magicItemsGroup.GET("/cursedItems", s.siteHandler.CursedItems(ctx))
	magicItemsGroup.GET("/specificCursedItems", s.siteHandler.SpecificCursedItems(ctx))
	magicItemsGroup.GET("/artifacts", s.siteHandler.Artifacts(ctx))
	bestiaryGroup := s.engine.Group("/bestiary")
	bestiaryGroup.GET("", s.siteHandler.Bestiary(ctx))
	bestiaryGroup.GET("/beast/:alias", s.siteHandler.BeastInfo(ctx))
	bestiaryGroup.GET("/description", s.siteHandler.BestiaryMainDescription(ctx))
	bestiaryGroup.GET("/animalCompanion", s.siteHandler.AnimalCompanion(ctx))
	bestiaryGroup.GET("/familiar", s.siteHandler.Familiar(ctx))
	bestiaryGroup.GET("/eidolon", s.siteHandler.Eidolon(ctx))
	bestiaryGroup.GET("/phantom", s.siteHandler.Phantom(ctx))
	bestiaryAppendixGroup := bestiaryGroup.Group("/appendix")
	bestiaryAppendixGroup.GET("/monsterCreation", s.siteHandler.MonsterCreation(ctx))
	bestiaryAppendixGroup.GET("/simpleTemplates", s.siteHandler.SimpleTemplates(ctx))
	bestiaryAppendixGroup.GET("/acquiredTemplates", s.siteHandler.AcquiredTemplates(ctx))
	bestiaryAppendixGroup.GET("/addingRacialHitDice", s.siteHandler.AddingRacialHitDice(ctx))
	bestiaryAppendixGroup.GET("/universalMonsterRules", s.siteHandler.UniversalMonsterRules(ctx))
	bestiaryAppendixGroup.GET("/creatureTypes", s.siteHandler.CreatureTypes(ctx))
	bestiaryAppendixGroup.GET("/monstersAsPCs", s.siteHandler.MonstersAsPCs(ctx))
	bestiaryAppendixGroup.GET("/monsterRoles", s.siteHandler.MonsterRoles(ctx))
	bestiaryAppendixGroup.GET("/encounterTables", s.siteHandler.EncounterTables(ctx))
	abilityGroup := s.engine.Group("/ability")
	abilityGroup.GET("", s.siteHandler.Abilities(ctx))
	abilityGroup.GET("/:alias", s.siteHandler.AbilityInfo(ctx))
	s.engine.GET("/about", s.siteHandler.AboutGame(ctx))
	s.engine.GET("/commonDictionary", s.siteHandler.CommonDictionary(ctx))
	s.engine.GET("/characterCreate", s.siteHandler.CharacterCreate(ctx))
	s.engine.GET("/alignment", s.siteHandler.Alignment(ctx))
	s.engine.GET("/vitalStatistics", s.siteHandler.VitalStatistics(ctx))
	s.engine.GET("/movement", s.siteHandler.Movement(ctx))
	s.engine.GET("/exploration", s.siteHandler.Exploration(ctx))
	s.engine.GET("/specialAbilities", s.siteHandler.SpecialAbilities(ctx))
	s.engine.GET("/conditions", s.siteHandler.Conditions(ctx))
	s.engine.GET("/afflictions", s.siteHandler.Afflictions(ctx))
	s.engine.GET("/gameExample", s.siteHandler.GameExample(ctx))
	s.engine.GET("/howCombatWorks", s.siteHandler.HowCombatWorks(ctx))
	s.engine.GET("/combatStatistics", s.siteHandler.CombatStatistics(ctx))
	s.engine.GET("/actionsInCombat", s.siteHandler.ActionsInCombat(ctx))
	s.engine.GET("/injuryAndDeath", s.siteHandler.InjuryAndDeath(ctx))
	s.engine.GET("/movementAndDistance", s.siteHandler.MovementAndDistance(ctx))
	s.engine.GET("/bigAndLittleCreatures", s.siteHandler.BigAndLittleCreatures(ctx))
	s.engine.GET("/combatModifiers", s.siteHandler.CombatModifiers(ctx))
	s.engine.GET("/specialAttacks", s.siteHandler.SpecialAttacks(ctx))
	s.engine.GET("/specialInitiativeActions", s.siteHandler.SpecialInitiativeActions(ctx))
	s.engine.GET("/startingCampaign", s.siteHandler.StartingCampaign(ctx))
	s.engine.GET("/buildingAdventure", s.siteHandler.BuildingAdventure(ctx))
	s.engine.GET("/preparingGame", s.siteHandler.PreparingGame(ctx))
	s.engine.GET("/duringGame", s.siteHandler.DuringGame(ctx))
	s.engine.GET("/campaignTips", s.siteHandler.CampaignTips(ctx))
	s.engine.GET("/endingCampaign", s.siteHandler.EndingCampaign(ctx))
	variantRulesGroup := s.engine.Group("/variantRules")
	variantRulesGroup.GET("/calledShots", s.siteHandler.CalledShots(ctx))
	variantRulesGroup.GET("/armorAsDamageReduction", s.siteHandler.ArmorAsDamageReduction(ctx))
	rulesGroup := s.engine.Group("/rules")
	rulesGroup.GET("/characterAdvancement", s.siteHandler.CharacterAdvancement(ctx))
	rulesGroup.GET("/class", s.siteHandler.ClassRules(ctx))
	rulesGroup.GET("/archetype", s.siteHandler.ArchetypeRules(ctx))
	rulesGroup.GET("/npc", s.siteHandler.NpcRules(ctx))
	rulesGroup.GET("/feat", s.siteHandler.FeatRules(ctx))
	rulesGroup.GET("/skill", s.siteHandler.SkillsRules(ctx))
	rulesGroup.GET("/trait", s.siteHandler.TraitsRules(ctx))
	rulesGroup.GET("/duels", s.siteHandler.Duels(ctx))
	spellRulesGroup := rulesGroup.Group("/spell")
	spellRulesGroup.GET("/description", s.siteHandler.SpellMainDescription(ctx))
	spellRulesGroup.GET("/castingSpells", s.siteHandler.CastingSpells(ctx))
	spellRulesGroup.GET("/spellDescriptions", s.siteHandler.SpellDescriptions(ctx))
	spellRulesGroup.GET("/schools", s.siteHandler.SpellSchools(ctx))
	spellRulesGroup.GET("/arcaneSpells", s.siteHandler.ArcaneSpells(ctx))
	spellRulesGroup.GET("/divineSpells", s.siteHandler.DivineSpells(ctx))
	spellRulesGroup.GET("/specialAbilities", s.siteHandler.SpellSpecialAbilities(ctx))
	s.engine.GET("/dungeons", s.siteHandler.Dungeons(ctx))
	trapsGroup := s.engine.Group("/traps")
	trapsGroup.GET("", s.siteHandler.Traps(ctx))
	trapsGroup.GET("/description", s.siteHandler.TrapsMainDescription(ctx))
	hauntsGroup := s.engine.Group("/haunts")
	hauntsGroup.GET("", s.siteHandler.Haunts(ctx))
	hauntsGroup.GET("/description", s.siteHandler.HauntsMainDescription(ctx))
	s.engine.GET("/wilderness", s.siteHandler.Wilderness(ctx))
	s.engine.GET("/urbanAdventures", s.siteHandler.UrbanAdventures(ctx))
	s.engine.GET("/weather", s.siteHandler.Weather(ctx))
	s.engine.GET("/thePlanes", s.siteHandler.ThePlanes(ctx))
	s.engine.GET("/environmentalRules", s.siteHandler.EnvironmentalRules(ctx))
	environmentGroup := s.engine.Group("/environment")
	environmentGroup.GET("/hazards", s.siteHandler.Hazards(ctx))
	specialAreaGroup := environmentGroup.Group("/specialArea")
	specialAreaGroup.GET("/irrisen", s.siteHandler.Irrisen(ctx))
	generatorGroup := s.engine.Group("/generator")
	generatorGroup.GET("/adventure", s.siteHandler.AdventureGenerator(ctx))
	generatorGroup.GET("/smallProblem", s.siteHandler.SmallProblemGenerator(ctx))
	s.engine.GET("/search/:query", s.siteHandler.Search(ctx))
}

func WithApiHandler(handler *a.Handler) WebServerOptions {
	return func(server *WebServer) {
		server.apiHandler = handler
	}
}

func WithSiteHandler(handler *s.Handler) WebServerOptions {
	return func(server *WebServer) {
		server.siteHandler = handler
	}
}
