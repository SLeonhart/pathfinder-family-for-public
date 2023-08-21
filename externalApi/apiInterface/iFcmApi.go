package apiInterface

import (
	"context"
	"pathfinder-family/model"
)

//go:generate mockgen -source=iFcmApi.go -destination=../apiMock/fcmApi.go -package=apiMock

type IFcmApi interface {
	CheckTokens(ctx context.Context, tokens []string) (*model.FcmApiResponse, error)
	SendPush(ctx context.Context, request model.SendPushRequest, token string) (*model.FcmApiResponse, error)
}
