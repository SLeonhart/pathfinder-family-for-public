package dbInterface

import (
	"context"
	"pathfinder-family/model"

	"github.com/jmoiron/sqlx"
)

//go:generate mockgen -source=iPostgres.go -destination=../dbMock/postgres.go -package=dbMock

type IPostgres interface {
	GetClient() *sqlx.DB
	Close()
	GetSearchInfo(ctx context.Context) ([]model.SearchInfo, error)
	GetNameByAlias(ctx context.Context, table string, alias string, withEng bool) (*string, error)
	GetNews(ctx context.Context, offset int, limit int, onlyActual bool) ([]model.News, error)
	GetThanks(ctx context.Context) ([]model.Thanks, error)
	GetTranslations(ctx context.Context, alias string) ([]model.Translation, error)
	GetBooks(ctx context.Context) ([]model.Books, error)
	GetBookInfo(ctx context.Context, bookAlias string) (*model.BookInfo, error)
	SendFeedback(ctx context.Context, theme string, email *string, message string) error
	GetRaces(ctx context.Context) ([]model.Race, error)
	GetRaceInfo(ctx context.Context, raceAlias string) (*model.RaceInfo, error)
	GetClasses(ctx context.Context) ([]model.Class, error)
	GetClassInfo(ctx context.Context, classAlias string) (*model.ClassInfo, error)
	GetArchetypes(ctx context.Context, classAlias *string) ([]model.ClassWithArchetypes, error)
	GetArchetypeInfo(ctx context.Context, archetypeAlias string) (*model.ArchetypeInfo, error)
	GetPrestigeClasses(ctx context.Context) ([]model.Class, error)
	GetPrestigeClassInfo(ctx context.Context, classAlias string) (*model.PrestigeClassInfo, error)
	GetNpcs(ctx context.Context) ([]model.Class, error)
	GetNpcInfo(ctx context.Context, classAlias string) (*model.NpcInfo, error)
	GetSkills(ctx context.Context) ([]model.SkillWithClasses, error)
	GetSkillsPerLvl(ctx context.Context) ([]model.SkillsPerLvlInfo, error)
	GetSkillInfo(ctx context.Context, skillAlias string) (*model.SkillInfo, error)
	GetFeats(ctx context.Context) ([]model.Feat, error)
	GetFeatInfo(ctx context.Context, featAlias string) (*model.FeatInfo, error)
	GetTraits(ctx context.Context) ([]model.Trait, error)
	GetGods(ctx context.Context) ([]model.God, error)
	GetDomains(ctx context.Context) ([]model.Domain, error)
	GetDomainInfo(ctx context.Context, domainType string, domainAlias string) (*model.DomainInfo, error)
	GetDomainName(ctx context.Context, domainType string, alias string) (*string, error)
	GetBloodlines(ctx context.Context, classAlias string) ([]model.Bloodline, error)
	GetBloodlineInfo(ctx context.Context, classAlias string, bloodlineAlias string) (*model.BloodlineInfo, error)
	GetBloodlineName(ctx context.Context, classAlias string, bloodlineAlias string) (*string, error)
	GetOrders(ctx context.Context) ([]model.Order, error)
	GetOrderInfo(ctx context.Context, bloodlineAlias string) (*model.OrderInfo, error)
	GetSpellSchools(ctx context.Context) ([]model.School, error)
	GetSpellSchoolInfo(ctx context.Context, schoolAlias string) (*model.SchoolInfo, error)
	GetSpells(ctx context.Context, classAlias *string) ([]model.Spell, error)
	GetSpellInfo(ctx context.Context, spellAlias string) (*model.SpellInfo, error)
	GetWeapons(ctx context.Context) ([]model.Weapon, error)
	GetArmors(ctx context.Context) ([]model.Armor, error)
	GetGoodsAndServices(ctx context.Context) ([]model.GoodAndService, error)
	GetMagicItemInfo(ctx context.Context, magicItemAlias string) (*model.MagicItemInfo, error)
	GetAllMagicItems(ctx context.Context) ([]model.MagicItemForList, error)
	GetMagicItemsByTypes(ctx context.Context, types []string) ([]model.MagicItemInfo, error)
	GetMagicItemAbilitiesByTypes(ctx context.Context, types []string) ([]model.MagicItemAbility, error)
	GetBeasts(ctx context.Context) ([]model.Beast, error)
	GetBeastInfo(ctx context.Context, beastAlias string) (*model.BeastInfo, error)
	GetMonsterAbilities(ctx context.Context) ([]model.MonsterAbility, error)
	GetCreatureTypes(ctx context.Context) ([]model.CreatureType, error)
	GetAnimalCompanions(ctx context.Context) ([]model.AnimalCompanion, error)
	GetPlantCompanions(ctx context.Context) ([]model.AnimalCompanion, error)
	GetMonstrousMounts(ctx context.Context) ([]model.MonstrousMount, error)
	GetAbilities(ctx context.Context) ([]model.NameAlias, error)
	GetAbilityInfo(ctx context.Context, abilityAlias string) (*model.AbilityInfo, error)
	GetAfflictions(ctx context.Context) ([]model.Affliction, error)
	GetTraps(ctx context.Context) ([]model.Trap, error)
	GetHaunts(ctx context.Context) ([]model.Haunt, error)
	GetGoods(ctx context.Context, token *string) ([]model.Good, error)
	GetUser(ctx context.Context, token string) (*model.UserData, error)
	AddGoodInWaitingList(ctx context.Context, userId int, goodId int) error
	UserAuth(ctx context.Context, request model.UserAuthRequest) (*model.UserData, error)
	UserRegister(ctx context.Context, request model.UserRegisterRequest) (*model.UserData, error)
	UserResetPassword(ctx context.Context, request model.UserResetPasswordRequest, password string) (*model.UserData, error)
	UserChangeData(ctx context.Context, userId int, request model.UserChangeDataRequest) (*model.UserData, error)
	GetUserFavourites(ctx context.Context, userId int) ([]model.Favourites, error)
	AddUserFavourites(ctx context.Context, userId int, guid string) error
	DeleteUserFavourites(ctx context.Context, userId int, guid string) error
	ChangeUserFavouritesItems(ctx context.Context, userId int, request model.ChangeUserFavouritesItemsRequest) error
	Logout(ctx context.Context, userId int) error
	RenameUserFavourites(ctx context.Context, userId int, request model.RenameUserFavouritesRequest) error
	AddDonate(ctx context.Context, request model.AddDonateRequest) error
	GetPushTokens(ctx context.Context) ([]model.PushToken, error)
	DeletePushToken(ctx context.Context, id int) error
	AddNews(ctx context.Context, request model.AddNewsRequest) error
	GetWildTalents(ctx context.Context) ([]model.WildTalent, error)
	GetAspects(ctx context.Context) ([]model.Aspect, error)
	GetCommonPage(ctx context.Context, id int) (*model.StaticPage, error)
	AddCommonPage(ctx context.Context, page model.StaticPage) (*int, error)
	UpdateCommonPage(ctx context.Context, page model.StaticPage) error
	GetShamanSpirits(ctx context.Context) ([]model.ShamanSpiritInfo, error)
	GetMediumSpirits(ctx context.Context, alias string) ([]model.MediumSpiritInfo, error)
	GetBotSpells(ctx context.Context, id *int, name *string, engName *string, classId *int, alias *string, level *int, rulebookIds []int) ([]model.BotSpellInfo, error)
	GetBotClasses(ctx context.Context, id *int, alias *string, magicClass *bool) ([]model.BotClassInfo, error)
	GetBotBooks(ctx context.Context, withSpells bool) ([]model.BotBook, error)
}
