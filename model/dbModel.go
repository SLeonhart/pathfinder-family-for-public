package model

import (
	"time"

	"github.com/jackc/pgx/pgtype"
)

type News struct {
	Id             int       `db:"id" json:"id"`
	PushTitle      *string   `db:"push_title" json:"pushTitle"`
	PushBody       *string   `db:"push_body" json:"pushBody"`
	PushImageLink  *string   `db:"push_image_link" json:"pushImageLink"`
	PushTargetLink *string   `db:"push_target_link" json:"pushTargetLink"`
	Dt             time.Time `db:"dt" json:"dt"`
	Body           string    `db:"body" json:"body"`
}

type Thanks struct {
	Type       string       `db:"type" json:"type"`
	ListJson   pgtype.JSONB `db:"list" json:"-"`
	HelperStat []HelperStat `db:"-" json:"helperStat"`
}

type HelperStat struct {
	Name  string  `json:"name"`
	Alias *string `json:"alias"`
	Cnt   int     `json:"cnt"`
}

type Translation struct {
	Group     *string           `db:"group" json:"group"`
	ItemsJson pgtype.JSONB      `db:"items" json:"-"`
	Items     []TranslationItem `db:"-" json:"items"`
}

type TranslationItem struct {
	Name     string   `json:"name"`
	EngName  string   `json:"engName"`
	Alias    string   `json:"alias"`
	ForLevel *string  `json:"forLevel"`
	Order    int      `json:"order"`
	Helpers  []Helper `json:"helpers"`
}

type Helper struct {
	Name   string  `json:"name"`
	Alias  *string `json:"alias"`
	IsMain *bool   `json:"isMain"`
}

type Books struct {
	Alias        string `db:"alias" json:"alias"`
	Name         string `db:"name" json:"name"`
	Order        int    `db:"order" json:"order"`
	Abbreviation string `db:"abbreviation" json:"abbreviation"`
}

type BotBook struct {
	Id   string `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

type BookInfo struct {
	RacesJson              pgtype.JSONB             `db:"races_json" json:"-"`
	Races                  []NameAlias              `db:"-" json:"races"`
	ClassesJson            pgtype.JSONB             `db:"classes_json" json:"-"`
	Classes                []NameAlias              `db:"-" json:"classes"`
	ArchetypesJson         pgtype.JSONB             `db:"archetypes_json" json:"-"`
	Archetypes             []BookInfoWithClassAlias `db:"-" json:"archetypes"`
	FeatsJson              pgtype.JSONB             `db:"feats_json" json:"-"`
	Feats                  []NameAlias              `db:"-" json:"feats"`
	PrestigeClassesJson    pgtype.JSONB             `db:"prestige_classes_json" json:"-"`
	PrestigeClasses        []NameAlias              `db:"-" json:"prestigeClasses"`
	TraitsJson             pgtype.JSONB             `db:"traits_json" json:"-"`
	Traits                 []NameAlias              `db:"-" json:"traits"`
	GodsJson               pgtype.JSONB             `db:"gods_json" json:"-"`
	Gods                   []NameAlias              `db:"-" json:"gods"`
	DomainsJson            pgtype.JSONB             `db:"domains_json" json:"-"`
	Domains                []NameAlias              `db:"-" json:"domains"`
	SubdomainsJson         pgtype.JSONB             `db:"subdomains_json" json:"-"`
	Subdomains             []NameAlias              `db:"-" json:"subdomains"`
	InquisitionsJson       pgtype.JSONB             `db:"inquisitions_json" json:"-"`
	Inquisitions           []NameAlias              `db:"-" json:"inquisitions"`
	BloodlinesJson         pgtype.JSONB             `db:"bloodlines_json" json:"-"`
	Bloodlines             []BookInfoWithClassAlias `db:"-" json:"bloodlines"`
	SchoolsJson            pgtype.JSONB             `db:"schools_json" json:"-"`
	Schools                []NameAlias              `db:"-" json:"schools"`
	SpellsJson             pgtype.JSONB             `db:"spells_json" json:"-"`
	Spells                 []NameAlias              `db:"-" json:"spells"`
	WeaponsJson            pgtype.JSONB             `db:"weapons_json" json:"-"`
	Weapons                []NameAlias              `db:"-" json:"weapons"`
	ArmorsJson             pgtype.JSONB             `db:"armors_json" json:"-"`
	Armors                 []NameAlias              `db:"-" json:"armors"`
	EquipmentsJson         pgtype.JSONB             `db:"equipments_json" json:"-"`
	Equipments             []NameAlias              `db:"-" json:"equipments"`
	MagicItemAbilitiesJson pgtype.JSONB             `db:"magic_item_abilities_json" json:"-"`
	MagicItemAbilities     []BookMagicItemAbility   `db:"-" json:"magicItemAbilities"`
	MagicItemsJson         pgtype.JSONB             `db:"magic_items_json" json:"-"`
	MagicItems             []NameAlias              `db:"-" json:"magicItems"`
	MonsterAbilitiesJson   pgtype.JSONB             `db:"monster_abilities_json" json:"-"`
	MonsterAbilities       []NameAlias              `db:"-" json:"monsterAbilities"`
	BeastsJson             pgtype.JSONB             `db:"beasts_json" json:"-"`
	Beasts                 []NameAlias              `db:"-" json:"beasts"`
	AfflictionsJson        pgtype.JSONB             `db:"afflictions_json" json:"-"`
	Afflictions            []NameAlias              `db:"-" json:"afflictions"`
	TrapsJson              pgtype.JSONB             `db:"traps_json" json:"-"`
	Traps                  []NameAlias              `db:"-" json:"traps"`
	HauntsJson             pgtype.JSONB             `db:"haunts_json" json:"-"`
	Haunts                 []NameAlias              `db:"-" json:"haunts"`
}

type BookMagicItemAbility struct {
	Alias string   `json:"alias"`
	Name  string   `json:"name"`
	Types []string `json:"types"`
}

type BookInfoWithClassAlias struct {
	Alias      string `json:"alias"`
	Name       string `json:"name"`
	ClassAlias string `json:"classAlias"`
}

type Race struct {
	Alias    string       `db:"alias" json:"alias"`
	Name     string       `db:"name" json:"name"`
	BookJson pgtype.JSONB `db:"book_json" json:"-"`
	Book     Books        `db:"-" json:"book"`
}

type RaceInfo struct {
	Description           *string      `db:"description" json:"description"`
	PhysicalDescription   *string      `db:"physical_description" json:"physicalDescription"`
	Society               *string      `db:"society" json:"society"`
	Relations             *string      `db:"relations" json:"relations"`
	AlignmentAndReligion  *string      `db:"alignment_and_religion" json:"alignmentAndReligion"`
	Adventurers           *string      `db:"adventurers" json:"adventurers"`
	NamesDescription      *string      `db:"names_description" json:"namesDescription"`
	AdditionalDescription *string      `db:"additional_description" json:"additionalDescription"`
	NamesJson             pgtype.JSONB `db:"names" json:"-"`
	Names                 []RaceName   `db:"-" json:"names"`
	BaseRaceTraitsJson    pgtype.JSONB `db:"base_race_traits_json" json:"-"`
	BaseRaceTraits        []RaceTrait  `db:"-" json:"baseRaceTraits"`
	AlterRaceTraitsJson   pgtype.JSONB `db:"alter_race_traits_json" json:"-"`
	AlterRaceTraits       []RaceTrait  `db:"-" json:"alterRaceTraits"`
	FavoredClassJson      pgtype.JSONB `db:"favored_class_json" json:"-"`
	FavoredClass          []RaceClass  `db:"-" json:"favoredClass"`
	AdventurerClassJson   pgtype.JSONB `db:"adventurer_class_json" json:"-"`
	AdventurerClass       []RaceClass  `db:"-" json:"adventurerClass"`
	BookJson              pgtype.JSONB `db:"book_json" json:"-"`
	Book                  Books        `db:"-" json:"book"`
	HelpersJson           pgtype.JSONB `db:"helpers_json" json:"-"`
	Helpers               []Helper     `db:"-" json:"helpers"`
}

type RaceName struct {
	Name   string `json:"name"`
	IsMale *bool  `json:"isMale"`
}

type RaceTrait struct {
	Name           string   `json:"name"`
	Description    string   `json:"description"`
	BaseRaceTraits []string `json:"baseRaceTraits"`
}

type RaceClass struct {
	Name        string `json:"name"`
	Alias       string `json:"alias"`
	Description string `json:"description"`
}

type Class struct {
	Alias       string       `db:"alias" json:"alias"`
	Name        string       `db:"name" json:"name"`
	Description string       `db:"description" json:"description"`
	BookJson    pgtype.JSONB `db:"book_json" json:"-"`
	Book        Books        `db:"-" json:"book"`
}

type ClassInfo struct {
	Name               string       `db:"name" json:"name"`
	FullDescription    *string      `db:"full_description" json:"fullDescription"`
	Role               *string      `db:"role" json:"role"`
	Alignment          string       `db:"alignment" json:"alignment"`
	HitDie             int          `db:"hit_die" json:"hitDie"`
	StartingWealth     string       `db:"starting_wealth" json:"startingWealth"`
	SkillRanksPerLevel int          `db:"skill_ranks_per_level" json:"skillRanksPerLevel"`
	TableFeatures      string       `db:"table_features" json:"tableFeatures"`
	TableSpellCount    *string      `db:"table_spell_count" json:"tableSpellCount"`
	Features           string       `db:"features" json:"features"`
	IsHaveArchetypes   bool         `db:"is_have_archetypes" json:"isHaveArchetypes"`
	InfoLinks          *string      `db:"info_links" json:"infoLinks"`
	SkillsJson         pgtype.JSONB `db:"skills" json:"-"`
	Skills             []Skill      `db:"-" json:"skills"`
	ParentClassesJson  pgtype.JSONB `db:"parent_classes" json:"-"`
	ParentClasses      []NameAlias  `db:"-" json:"parentClasses"`
	BookJson           pgtype.JSONB `db:"book_json" json:"-"`
	Book               Books        `db:"-" json:"book"`
	HelpersJson        pgtype.JSONB `db:"helpers_json" json:"-"`
	Helpers            []Helper     `db:"-" json:"helpers"`
}

type NameAlias struct {
	Name  string `db:"name" json:"name"`
	Alias string `db:"alias" json:"alias"`
}

type Skill struct {
	Name    string  `json:"name"`
	Alias   string  `json:"alias"`
	Ability Ability `json:"ability"`
}

type Ability struct {
	Name                string `json:"name"`
	Alias               string `json:"alias"`
	ShortName           string `json:"shortName"`
	IsSkillArmorPenalty bool   `json:"isSkillArmorPenalty"`
}

type ClassWithArchetypes struct {
	Name                  string       `db:"name" json:"name"`
	Alias                 string       `db:"alias" json:"alias"`
	ArchetypesDescription string       `db:"archetypes_description" json:"archetypesDescription"`
	ArchetypesJson        pgtype.JSONB `db:"archetypes" json:"-"`
	Archetypes            []Archetype  `db:"-" json:"archetypes"`
}

type Archetype struct {
	Name        string `json:"name"`
	Alias       string `json:"alias"`
	Description string `json:"description"`
	Book        Books  `json:"book"`
}

type ArchetypeInfo struct {
	Name               string       `db:"name" json:"name"`
	Description        string       `db:"description" json:"description"`
	ArchetypeFeatures  string       `db:"archetype_features" json:"archetypeFeatures"`
	Alignment          string       `db:"alignment" json:"alignment"`
	HitDie             int          `db:"hit_die" json:"hitDie"`
	StartingWealth     string       `db:"starting_wealth" json:"startingWealth"`
	SkillRanksPerLevel int          `db:"skill_ranks_per_level" json:"skillRanksPerLevel"`
	TableFeatures      string       `db:"table_features" json:"tableFeatures"`
	TableSpellCount    *string      `db:"table_spell_count" json:"tableSpellCount"`
	Features           string       `db:"features" json:"features"`
	InfoLinks          *string      `db:"info_links" json:"infoLinks"`
	SkillsJson         pgtype.JSONB `db:"skills" json:"-"`
	Skills             []Skill      `db:"-" json:"skills"`
	ParentClassJson    pgtype.JSONB `db:"parent_class" json:"-"`
	ParentClass        NameAlias    `db:"-" json:"parentClass"`
	BookJson           pgtype.JSONB `db:"book_json" json:"-"`
	Book               Books        `db:"-" json:"book"`
	HelpersJson        pgtype.JSONB `db:"helpers_json" json:"-"`
	Helpers            []Helper     `db:"-" json:"helpers"`
}

type PrestigeClassInfo struct {
	Name               string       `db:"name" json:"name"`
	FullDescription    *string      `db:"full_description" json:"fullDescription"`
	Role               *string      `db:"role" json:"role"`
	Alignment          *string      `db:"alignment" json:"alignment"`
	HitDie             int          `db:"hit_die" json:"hitDie"`
	Requirements       string       `db:"requirements" json:"requirements"`
	SkillRanksPerLevel int          `db:"skill_ranks_per_level" json:"skillRanksPerLevel"`
	TableFeatures      string       `db:"table_features" json:"tableFeatures"`
	TableSpellCount    *string      `db:"table_spell_count" json:"tableSpellCount"`
	Features           string       `db:"features" json:"features"`
	SkillsJson         pgtype.JSONB `db:"skills" json:"-"`
	Skills             []Skill      `db:"-" json:"skills"`
	BookJson           pgtype.JSONB `db:"book_json" json:"-"`
	Book               Books        `db:"-" json:"book"`
	HelpersJson        pgtype.JSONB `db:"helpers_json" json:"-"`
	Helpers            []Helper     `db:"-" json:"helpers"`
}

type NpcInfo struct {
	Name               string       `db:"name" json:"name"`
	Alignment          string       `db:"alignment" json:"alignment"`
	HitDie             int          `db:"hit_die" json:"hitDie"`
	SkillRanksPerLevel int          `db:"skill_ranks_per_level" json:"skillRanksPerLevel"`
	TableFeatures      string       `db:"table_features" json:"tableFeatures"`
	TableSpellCount    *string      `db:"table_spell_count" json:"tableSpellCount"`
	Features           string       `db:"features" json:"features"`
	SkillsJson         pgtype.JSONB `db:"skills" json:"-"`
	Skills             []Skill      `db:"-" json:"skills"`
	BookJson           pgtype.JSONB `db:"book_json" json:"-"`
	Book               Books        `db:"-" json:"book"`
	HelpersJson        pgtype.JSONB `db:"helpers_json" json:"-"`
	Helpers            []Helper     `db:"-" json:"helpers"`
}

type SkillWithClasses struct {
	Name                string          `db:"name" json:"name"`
	Alias               string          `db:"alias" json:"alias"`
	ClassesJson         pgtype.JSONB    `db:"classes" json:"-"`
	Classes             []ClassForSkill `db:"-" json:"classes"`
	PrestigeClassesJson pgtype.JSONB    `db:"prestige_classes" json:"-"`
	PrestigeClasses     []ClassForSkill `db:"-" json:"prestigeClasses"`
}

type ClassForSkill struct {
	Name         string `json:"name"`
	Alias        string `json:"alias"`
	ShortName    string `json:"shortName"`
	Book         Books  `json:"book"`
	IsClassSkill bool   `json:"isClassSkill"`
}

type SkillsPerLvlInfo struct {
	Name               string `db:"name" json:"name"`
	Alias              string `db:"alias" json:"alias"`
	SkillRanksPerLevel int    `db:"skill_ranks_per_level" json:"skillRanksPerLevel"`
	IsPrestige         bool   `db:"is_prestige" json:"isPrestige"`
}

type SkillInfo struct {
	Description     string       `db:"description" json:"description"`
	FullDescription string       `db:"full_description" json:"fullDescription"`
	AbilityJson     pgtype.JSONB `db:"ability" json:"-"`
	Ability         Ability      `db:"-" json:"ability"`
}

type Feat struct {
	Id            int          `db:"id" json:"id"`
	Name          string       `db:"name" json:"name"`
	Alias         string       `db:"alias" json:"alias"`
	Requirements  *string      `db:"requirements" json:"requirements"`
	Description   string       `db:"description" json:"description"`
	ParentFeatId  *int         `db:"parent_feat_id" json:"parentFeatId"`
	FeatTypesJson pgtype.JSONB `db:"feat_types" json:"-"`
	FeatTypes     []NameAlias  `db:"-" json:"types"`
	BookJson      pgtype.JSONB `db:"book" json:"-"`
	Book          Books        `db:"-" json:"book"`
}

type FeatInfo struct {
	Name            string       `db:"name" json:"name"`
	FullDescription *string      `db:"full_description" json:"fullDescription"`
	Prerequisites   *string      `db:"prerequisites" json:"prerequisites"`
	Benefit         string       `db:"benefit" json:"benefit"`
	Normal          *string      `db:"normal" json:"normal"`
	Special         *string      `db:"special" json:"special"`
	FeatTypesJson   pgtype.JSONB `db:"feat_types" json:"-"`
	FeatTypes       []NameAlias  `db:"-" json:"types"`
	BookJson        pgtype.JSONB `db:"book" json:"-"`
	Book            Books        `db:"-" json:"book"`
	HelpersJson     pgtype.JSONB `db:"helpers" json:"-"`
	Helpers         []Helper     `db:"-" json:"helpers"`
}

type Trait struct {
	Alias         string       `db:"alias" json:"alias"`
	Name          string       `db:"name" json:"name"`
	EngName       string       `db:"eng_name" json:"engName"`
	Benefit       string       `db:"benefit" json:"benefit"`
	Prerequisites *string      `db:"prerequisites" json:"prerequisites"`
	TraitTypeJson pgtype.JSONB `db:"trait_type" json:"-"`
	TraitType     TraitType    `db:"-" json:"type"`
	BookJson      pgtype.JSONB `db:"book" json:"-"`
	Book          Books        `db:"-" json:"book"`
}

type TraitType struct {
	Name        string     `json:"name"`
	Alias       string     `json:"alias"`
	Description *string    `json:"description"`
	ParentType  *TraitType `json:"parentType"`
}

type God struct {
	Alias         string               `db:"alias" json:"alias"`
	Name          string               `db:"name" json:"name"`
	EngName       string               `db:"eng_name" json:"engName"`
	Aligment      string               `db:"aligment" json:"aligment"`
	Title         string               `db:"title" json:"title"`
	Portfolios    string               `db:"portfolios" json:"portfolios"`
	FavoredWeapon string               `db:"favored_weapon" json:"favoredWeapon"`
	Symbol        *string              `db:"symbol" json:"symbol"`
	SacredAnimal  *string              `db:"sacred_animal" json:"sacredAnimal"`
	SacredColors  *string              `db:"sacred_colors" json:"sacredColors"`
	GodTypeJson   pgtype.JSONB         `db:"god_type" json:"-"`
	GodType       NameAliasDescription `db:"-" json:"type"`
	BookJson      pgtype.JSONB         `db:"book" json:"-"`
	Book          Books                `db:"-" json:"book"`
	DomainsJson   pgtype.JSONB         `db:"domains" json:"-"`
	Domains       []GodDomain          `db:"-" json:"domains"`
}

type NameAliasDescription struct {
	Name        string  `json:"name"`
	Alias       string  `json:"alias"`
	Description *string `json:"description"`
}

type GodDomain struct {
	Alias string `json:"alias"`
	Name  string `json:"name"`
	Type  string `json:"type"`
}

type Domain struct {
	Name        string       `db:"name" json:"name"`
	Alias       string       `db:"alias" json:"alias"`
	Description *string      `db:"description" json:"description"`
	Type        string       `db:"domain_type" json:"type"`
	BookJson    pgtype.JSONB `db:"book" json:"-"`
	Book        Books        `db:"-" json:"book"`
	ChildsJson  pgtype.JSONB `db:"child_domains" json:"-"`
	Childs      []Domain     `db:"-" json:"childs"`
}

type DomainInfo struct {
	Name              string           `db:"name" json:"name"`
	Description       *string          `db:"description" json:"description"`
	InnerDescription  *string          `db:"inner_description" json:"innerDescription"`
	Power0Name        *string          `db:"power0_name" json:"power0Name"`
	Power0Description *string          `db:"power0_description" json:"power0Description"`
	Power1Name        *string          `db:"power1_name" json:"power1Name"`
	Power1Description *string          `db:"power1_description" json:"power1Description"`
	Power2Name        *string          `db:"power2_name" json:"power2Name"`
	Power2Description *string          `db:"power2_description" json:"power2Description"`
	GodsArray         pgtype.TextArray `db:"gods" json:"-"`
	Gods              []string         `db:"-" json:"gods"`
	BookJson          pgtype.JSONB     `db:"book" json:"-"`
	Book              Books            `db:"-" json:"book"`
	ChildsJson        pgtype.JSONB     `db:"childs" json:"-"`
	Childs            []NameAlias      `db:"-" json:"childs"`
	ParentsJson       pgtype.JSONB     `db:"parents" json:"-"`
	Parents           []DomainParent   `db:"-" json:"parents"`
	SpellsJson        pgtype.JSONB     `db:"spells" json:"-"`
	Spells            []DomainSpell    `db:"-" json:"spells"`
}

type DomainSpell struct {
	Alias   string  `json:"alias"`
	Name    string  `json:"name"`
	Level   int     `json:"level"`
	Comment *string `json:"comment"`
}

type DomainParent struct {
	Alias             string        `json:"alias"`
	Name              string        `json:"name"`
	Power1Name        *string       `json:"power1Name"`
	Power1Description *string       `json:"power1Description"`
	Power2Name        *string       `json:"power2Name"`
	Power2Description *string       `json:"power2Description"`
	Spells            []DomainSpell `json:"spells"`
}

type Bloodline struct {
	Name        string       `db:"name" json:"name"`
	Alias       string       `db:"alias" json:"alias"`
	Description string       `db:"description" json:"description"`
	BookJson    pgtype.JSONB `db:"book" json:"-"`
	Book        Books        `db:"-" json:"book"`
}

type BloodlineInfo struct {
	Name            string       `db:"name" json:"name"`
	FullDescription string       `db:"full_description" json:"fullDescription"`
	Description     string       `db:"description" json:"description"`
	BookJson        pgtype.JSONB `db:"book" json:"-"`
	Book            Books        `db:"-" json:"book"`
}

type Order struct {
	Name        string `db:"name" json:"name"`
	Alias       string `db:"alias" json:"alias"`
	Description string `db:"description" json:"description"`
	ClassName   string `db:"class_name" json:"className"`
}

type OrderInfo struct {
	Name            string `db:"name" json:"name"`
	FullDescription string `db:"full_description" json:"fullDescription"`
	Description     string `db:"description" json:"description"`
}

type School struct {
	Name        string               `db:"name" json:"name"`
	Alias       string               `db:"alias" json:"alias"`
	Description *string              `db:"description" json:"description"`
	TypeJson    pgtype.JSONB         `db:"school_type" json:"-"`
	Type        NameAliasDescription `db:"-" json:"type"`
	BookJson    pgtype.JSONB         `db:"book" json:"-"`
	Book        Books                `db:"-" json:"book"`
}

type SchoolInfo struct {
	Name            string               `db:"name" json:"name"`
	Alias           string               `db:"alias" json:"alias"`
	Description     *string              `db:"description" json:"description"`
	FullDescription string               `db:"full_description" json:"fullDescription"`
	TypeJson        pgtype.JSONB         `db:"school_type" json:"-"`
	Type            NameAliasDescription `db:"-" json:"type"`
	BookJson        pgtype.JSONB         `db:"book" json:"-"`
	Book            Books                `db:"-" json:"book"`
}

type Spell struct {
	Name                       string              `db:"name" json:"name"`
	Alias                      string              `db:"alias" json:"alias"`
	ShortDescriptionComponents *string             `db:"short_description_components" json:"shortDescriptionComponents"`
	ShortDescription           *string             `db:"short_description" json:"shortDescription"`
	SchoolsJson                pgtype.JSONB        `db:"schools" json:"-"`
	Schools                    []SchoolForList     `db:"-" json:"schools"`
	ClassesJson                pgtype.JSONB        `db:"classes" json:"-"`
	Classes                    []ClassForSpellList `db:"-" json:"classes"`
	BookJson                   pgtype.JSONB        `db:"book" json:"-"`
	Book                       Books               `db:"-" json:"book"`
}

type SchoolForList struct {
	Name  string    `json:"name"`
	Alias string    `json:"alias"`
	Type  NameAlias `json:"type"`
}

type ClassForSpellList struct {
	Name  string `json:"name"`
	Alias string `json:"alias"`
	Level int    `json:"level"`
}

type SpellInfo struct {
	Name            string              `db:"name" json:"name"`
	EngName         *string             `db:"eng_name" json:"engName"`
	CastingTime     string              `db:"casting_time" json:"castingTime"`
	Components      *string             `db:"components" json:"components"`
	Description     string              `db:"description" json:"description"`
	Range           *string             `db:"range" json:"range"`
	Target          *string             `db:"target" json:"target"`
	Area            *string             `db:"area" json:"area"`
	Effect          *string             `db:"effect" json:"effect"`
	Duration        string              `db:"duration" json:"duration"`
	SavingThrow     *string             `db:"saving_throw" json:"savingThrow"`
	SpellResistance *int                `db:"spell_resistance" json:"spellResistance"`
	SubSchool       *string             `db:"sub_school" json:"subSchool"`
	SchoolJson      pgtype.JSONB        `db:"school" json:"-"`
	School          SchoolForList       `db:"-" json:"school"`
	ClassesJson     pgtype.JSONB        `db:"classes" json:"-"`
	Classes         []ClassForSpellList `db:"-" json:"classes"`
	RacesJson       pgtype.JSONB        `db:"races" json:"-"`
	Races           []NameAlias         `db:"-" json:"races"`
	BookJson        pgtype.JSONB        `db:"book" json:"-"`
	Book            Books               `db:"-" json:"book"`
	GodJson         pgtype.JSONB        `db:"god" json:"-"`
	God             *NameAlias          `db:"-" json:"god"`
	HelpersJson     pgtype.JSONB        `db:"helpers" json:"-"`
	Helpers         []Helper            `db:"-" json:"helpers"`
}

type ClassForBotSpellList struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Alias string `json:"alias"`
	Level int    `json:"level"`
}

type BookForBotSpellList struct {
	Id           int    `db:"id" json:"id"`
	Alias        string `db:"alias" json:"alias"`
	Name         string `db:"name" json:"name"`
	Order        int    `db:"order" json:"order"`
	Abbreviation string `db:"abbreviation" json:"abbreviation"`
}

type BotSpellInfo struct {
	Id                         int                    `db:"id" json:"id"`
	Alias                      string                 `db:"alias" json:"alias"`
	Name                       string                 `db:"name" json:"name"`
	EngName                    *string                `db:"eng_name" json:"engName"`
	ShortDescriptionComponents *string                `db:"short_description_components" json:"shortDescriptionComponents,omitempty"`
	ShortDescription           *string                `db:"short_description" json:"shortDescription,omitempty"`
	CastingTime                *string                `db:"casting_time" json:"castingTime,omitempty"`
	Components                 *string                `db:"components" json:"components,omitempty"`
	Description                *string                `db:"description" json:"description,omitempty"`
	Range                      *string                `db:"range" json:"range,omitempty"`
	Target                     *string                `db:"target" json:"target,omitempty"`
	Area                       *string                `db:"area" json:"area,omitempty"`
	Effect                     *string                `db:"effect" json:"effect,omitempty"`
	Duration                   *string                `db:"duration" json:"duration,omitempty"`
	SavingThrow                *string                `db:"saving_throw" json:"savingThrow,omitempty"`
	SpellResistance            *int                   `db:"spell_resistance" json:"spellResistance,omitempty"`
	SubSchool                  *string                `db:"sub_school" json:"subSchool,omitempty"`
	SchoolJson                 pgtype.JSONB           `db:"school" json:"-"`
	School                     *SchoolForList         `db:"-" json:"school,omitempty"`
	ClassesJson                pgtype.JSONB           `db:"classes" json:"-"`
	Classes                    []ClassForBotSpellList `db:"-" json:"classes"`
	RacesJson                  pgtype.JSONB           `db:"races" json:"-"`
	Races                      []NameAlias            `db:"-" json:"races,omitempty"`
	BookJson                   pgtype.JSONB           `db:"book" json:"-"`
	Book                       *BookForBotSpellList   `db:"-" json:"book,omitempty"`
	GodJson                    pgtype.JSONB           `db:"god" json:"-"`
	God                        *NameAlias             `db:"-" json:"god,omitempty"`
	HelpersJson                pgtype.JSONB           `db:"helpers" json:"-"`
	Helpers                    []Helper               `db:"-" json:"helpers,omitempty"`
}

type BotClassInfo struct {
	Id              int              `db:"id" json:"id"`
	Alias           string           `db:"alias" json:"alias"`
	Name            string           `db:"name" json:"name"`
	Description     string           `db:"description" json:"description"`
	TableFeatures   *string          `db:"table_features" json:"tableFeatures,omitempty"`
	TableSpellCount *string          `db:"table_spell_count" json:"tableSpellCount,omitempty"`
	SpellLevelsDB   pgtype.Int4Array `db:"spell_levels" json:"-"`
	SpellLevels     []int            `db:"-" json:"spellLevels"`
}

type Weapon struct {
	Alias                   string           `db:"alias" json:"alias"`
	Name                    string           `db:"name" json:"name"`
	EngName                 *string          `db:"eng_name" json:"engName"`
	Type                    *string          `db:"type" json:"type"`
	Cost                    *float64         `db:"cost" json:"cost"`
	DamageS                 *string          `db:"damage_s" json:"damageS"`
	DamageM                 *string          `db:"damage_m" json:"damageM"`
	CriticalRoll            *string          `db:"critical_roll" json:"criticalRoll"`
	CriticalDamage          *string          `db:"critical_damage" json:"criticalDamage"`
	Range                   *string          `db:"range" json:"range"`
	Misfire                 *string          `db:"misfire" json:"misfire"`
	Capacity                *int             `db:"capacity" json:"capacity"`
	Weight                  *float64         `db:"weight" json:"weight"`
	Special                 *string          `db:"special" json:"special"`
	Description             *string          `db:"description" json:"description"`
	ProficientCategoryJson  pgtype.JSONB     `db:"proficient_category" json:"-"`
	ProficientCategory      *NameAlias       `db:"-" json:"proficientCategory"`
	RangeCategoryJson       pgtype.JSONB     `db:"range_category" json:"-"`
	RangeCategory           *NameAlias       `db:"-" json:"rangeCategory"`
	EncumbranceCategoryJson pgtype.JSONB     `db:"encumbrance_category" json:"-"`
	EncumbranceCategory     *NameAlias       `db:"-" json:"encumbranceCategory"`
	ParentsArray            pgtype.TextArray `db:"parents" json:"-"`
	Parents                 []string         `db:"-" json:"parents"`
	BookJson                pgtype.JSONB     `db:"book" json:"-"`
	Book                    Books            `db:"-" json:"book"`
	ChildsJson              pgtype.JSONB     `db:"childs" json:"-"`
	Childs                  []Weapon         `db:"-" json:"childs"`
}

type Armor struct {
	Alias                    string       `db:"alias" json:"alias"`
	Name                     string       `db:"name" json:"name"`
	EngName                  *string      `db:"eng_name" json:"engName"`
	ArmorBonus               *int         `db:"armor_bonus" json:"armorBonus"`
	MaxDexBonus              *int         `db:"max_dex_bonus" json:"maxDexBonus"`
	ArmorCheckPenalty        *int         `db:"armor_check_penalty" json:"armorCheckPenalty"`
	Cost                     *float64     `db:"cost" json:"cost"`
	ArcaneSpellFailureChance *int         `db:"arcane_spell_failure_chance" json:"arcaneSpellFailureChance"`
	Speed30                  *int         `db:"speed30" json:"speed30"`
	Speed20                  *int         `db:"speed20" json:"speed20"`
	Weight                   *float64     `db:"weight" json:"weight"`
	Description              *string      `db:"description" json:"description"`
	TypeJson                 pgtype.JSONB `db:"armor_type" json:"-"`
	Type                     NameAlias    `db:"-" json:"type"`
	BookJson                 pgtype.JSONB `db:"book" json:"-"`
	Book                     Books        `db:"-" json:"book"`
}

type GoodAndService struct {
	Alias             string                `db:"alias" json:"alias"`
	Name              string                `db:"name" json:"name"`
	EngName           *string               `db:"eng_name" json:"engName"`
	EquipmentSubType  *string               `db:"equipment_sub_type" json:"equipmentSubType"`
	Cost              *float64              `db:"cost" json:"cost"`
	CostSpecial       *string               `db:"cost_special" json:"costSpecial"`
	CostDescription   *string               `db:"cost_description" json:"costDescription"`
	CostOfPassage     *float64              `db:"cost_of_passage" json:"costOfPassage"`
	Weight            *float64              `db:"weight" json:"weight"`
	WeightSpecial     *string               `db:"weight_special" json:"weightSpecial"`
	WeightDescription *string               `db:"weight_description" json:"weightDescription"`
	CraftDc           *int                  `db:"craft_dc" json:"craftDc"`
	Description       *string               `db:"description" json:"description"`
	TypeJson          pgtype.JSONB          `db:"equipment_type" json:"-"`
	Type              *NameAliasDescription `db:"-" json:"type"`
	ParentsArray      pgtype.TextArray      `db:"parents" json:"-"`
	Parents           []string              `db:"-" json:"parents"`
	BookJson          pgtype.JSONB          `db:"book" json:"-"`
	Book              Books                 `db:"-" json:"book"`
	ChildsJson        pgtype.JSONB          `db:"childs" json:"-"`
	Childs            []GoodAndService      `db:"-" json:"childs"`
}

type MagicItemInfo struct {
	Alias                    string       `db:"alias" json:"alias"`
	Name                     string       `db:"name" json:"name"`
	EngName                  *string      `db:"eng_name" json:"engName"`
	Aura                     string       `db:"aura" json:"aura"`
	Cl                       *int         `db:"cl" json:"cl"`
	SlotComment              *string      `db:"slot_comment" json:"slotComment"`
	Price                    *float64     `db:"price" json:"price"`
	PriceComment             *string      `db:"price_comment" json:"priceComment"`
	Weight                   *float64     `db:"weight" json:"weight"`
	WeightComment            *string      `db:"weight_comment" json:"weightComment"`
	Description              string       `db:"description" json:"description"`
	ConstructionRequirements *string      `db:"construction_requirements" json:"constructionRequirements"`
	ConstructionCost         *float64     `db:"construction_cost" json:"constructionCost"`
	CreationMagicItems       *string      `db:"creation_magic_items" json:"creationMagicItems"`
	Statistics               *string      `db:"statistics" json:"statistics"`
	Destruction              *string      `db:"destruction" json:"destruction"`
	SlotJson                 pgtype.JSONB `db:"slot" json:"-"`
	Slot                     *NameAlias   `db:"-" json:"slot"`
	TypeJson                 pgtype.JSONB `db:"magic_item_type" json:"-"`
	Type                     NameAlias    `db:"-" json:"type"`
	BookJson                 pgtype.JSONB `db:"book" json:"-"`
	Book                     Books        `db:"-" json:"book"`
	HelpersJson              pgtype.JSONB `db:"helpers" json:"-"`
	Helpers                  []Helper     `db:"-" json:"helpers"`
}

type MagicItemForList struct {
	Alias    string       `db:"alias" json:"alias"`
	Name     string       `db:"name" json:"name"`
	EngName  *string      `db:"eng_name" json:"engName"`
	SlotJson pgtype.JSONB `db:"slot" json:"-"`
	Slot     *NameAlias   `db:"-" json:"slot"`
	TypeJson pgtype.JSONB `db:"magic_item_type" json:"-"`
	Type     NameAlias    `db:"-" json:"type"`
	BookJson pgtype.JSONB `db:"book" json:"-"`
	Book     Books        `db:"-" json:"book"`
}

type MagicItemAbility struct {
	Alias                    string       `db:"alias" json:"alias"`
	Name                     string       `db:"name" json:"name"`
	EngName                  *string      `db:"eng_name" json:"engName"`
	Aura                     string       `db:"aura" json:"aura"`
	Cl                       *int         `db:"cl" json:"cl"`
	BonusPrice               *int         `db:"bonus_price" json:"bonusPrice"`
	MoneyPrice               *float64     `db:"money_price" json:"moneyPrice"`
	Description              string       `db:"description" json:"description"`
	ConstructionRequirements *string      `db:"construction_requirements" json:"constructionRequirements"`
	BookJson                 pgtype.JSONB `db:"book" json:"-"`
	Book                     Books        `db:"-" json:"book"`
}

type Beast struct {
	Id               int              `db:"id" json:"id"`
	Alias            string           `db:"alias" json:"alias"`
	Name             string           `db:"name" json:"name"`
	EngName          *string          `db:"eng_name" json:"engName"`
	Cr               *float64         `db:"cr" json:"cr"`
	Description      *string          `db:"description" json:"description"`
	Exp              *int             `db:"exp" json:"exp"`
	FullCreatureType *string          `db:"full_creature_type" json:"fullCreatureType"`
	Names            *string          `db:"names" json:"names"`
	IsUnique         *bool            `db:"is_unique" json:"isUnique"`
	ClimateJson      pgtype.JSONB     `db:"climate" json:"-"`
	Climate          *NameAlias       `db:"-" json:"climate"`
	CreatureTypeJson pgtype.JSONB     `db:"creature_type" json:"-"`
	CreatureType     *NameAlias       `db:"-" json:"creatureType"`
	TerrainJson      pgtype.JSONB     `db:"terrain" json:"-"`
	Terrain          *NameAlias       `db:"-" json:"terrain"`
	RolesJson        pgtype.JSONB     `db:"roles" json:"-"`
	Roles            []NameAlias      `db:"-" json:"roles"`
	BookJson         pgtype.JSONB     `db:"book" json:"-"`
	Book             Books            `db:"-" json:"book"`
	ChildsArray      pgtype.Int4Array `db:"childs" json:"-"`
	Childs           []int            `db:"-" json:"childs"`
}

type BeastInfo struct {
	Alias                        string                     `db:"alias" json:"alias"`
	Name                         string                     `db:"name" json:"name"`
	EngName                      *string                    `db:"eng_name" json:"engName"`
	Description                  *string                    `db:"description" json:"description"`
	Names                        *string                    `db:"names" json:"names"`
	RootPageDescription          *string                    `db:"root_page_description" json:"rootPageDescription"`
	Cr                           *float64                   `db:"cr" json:"cr"`
	Exp                          *int                       `db:"exp" json:"exp"`
	FullCreatureType             *string                    `db:"full_creature_type" json:"fullCreatureType"`
	Initiative                   *int                       `db:"initiative" json:"initiative"`
	Senses                       *string                    `db:"senses" json:"senses"`
	Perception                   *int                       `db:"perception" json:"perception"`
	PerceptionComment            *string                    `db:"perception_comment" json:"perceptionComment"`
	Aura                         *string                    `db:"aura" json:"aura"`
	Strength                     *int                       `db:"strength" json:"strength"`
	Dexterity                    *int                       `db:"dexterity" json:"dexterity"`
	Constitution                 *int                       `db:"constitution" json:"constitution"`
	Intelligence                 *int                       `db:"intelligence" json:"intelligence"`
	Wisdom                       *int                       `db:"wisdom" json:"wisdom"`
	Charisma                     *int                       `db:"charisma" json:"charisma"`
	MaxAcDexterity               *int                       `db:"max_ac_dexterity" json:"maxAcDexterity"`
	AcNatural                    *int                       `db:"ac_natural" json:"acNatural"`
	AcArmor                      *int                       `db:"ac_armor" json:"acArmor"`
	AcShield                     *int                       `db:"ac_shield" json:"acShield"`
	AcDodge                      *int                       `db:"ac_dodge" json:"acDodge"`
	AcDeflection                 *int                       `db:"ac_deflection" json:"acDeflection"`
	AcInsight                    *int                       `db:"ac_insight" json:"acInsight"`
	AcRage                       *int                       `db:"ac_rage" json:"acRage"`
	AcWisdom                     *int                       `db:"ac_wisdom" json:"acWisdom"`
	AcMonk                       *int                       `db:"ac_monk" json:"acMonk"`
	AcString                     *string                    `db:"ac_string" json:"acString"`
	AcDescription                *string                    `db:"ac_description" json:"acDescription"`
	HitPoints                    *int                       `db:"hit_points" json:"hitPoints"`
	HitPointsDescription         *string                    `db:"hit_points_description" json:"hitPointsDescription"`
	HitPointsComment             *string                    `db:"hit_points_comment" json:"hitPointsComment"`
	FastHealing                  *string                    `db:"fast_healing" json:"fastHealing"`
	Regeneration                 *string                    `db:"regeneration" json:"regeneration"`
	Fortitude                    *int                       `db:"fortitude" json:"fortitude"`
	Reflex                       *int                       `db:"reflex" json:"reflex"`
	Will                         *int                       `db:"will" json:"will"`
	WillComment                  *string                    `db:"will_comment" json:"willComment"`
	DefensiveAbilities           *string                    `db:"defensive_abilities" json:"defensiveAbilities"`
	DamageResist                 *string                    `db:"damage_resist" json:"damageResist"`
	Immune                       *string                    `db:"immune" json:"immune"`
	Resist                       *string                    `db:"resist" json:"resist"`
	SpellResist                  *int                       `db:"spell_resist" json:"spellResist"`
	SpellResistComment           *string                    `db:"spell_resist_comment" json:"spellResistComment"`
	Weaknesses                   *string                    `db:"weaknesses" json:"weaknesses"`
	Speed                        *string                    `db:"speed" json:"speed"`
	MeleeAttacks                 *string                    `db:"melee_attacks" json:"meleeAttacks"`
	RangedAttacks                *string                    `db:"ranged_attacks" json:"rangedAttacks"`
	Space                        *string                    `db:"space" json:"space"`
	Reach                        *string                    `db:"reach" json:"reach"`
	SpecialAttacks               *string                    `db:"special_attacks" json:"specialAttacks"`
	SpellLikeAbilities           *string                    `db:"spell_like_abilities" json:"spellLikeAbilities"`
	SpellsPrepared               *string                    `db:"spells_prepared" json:"spellsPrepared"`
	SpellsKnown                  *string                    `db:"spells_known" json:"spellsKnown"`
	Domains                      *string                    `db:"domains" json:"domains"`
	Patron                       *string                    `db:"patron" json:"patron"`
	Bloodline                    *string                    `db:"bloodline" json:"bloodline"`
	School                       *string                    `db:"school" json:"school"`
	OppositionSchools            *string                    `db:"opposition_schools" json:"oppositionSchools"`
	BeforeCombat                 *string                    `db:"before_combat" json:"beforeCombat"`
	DuringCombat                 *string                    `db:"during_combat" json:"duringCombat"`
	Morale                       *string                    `db:"morale" json:"morale"`
	BaseParameters               *string                    `db:"base_parameters" json:"baseParameters"`
	BaseAttack                   *int                       `db:"base_attack" json:"baseAttack"`
	CombatManeuverBonus          *int                       `db:"combat_maneuver_bonus" json:"combatManeuverBonus"`
	CombatManeuverBonusComment   *string                    `db:"combat_maneuver_bonus_comment" json:"combatManeuverBonusComment"`
	CombatManeuverDefense        *int                       `db:"combat_maneuver_defense" json:"combatManeuverDefense"`
	CombatManeuverDefenseComment *string                    `db:"combat_maneuver_defense_comment" json:"combatManeuverDefenseComment"`
	Feats                        *string                    `db:"feats" json:"feats"`
	Skills                       *string                    `db:"skills" json:"skills"`
	SkillsRacialModifiers        *string                    `db:"skills_racial_modifiers" json:"skillsRacialModifiers"`
	Languages                    *string                    `db:"languages" json:"languages"`
	CombatGear                   *string                    `db:"combat_gear" json:"combatGear"`
	OtherGear                    *string                    `db:"other_gear" json:"otherGear"`
	Gear                         *string                    `db:"gear" json:"gear"`
	Spellbook                    *string                    `db:"spellbook" json:"spellbook"`
	SpecialQualities             *string                    `db:"special_qualities" json:"specialQualities"`
	Environment                  *string                    `db:"environment" json:"environment"`
	Organization                 *string                    `db:"organization" json:"organization"`
	Treasure                     *string                    `db:"treasure" json:"treasure"`
	SpecialAbilities             *string                    `db:"special_abilities" json:"specialAbilities"`
	FullDescription              *string                    `db:"full_description" json:"fullDescription"`
	ConstructionDescription      *string                    `db:"construction_description" json:"constructionDescription"`
	ConstructionCasterLevel      *int                       `db:"construction_caster_level" json:"constructionCasterLevel"`
	ConstructionPrice            *float64                   `db:"construction_price" json:"constructionPrice"`
	ConstructionPriceComment     *string                    `db:"construction_price_comment" json:"constructionPriceComment"`
	ConstructionRequirements     *string                    `db:"construction_requirements" json:"constructionRequirements"`
	ConstructionSkill            *string                    `db:"construction_skill" json:"constructionSkill"`
	CostPrice                    *float64                   `db:"cost_price" json:"costPrice"`
	CostPriceComment             *string                    `db:"cost_price_comment" json:"costPriceComment"`
	ClimateJson                  pgtype.JSONB               `db:"climate" json:"-"`
	Climate                      *NameAlias                 `db:"-" json:"climate"`
	CreatureTypeJson             pgtype.JSONB               `db:"creature_type" json:"-"`
	CreatureType                 *NameAlias                 `db:"-" json:"creatureType"`
	TerrainJson                  pgtype.JSONB               `db:"terrain" json:"-"`
	Terrain                      *NameAlias                 `db:"-" json:"terrain"`
	AnimalCompanionJson          pgtype.JSONB               `db:"animal_companion" json:"-"`
	AnimalCompanion              []AnimalCompanionNameAlias `db:"-" json:"animalCompanion"`
	SizeTypeJson                 pgtype.JSONB               `db:"size_type" json:"-"`
	SizeType                     *NameAlias                 `db:"-" json:"sizeType"`
	ParentJson                   pgtype.JSONB               `db:"parent" json:"-"`
	Parent                       *NameAlias                 `db:"-" json:"parent"`
	RolesJson                    pgtype.JSONB               `db:"roles" json:"-"`
	Roles                        []NameAlias                `db:"-" json:"roles"`
	BookJson                     pgtype.JSONB               `db:"book" json:"-"`
	Book                         Books                      `db:"-" json:"book"`
	HelpersJson                  pgtype.JSONB               `db:"helpers" json:"-"`
	Helpers                      []Helper                   `db:"-" json:"helpers"`
	ChildsJson                   pgtype.JSONB               `db:"childs" json:"-"`
	Childs                       []NameAlias                `db:"-" json:"childs"`
}

type AnimalCompanionNameAlias struct {
	Name      string `db:"name" json:"name"`
	Alias     string `db:"alias" json:"alias"`
	TypeAlias string `db:"typeAlias" json:"typeAlias"`
	TypeName  string `db:"typeName" json:"typeName"`
}

type MonsterAbility struct {
	Alias       string       `db:"alias" json:"alias"`
	Name        string       `db:"name" json:"name"`
	EngName     string       `db:"eng_name" json:"engName"`
	Description string       `db:"description" json:"description"`
	Type        *string      `db:"type" json:"type"`
	Format      string       `db:"format" json:"format"`
	Location    string       `db:"location" json:"location"`
	BookJson    pgtype.JSONB `db:"book" json:"-"`
	Book        Books        `db:"-" json:"book"`
}

type CreatureType struct {
	Alias       string  `db:"alias" json:"alias"`
	Name        string  `db:"name" json:"name"`
	EngName     string  `db:"eng_name" json:"engName"`
	Description *string `db:"description" json:"description"`
	Features    *string `db:"features" json:"features"`
	Traits      *string `db:"traits" json:"traits"`
	IsSubtype   bool    `db:"is_subtype" json:"isSubtype"`
}

type AnimalCompanion struct {
	Type                        string       `db:"type" json:"type"`
	Alias                       string       `db:"alias" json:"alias"`
	Name                        string       `db:"name" json:"name"`
	EngName                     *string      `db:"eng_name" json:"engName"`
	Prerequisites               *string      `db:"prerequisites" json:"prerequisites"`
	Description                 *string      `db:"description" json:"description"`
	StartRacialSkillModifiers   *string      `db:"start_racial_skill_modifiers" json:"startRacialSkillModifiers"`
	UpdateRacialSkillModifiers  *string      `db:"update_racial_skill_modifiers" json:"updateRacialSkillModifiers"`
	StartSize                   *string      `db:"start_size" json:"startSize"`
	UpdateSize                  *string      `db:"update_size" json:"updateSize"`
	StartSpeed                  *string      `db:"start_speed" json:"startSpeed"`
	UpdateSpeed                 *string      `db:"update_speed" json:"updateSpeed"`
	StartAc                     *string      `db:"start_ac" json:"startAc"`
	UpdateAc                    *string      `db:"update_ac" json:"updateAc"`
	StartAttack                 *string      `db:"start_attack" json:"startAttack"`
	UpdateAttack                *string      `db:"update_attack" json:"updateAttack"`
	StartStrength               *int         `db:"start_strength" json:"startStrength"`
	UpdateStrength              *int         `db:"update_strength" json:"updateStrength"`
	StartDexterity              *int         `db:"start_dexterity" json:"startDexterity"`
	UpdateDexterity             *int         `db:"update_dexterity" json:"updateDexterity"`
	StartConstitution           *int         `db:"start_constitution" json:"startConstitution"`
	UpdateConstitution          *int         `db:"update_constitution" json:"updateConstitution"`
	StartIntelligence           *int         `db:"start_intelligence" json:"startIntelligence"`
	UpdateIntelligence          *int         `db:"update_intelligence" json:"updateIntelligence"`
	StartWisdom                 *int         `db:"start_wisdom" json:"startWisdom"`
	UpdateWisdom                *int         `db:"update_wisdom" json:"updateWisdom"`
	StartCharisma               *int         `db:"start_charisma" json:"startCharisma"`
	UpdateCharisma              *int         `db:"update_charisma" json:"updateCharisma"`
	StartLanguages              *string      `db:"start_languages" json:"startLanguages"`
	UpdateLanguages             *string      `db:"update_languages" json:"updateLanguages"`
	StartSpecialAttacks         *string      `db:"start_special_attacks" json:"startSpecialAttacks"`
	UpdateSpecialAttacks        *string      `db:"update_special_attacks" json:"updateSpecialAttacks"`
	StartSpecialQualities       *string      `db:"start_special_qualities" json:"startSpecialQualities"`
	UpdateSpecialQualities      *string      `db:"update_special_qualities" json:"updateSpecialQualities"`
	StartCombatManeuverDefense  *string      `db:"start_combat_maneuver_defense" json:"startCombatManeuverDefense"`
	UpdateCombatManeuverDefense *string      `db:"update_combat_maneuver_defense" json:"updateCombatManeuverDefense"`
	StartBonusFeat              *string      `db:"start_bonus_feat" json:"startBonusFeat"`
	UpdateBonusFeat             *string      `db:"update_bonus_feat" json:"updateBonusFeat"`
	UpdateLevel                 int          `db:"update_level" json:"updateLevel"`
	MasteryLevel                int          `db:"mastery_level" json:"masteryLevel"`
	MasteryDescription          string       `db:"mastery_description" json:"masteryDescription"`
	BeastJson                   pgtype.JSONB `db:"beast" json:"-"`
	Beast                       *NameAlias   `db:"-" json:"beast"`
	BookJson                    pgtype.JSONB `db:"book" json:"-"`
	Book                        Books        `db:"-" json:"book"`
}

type AbilityInfo struct {
	Name        string  `db:"name" json:"name"`
	EngName     *string `db:"eng_name" json:"engName"`
	Alias       string  `db:"alias" json:"alias"`
	Description string  `db:"description" json:"description"`
	ShortName   string  `db:"short_name" json:"shortName"`
}

type Affliction struct {
	Alias              string       `db:"alias" json:"alias"`
	Name               string       `db:"name" json:"name"`
	EngName            *string      `db:"eng_name" json:"engName"`
	TypeDescription    *string      `db:"type_description" json:"typeDescription"`
	Save               *string      `db:"save" json:"save"`
	Onset              *string      `db:"onset" json:"onset"`
	Frequency          *string      `db:"frequency" json:"frequency"`
	Effect             *string      `db:"effect" json:"effect"`
	InitialEffect      *string      `db:"initial_effect" json:"initialEffect"`
	SecondaryEffect    *string      `db:"secondary_effect" json:"secondaryEffect"`
	Cure               *string      `db:"cure" json:"cure"`
	Cost               *float64     `db:"cost" json:"cost"`
	Description        *string      `db:"description" json:"description"`
	MainTypeJson       pgtype.JSONB `db:"main_type" json:"-"`
	MainType           NameAlias    `db:"-" json:"mainType"`
	SecondaryTypesJson pgtype.JSONB `db:"secondary_types" json:"-"`
	SecondaryTypes     []NameAlias  `db:"-" json:"secondaryTypes"`
	BookJson           pgtype.JSONB `db:"book" json:"-"`
	Book               Books        `db:"-" json:"book"`
}

type Trap struct {
	Alias           string       `db:"alias" json:"alias"`
	Name            string       `db:"name" json:"name"`
	EngName         *string      `db:"eng_name" json:"engName"`
	Cr              float64      `db:"cr" json:"cr"`
	Type            *string      `db:"type" json:"type"`
	DcPerception    int          `db:"dc_perception" json:"dcPerception"`
	DcDisableDevice int          `db:"dc_disable_device" json:"dcDisableDevice"`
	Trigger         string       `db:"trigger" json:"trigger"`
	Duration        *string      `db:"duration" json:"duration"`
	Reset           string       `db:"reset" json:"reset"`
	Effect          string       `db:"effect" json:"effect"`
	BookJson        pgtype.JSONB `db:"book" json:"-"`
	Book            Books        `db:"-" json:"book"`
}

type Haunt struct {
	Alias         string       `db:"alias" json:"alias"`
	Name          string       `db:"name" json:"name"`
	EngName       *string      `db:"eng_name" json:"engName"`
	Cr            float64      `db:"cr" json:"cr"`
	CrStr         string       `db:"cr_str" json:"crStr"`
	Exp           int          `db:"exp" json:"exp"`
	AlignmentArea string       `db:"alignment_area" json:"alignmentArea"`
	CasterLevel   int          `db:"caster_level" json:"casterLevel"`
	Hp            int          `db:"hp" json:"hp"`
	Notice        string       `db:"notice" json:"notice"`
	Weakness      *string      `db:"weakness" json:"weakness"`
	Trigger       string       `db:"trigger" json:"trigger"`
	Reset         string       `db:"reset" json:"reset"`
	Destruction   *string      `db:"destruction" json:"destruction"`
	BookJson      pgtype.JSONB `db:"book" json:"-"`
	Book          Books        `db:"-" json:"book"`
}

type Good struct {
	Id            int    `db:"id" json:"id"`
	Cnt           int    `db:"cnt" json:"cnt"`
	Priority      int    `db:"priority" json:"priority"`
	Name          string `db:"name" json:"name"`
	Url           string `db:"url" json:"url"`
	ImageUrls     string `db:"image_urls" json:"imageUrls"`
	Price         int    `db:"price" json:"price"`
	InWaitingList bool   `db:"in_waiting_list" json:"inWaitingList"`
}

type UserData struct {
	Id    int     `db:"id" json:"id"`
	Login string  `db:"login" json:"login"`
	Email *string `db:"email" json:"email"`
	Token *string `db:"token" json:"token"`
	Role  string  `db:"role" json:"role"`
}

type Favourites struct {
	Name           string       `db:"name" json:"name"`
	Guid           string       `db:"guid" json:"guid"`
	FavouritesJson pgtype.JSONB `db:"favourites" json:"-"`
	Favourites     []Favourite  `db:"-" json:"favourites"`
}

type Favourite struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PushToken struct {
	Id    int    `db:"id" json:"id"`
	Token string `db:"token" json:"token"`
}

type SearchInfo struct {
	Id       string    `db:"id" json:"id"`
	Type     string    `db:"type" json:"type"`
	Url      string    `db:"url" json:"url"`
	H1       string    `db:"h1" json:"h1"`
	Content  *string   `db:"content" json:"content"`
	DtUpdate time.Time `db:"-" json:"dtUpdate"`
	Score    string    `db:"-" json:"score"`
}

type SearchResult struct {
	Url   string  `db:"url" json:"url"`
	H1    string  `db:"h1" json:"h1"`
	Type  string  `db:"type" json:"type"`
	Score float64 `db:"-" json:"score"`
}

type ElasticResult struct {
	Hits ElasticResultHits `json:"hits"`
}

type ElasticResultHits struct {
	Hits []ElasticResultHitsHits `json:"hits"`
}

type ElasticResultHitsHits struct {
	Score  float64    `json:"_score"`
	Source SearchInfo `json:"_source"`
}

type WildTalent struct {
	Alias            string         `db:"alias" json:"alias"`
	Name             string         `db:"name" json:"name"`
	EngName          string         `db:"eng_name" json:"engName"`
	Description      *string        `db:"description" json:"description"`
	Element          string         `db:"element" json:"element"`
	TypeDescription  *string        `db:"type_description" json:"typeDescription"`
	Level            *int           `db:"level" json:"level"`
	Burn             *int           `db:"burn" json:"burn"`
	BurnDescription  *string        `db:"burn_description" json:"burnDescription"`
	Prerequisites    *string        `db:"prerequisites" json:"prerequisites"`
	BlastType        *string        `db:"blast_type" json:"blastType"`
	Damage           *string        `db:"damage" json:"damage"`
	SavingThrow      *string        `db:"saving_throw" json:"savingThrow"`
	SpellResistance  *string        `db:"spell_resistance" json:"spellResistance"`
	AssociatedBlasts *string        `db:"associated_blasts" json:"associatedBlasts"`
	TypeJson         pgtype.JSONB   `db:"type" json:"-"`
	Type             WildTalentType `db:"-" json:"type"`
	BookJson         pgtype.JSONB   `db:"book" json:"-"`
	Book             Books          `db:"-" json:"book"`
}

type WildTalentType struct {
	Alias       string  `db:"alias" json:"alias"`
	Name        string  `db:"name" json:"name"`
	ShortName   string  `db:"short_name" json:"shortName"`
	Description *string `db:"description" json:"description"`
	Order       int     `db:"order" json:"order"`
}

type ShamanSpiritInfo struct {
	Alias                    string               `db:"alias" json:"alias"`
	Name                     string               `db:"name" json:"name"`
	EngName                  string               `db:"eng_name" json:"engName"`
	Description              *string              `db:"description" json:"description"`
	HexesJson                pgtype.JSONB         `db:"hexes" json:"-"`
	Hexes                    []AbilityDescription `db:"-" json:"hexes"`
	SpiritAnimal             string               `db:"spirit_animal" json:"spiritAnimal"`
	SpiritAbilityJson        pgtype.JSONB         `db:"spirit_ability" json:"-"`
	SpiritAbility            AbilityDescription   `db:"-" json:"spiritAbility"`
	GreaterSpiritAbilityJson pgtype.JSONB         `db:"greater_spirit_ability" json:"-"`
	GreaterSpiritAbility     AbilityDescription   `db:"-" json:"greaterSpiritAbility"`
	TrueSpiritAbilityJson    pgtype.JSONB         `db:"true_spirit_ability" json:"-"`
	TrueSpiritAbility        AbilityDescription   `db:"-" json:"trueSpiritAbility"`
	Manifestation            string               `db:"manifestation" json:"manifestation"`
	ParentJson               pgtype.JSONB         `db:"parent" json:"-"`
	Parent                   *NameAlias           `db:"-" json:"parent"`
	SpellsJson               pgtype.JSONB         `db:"spells" json:"-"`
	Spells                   []SpiritSpell        `db:"-" json:"spells"`
	BookJson                 pgtype.JSONB         `db:"book" json:"-"`
	Book                     Books                `db:"-" json:"book"`
}

type SpiritSpell struct {
	Alias   string  `json:"alias"`
	Name    string  `json:"name"`
	Level   int     `json:"level"`
	Comment *string `json:"comment"`
}

type AbilityDescription struct {
	Name        string  `json:"name"`
	EngName     string  `json:"engName"`
	Type        *string `json:"type"`
	Description string  `json:"description"`
}

type StaticPage struct {
	Id          *int         `db:"id"       json:"id"`
	Name        string       `db:"name"     json:"name"`
	Slug        string       `db:"slug"     json:"slug"`
	ContentJson pgtype.JSONB `db:"content"  json:"-"`
	Content     *string      `db:"-"        json:"content"`
}

type MediumSpiritInfo struct {
	Alias                       string             `db:"alias" json:"alias"`
	Name                        string             `db:"name" json:"name"`
	EngName                     string             `db:"eng_name" json:"engName"`
	Description                 *string            `db:"description" json:"description"`
	Type                        *string            `db:"type" json:"type"`
	SpiritBonus                 *string            `db:"spirit_bonus" json:"spiritBonus"`
	SeanceBoon                  *string            `db:"seance_boon" json:"seanceBoon"`
	FavoredLocations            *string            `db:"favored_locations" json:"favoredLocations"`
	InfluencePenalty            *string            `db:"influence_penalty" json:"influencePenalty"`
	Taboos                      *string            `db:"taboos" json:"taboos"`
	SpiritPowerBaseJson         pgtype.JSONB       `db:"spirit_power_base" json:"-"`
	SpiritPowerBase             AbilityDescription `db:"-" json:"spiritPowerBase"`
	SpiritPowerIntermediateJson pgtype.JSONB       `db:"spirit_power_intermediate" json:"-"`
	SpiritPowerIntermediate     AbilityDescription `db:"-" json:"spiritPowerIntermediate"`
	SpiritPowerGreaterJson      pgtype.JSONB       `db:"spirit_power_greater" json:"-"`
	SpiritPowerGreater          AbilityDescription `db:"-" json:"spiritPowerGreater"`
	SpiritPowerSupremeJson      pgtype.JSONB       `db:"spirit_power_supreme" json:"-"`
	SpiritPowerSupreme          AbilityDescription `db:"-" json:"spiritPowerSupreme"`
	BookJson                    pgtype.JSONB       `db:"book" json:"-"`
	Book                        Books              `db:"-" json:"book"`
}

type Aspect struct {
	Alias                 string       `db:"alias" json:"alias"`
	Name                  string       `db:"name" json:"name"`
	EngName               string       `db:"eng_name" json:"engName"`
	Description           string       `db:"description" json:"description"`
	MinorForm             string       `db:"minor_form" json:"minorForm"`
	MajorForm             string       `db:"major_form" json:"majorForm"`
	AdditionalDescription *string      `db:"additional_description" json:"additionalDescription"`
	BookJson              pgtype.JSONB `db:"book" json:"-"`
	Book                  Books        `db:"-" json:"book"`
}
