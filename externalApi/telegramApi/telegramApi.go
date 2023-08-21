package telegramApi

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

type TelegramApi struct {
	client *resty.Client
	cfg    *config.Config
}

func NewTelegramApi(cfg *config.Config) *TelegramApi {
	return &TelegramApi{
		client: resty.New().
			SetDebug(cfg.API.Debug).
			SetTimeout(time.Duration(cfg.API.Timeout) * time.Millisecond),
		cfg: cfg,
	}
}

func (a *TelegramApi) Send(ctx context.Context, request model.SendTelegramRequest) (*model.TelegramApiResponse, error) {
	/*	span := jaeger.GetSpan(ctx, "Send")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	url := fmt.Sprintf(a.cfg.API.TelegramAPI.Send, a.cfg.API.TelegramAPI.BotToken)
	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"url": url, "request": request}, "Send")

	requestBody := map[string]interface{}{
		"text":                     request.Body,
		"parse_mode":               "HTML",
		"disable_web_page_preview": true,
	}
	if request.IsTest != nil && *request.IsTest {
		requestBody["chat_id"] = a.cfg.API.TelegramAPI.AdminChatId
	} else {
		requestBody["chat_id"] = a.cfg.API.TelegramAPI.ChannelId
	}

	body, err := a.client.R().
		SetContext(ctx).
		SetHeader("Content-Type", "application/json").
		SetBody(requestBody).
		Post(url)

	if err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"url": url, "request": request}, "Network error while dialing TelegramApi", err)
		return nil, errs.NetworkError
	}
	if body.StatusCode() != 200 {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"url": url, "request": request, "body": string(body.Body()), "status": body.Status()}, "TelegramApi returned non OK status", err)
		return nil, errs.NonOkStatusCode
	}
	//logger.TracefWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"url": url, "request": request}, "TelegramApi responded in %v seconds", body.Time().Seconds())

	var response model.TelegramApiResponse
	if err = json.Unmarshal(body.Body(), &response); err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"url": url, "request": request, "body": string(body.Body())}, "Can't parse TelegramApi response", err)
		return nil, errs.InvalidResponse
	}

	return &response, nil
}
