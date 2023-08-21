package model

import "pathfinder-family/config"

type Page struct {
	Cfg              config.App
	Title            string
	Url              string
	Description      string
	IsNotFixedNavBar bool
}

type DeafultPageGenerator struct {
	Page
	JsPaths []string
}

type DeafultPage struct {
	Page
	H1Class *string
	Alias   string
	Alias2  *string
	JsPaths []string
}

type CommonPage struct {
	Page
}

type AliasPage struct {
	Page
	Alias string
}

type DomainPage struct {
	Page
	Alias string
	Type  string
}

type BloodlinePage struct {
	Page
	Alias      string
	ClassAlias string
}

type ProfilePage struct {
	Page
	Role  string
	Login string
	Email *string
}
