package fcmApi

import (
	"context"
	"encoding/json"
	"fmt"
	"pathfinder-family/config"

	"pathfinder-family/infrastructure/logger"
	"pathfinder-family/model"
	"pathfinder-family/model/errs"
	"time"

	"github.com/go-resty/resty/v2"
)

type FcmApi struct {
	client *resty.Client
	cfg    *config.Config
}

func NewFcmApi(cfg *config.Config) *FcmApi {
	return &FcmApi{
		client: resty.New().
			SetDebug(cfg.API.Debug).
			SetTimeout(time.Duration(cfg.API.Timeout) * time.Millisecond),
		cfg: cfg,
	}
}

func (a *FcmApi) CheckTokens(ctx context.Context, tokens []string) (*model.FcmApiResponse, error) {
	/*	span := jaeger.GetSpan(ctx, "CheckTokens")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	url := a.cfg.API.FcmAPI.Send
	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"url": url}, "CheckTokens")

	body, err := a.client.R().
		SetContext(ctx).
		SetHeader("Authorization", fmt.Sprintf("key=%v", a.cfg.API.FcmAPI.Token)).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"registration_ids": tokens,
		}).
		Post(url)

	if err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"url": url}, "Network error while dialing FcmApi", err)
		return nil, errs.NetworkError
	}
	if body.StatusCode() != 200 {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"url": url, "body": string(body.Body()), "status": body.Status()}, "FcmApi returned non OK status", err)
		return nil, errs.NonOkStatusCode
	}
	//logger.TracefWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"url": url}, "FcmApi responded in %v seconds", body.Time().Seconds())

	var response model.FcmApiResponse
	if err = json.Unmarshal(body.Body(), &response); err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"url": url, "body": string(body.Body())}, "Can't parse FcmApi response", err)
		return nil, errs.InvalidResponse
	}

	return &response, nil
}

func (a *FcmApi) SendPush(ctx context.Context, request model.SendPushRequest, token string) (*model.FcmApiResponse, error) {
	/*	span := jaeger.GetSpan(ctx, "SendPush")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	url := a.cfg.API.FcmAPI.Send
	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"url": url, "request": request, "token": token}, "SendPush")

	body, err := a.client.R().
		SetContext(ctx).
		SetHeader("Authorization", fmt.Sprintf("key=%v", a.cfg.API.FcmAPI.Token)).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"notification": map[string]interface{}{
				"title":        request.Title,
				"body":         request.Body,
				"icon":         "https://pathfinder.family/icon.png",
				"image":        request.Image,
				"click_action": request.Link,
			},
			"to": token,
		}).
		Post(url)

	if err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"url": url, "request": request, "token": token}, "Network error while dialing FcmApi", err)
		return nil, errs.NetworkError
	}
	if body.StatusCode() != 200 {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"url": url, "request": request, "token": token, "body": string(body.Body()), "status": body.Status()}, "FcmApi returned non OK status", err)
		return nil, errs.NonOkStatusCode
	}
	//logger.TracefWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"url": url, "request": request, "token": token}, "FcmApi responded in %v seconds", body.Time().Seconds())

	var response model.FcmApiResponse
	if err = json.Unmarshal(body.Body(), &response); err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"url": url, "request": request, "token": token, "body": string(body.Body())}, "Can't parse FcmApi response", err)
		return nil, errs.InvalidResponse
	}

	return &response, nil
}
