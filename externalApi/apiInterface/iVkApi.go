package apiInterface

import (
	"context"
	"pathfinder-family/model"
)

//go:generate mockgen -source=iVkApi.go -destination=../apiMock/vkApi.go -package=apiMock

type IVkApi interface {
	Send(ctx context.Context, request model.SendVkRequest, attachments []string) (*model.VkApiResponse, error)
	GetPhotoServer(ctx context.Context, request model.SendVkRequest) (*model.VkApiGetPhotoServerResponse, error)
	GetPhoto(ctx context.Context, url string) ([]byte, error)
	LoadPhotoIntoServer(ctx context.Context, url string, photoName string, photo []byte) (*model.VkApiLoadPhotoIntoServerResponse, error)
	SavePhoto(ctx context.Context, loadPhotoRes model.VkApiLoadPhotoIntoServerResponse, isTest *bool) (*model.VkApiSavePhotoResponse, error)
}
