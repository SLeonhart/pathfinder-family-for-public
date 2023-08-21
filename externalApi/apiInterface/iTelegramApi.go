package apiInterface

import (
	"context"
	"pathfinder-family/model"
)

//go:generate mockgen -source=iTelegramApi.go -destination=../apiMock/telegramApi.go -package=apiMock

type ITelegramApi interface {
	Send(ctx context.Context, request model.SendTelegramRequest) (*model.TelegramApiResponse, error)
}
