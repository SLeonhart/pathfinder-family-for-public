package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"pathfinder-family/config"
	"pathfinder-family/data/cache/inmemory"
	"pathfinder-family/data/db/dbInterface"
	"pathfinder-family/data/db/postgres"
	"pathfinder-family/externalApi/fcmApi"
	"pathfinder-family/externalApi/telegramApi"
	"pathfinder-family/externalApi/vkApi"

	"pathfinder-family/infrastructure/logger"
	a "pathfinder-family/presentation/api/handler"
	s "pathfinder-family/presentation/site/handler"
	"pathfinder-family/presentation/swagger"
	"pathfinder-family/scheduller"
	"pathfinder-family/server"
	"pathfinder-family/service/elasticSearchService"
	"pathfinder-family/service/emailService"
)

//go:generate go get github.com/swaggo/swag/gen
//go:generate go get github.com/swaggo/swag/cmd/swag
//go:generate go run github.com/swaggo/swag/cmd/swag init -g main.go -o ./presentation/swagger

// @title Pathfinder Family
// @version 20220823
// @description Сайт с информацией для настольной игры Pathfinder
// @license.name Чудинов Андрей Дмитриевич

func main() {
	cfg := config.Get()
	swagger.SwaggerInfo.Version = cfg.App.Version

	logger.Init(cfg.Logger.Level)

	inmemory := inmemory.NewInMemory(cfg)

	postgres := postgres.NewPostgres(cfg, inmemory)
	//singleton.NewSimpleSingleton(cfg, postgres)

	fcmApi := fcmApi.NewFcmApi(cfg)
	telegramApi := telegramApi.NewTelegramApi(cfg)
	vkApi := vkApi.NewVkApi(cfg)

	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		destroy(cancel, postgres)
	}()

	elasticSearchService := elasticSearchService.NewElasticSearchService(cfg, postgres)
	emailService := emailService.NewEmailService(cfg)
	go elasticSearchService.UpdatePathfinderSearch(ctx)

	done := make(chan os.Signal)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sched := scheduller.NewScheduller(cfg)
		sched.Run()

		apiHandler := a.NewHandler(cfg, postgres, fcmApi, telegramApi, vkApi, elasticSearchService, emailService)
		siteHandler := s.NewHandler(cfg, postgres, inmemory)

		webServer := server.NewWebServer(
			cfg,
			inmemory,
			postgres,
			server.WithApiHandler(apiHandler),
			server.WithSiteHandler(siteHandler),
		)

		if err := webServer.Run(ctx); err != nil {
			logger.Errorf(logger.CreateRequestIDField(ctx), "server listen error: %s", err.Error())
		}
	}()

	logger.Infof(logger.CreateRequestIDField(ctx), "%s server started", cfg.App.ServiceName)

	<-done

	logger.Infof(logger.CreateRequestIDField(ctx), "%s server stopped", cfg.App.ServiceName)
}

// destroy функционал, который нужно вызвать перед завершением программы
func destroy(cancel context.CancelFunc, postgres dbInterface.IPostgres) {
	postgres.Close()

	cancel()
}
