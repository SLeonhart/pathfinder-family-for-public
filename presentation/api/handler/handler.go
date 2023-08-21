package handler

import (
	"pathfinder-family/config"
	"pathfinder-family/data/db/dbInterface"
	"pathfinder-family/externalApi/apiInterface"
	"pathfinder-family/service/serviceInterface"
)

type Handler struct {
	cfg                  *config.Config
	postgres             dbInterface.IPostgres
	fcmApi               apiInterface.IFcmApi
	telegramApi          apiInterface.ITelegramApi
	vkApi                apiInterface.IVkApi
	elasticSearchService serviceInterface.IElasticSearchService
	emailService         serviceInterface.IEmailService
}

func NewHandler(cfg *config.Config, postgres dbInterface.IPostgres, fcmApi apiInterface.IFcmApi, telegramApi apiInterface.ITelegramApi, vkApi apiInterface.IVkApi, elasticSearchService serviceInterface.IElasticSearchService, emailService serviceInterface.IEmailService) *Handler {
	return &Handler{
		cfg:                  cfg,
		postgres:             postgres,
		fcmApi:               fcmApi,
		telegramApi:          telegramApi,
		vkApi:                vkApi,
		elasticSearchService: elasticSearchService,
		emailService:         emailService,
	}
}
