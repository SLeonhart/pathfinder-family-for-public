package model

import "time"

type FeedbackRequest struct {
	Theme   string  `json:"theme"`
	Email   *string `json:"email"`
	Message string  `json:"message"`
}

type AddGoodInWaitingListRequest struct {
	Id int `json:"id"`
}

type UserAuthRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserAuthResponse struct {
	Url string `json:"url"`
}

type UserRegisterRequest struct {
	Login    string  `json:"login"`
	Password string  `json:"password"`
	Email    *string `json:"email"`
}

type UserResetPasswordRequest struct {
	Login string `json:"login"`
	Email string `json:"email"`
}

type UserChangeDataRequest struct {
	Password *string `json:"password"`
	Email    *string `json:"email"`
}

type UserFavouritesRequest struct {
	Guid string `json:"guid"`
}

type ChangeUserFavouritesItemsRequest struct {
	ForAdd   bool    `json:"forAdd"`
	Guid     *string `json:"guid"`
	PageName string  `json:"pageName"`
	Url      string  `json:"url"`
}

type RenameUserFavouritesRequest struct {
	Guid string `json:"guid"`
	Name string `json:"name"`
}

type AddDonateRequest struct {
	Dt     time.Time `json:"dt"`
	Sum    float64   `json:"sum"`
	Helper string    `json:"helper"`
}

type SendPushRequest struct {
	Title string  `json:"title"`
	Body  string  `json:"body"`
	Image string  `json:"image"`
	Link  string  `json:"link"`
	Token *string `json:"token"`
}

type SendTelegramRequest struct {
	Body   string `json:"body"`
	IsTest *bool  `json:"isTest"`
}

type AddNewsRequest struct {
	PushTitle      *string `json:"pushTitle"`
	PushBody       *string `json:"pushBody"`
	PushImageLink  *string `json:"pushImageLink"`
	PushTargetLink *string `json:"pushTargetLink"`
	Body           string  `json:"body"`
}

type SendVkRequest struct {
	Body   string   `json:"body"`
	Images []string `json:"images"`
	IsTest *bool    `json:"isTest"`
}

type UpsertResponse struct {
	Id int
}

type ElasticRequest struct {
	Dt time.Time `json:"dt"`
}
