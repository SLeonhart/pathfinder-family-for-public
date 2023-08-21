package vkApi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"pathfinder-family/config"

	"pathfinder-family/infrastructure/logger"
	"pathfinder-family/model"
	"pathfinder-family/model/errs"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
)

type VkApi struct {
	client *resty.Client
	cfg    *config.Config
}

func NewVkApi(cfg *config.Config) *VkApi {
	return &VkApi{
		client: resty.New().
			SetDebug(cfg.API.Debug).
			SetTimeout(time.Duration(cfg.API.Timeout) * time.Millisecond),
		cfg: cfg,
	}
}

func (a *VkApi) Send(ctx context.Context, request model.SendVkRequest, attachments []string) (*model.VkApiResponse, error) {
	/*	span := jaeger.GetSpan(ctx, "Send")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	url := a.cfg.API.VkAPI.Send
	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"url": url, "request": request}, "Send")

	formData := map[string]string{
		"v":            "5.131",
		"access_token": a.cfg.API.VkAPI.AccessToken,
		"from_group":   "1",
		"copyright":    "https://pathfinder.family/",
		"message":      request.Body,
	}
	if request.IsTest != nil && *request.IsTest {
		formData["owner_id"] = fmt.Sprintf("-%v", a.cfg.API.VkAPI.TestGroupId)
	} else {
		formData["owner_id"] = fmt.Sprintf("-%v", a.cfg.API.VkAPI.GroupId)
	}
	if len(attachments) > 0 {
		formData["attachments"] = strings.Join(attachments, ",")
	}

	body, err := a.client.R().
		SetContext(ctx).
		SetHeader("Content-Type", "application/json").
		SetFormData(formData).
		Post(url)

	if err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"url": url, "request": request}, "Network error while dialing VkApi", err)
		return nil, errs.NetworkError
	}
	if body.StatusCode() != 200 {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"url": url, "request": request, "body": string(body.Body()), "status": body.Status()}, "VkApi returned non OK status", err)
		return nil, errs.NonOkStatusCode
	}
	//logger.TracefWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"url": url, "request": request}, "VkApi responded in %v seconds", body.Time().Seconds())

	var response model.VkApiResponse
	if err = json.Unmarshal(body.Body(), &response); err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"url": url, "request": request, "body": string(body.Body())}, "Can't parse VkApi response", err)
		return nil, errs.InvalidResponse
	}

	return &response, nil
}

func (a *VkApi) GetPhotoServer(ctx context.Context, request model.SendVkRequest) (*model.VkApiGetPhotoServerResponse, error) {
	/*	span := jaeger.GetSpan(ctx, "GetPhotoServer")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	url := a.cfg.API.VkAPI.GetPhotoServer
	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"url": url, "request": request}, "GetPhotoServer")

	formData := map[string]string{
		"v":            "5.131",
		"access_token": a.cfg.API.VkAPI.AccessToken,
	}
	if request.IsTest != nil && *request.IsTest {
		formData["group_id"] = fmt.Sprintf("%v", a.cfg.API.VkAPI.TestGroupId)
	} else {
		formData["group_id"] = fmt.Sprintf("%v", a.cfg.API.VkAPI.GroupId)
	}

	body, err := a.client.R().
		SetContext(ctx).
		SetHeader("Content-Type", "application/json").
		SetFormData(formData).
		Post(url)

	if err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"url": url, "request": request}, "Network error while dialing VkApi", err)
		return nil, errs.NetworkError
	}
	if body.StatusCode() != 200 {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"url": url, "request": request, "body": string(body.Body()), "status": body.Status()}, "VkApi returned non OK status", err)
		return nil, errs.NonOkStatusCode
	}
	//logger.TracefWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"url": url, "request": request}, "VkApi responded in %v seconds", body.Time().Seconds())

	var response model.VkApiGetPhotoServerResponse
	if err = json.Unmarshal(body.Body(), &response); err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"url": url, "request": request, "body": string(body.Body())}, "Can't parse VkApi response", err)
		return nil, errs.InvalidResponse
	}

	return &response, nil
}

func (a *VkApi) GetPhoto(ctx context.Context, url string) ([]byte, error) {
	/*	span := jaeger.GetSpan(ctx, "GetPhoto")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"url": url}, "GetPhoto")

	body, err := a.client.R().
		SetContext(ctx).
		Get(url)

	if err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"url": url}, "Network error while getting photo", err)
		return nil, errs.NetworkError
	}
	if body.StatusCode() != 200 {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"url": url, "body": string(body.Body()), "status": body.Status()}, "Getting photo returned non OK status", err)
		return nil, errs.NonOkStatusCode
	}
	//logger.TracefWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"url": url}, "Getting photo responded in %v seconds", body.Time().Seconds())

	return body.Body(), nil
}

func (a *VkApi) LoadPhotoIntoServer(ctx context.Context, url string, photoName string, photo []byte) (*model.VkApiLoadPhotoIntoServerResponse, error) {
	/*	span := jaeger.GetSpan(ctx, "LoadPhotoIntoServer")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"url": url}, "LoadPhotoIntoServer")

	reader := bytes.NewReader(photo)

	body, err := a.client.R().
		SetContext(ctx).
		SetFileReader("photo", photoName, reader).
		Post(url)

	if err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"url": url}, "Network error while dialing VkApi", err)
		return nil, errs.NetworkError
	}
	if body.StatusCode() != 200 {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"url": url, "body": string(body.Body()), "status": body.Status()}, "VkApi returned non OK status", err)
		return nil, errs.NonOkStatusCode
	}
	//logger.TracefWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"url": url}, "VkApi responded in %v seconds", body.Time().Seconds())

	var response model.VkApiLoadPhotoIntoServerResponse
	if err = json.Unmarshal(body.Body(), &response); err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"url": url, "body": string(body.Body())}, "Can't parse VkApi response", err)
		return nil, errs.InvalidResponse
	}

	return &response, nil
}

func (a *VkApi) SavePhoto(ctx context.Context, loadPhotoRes model.VkApiLoadPhotoIntoServerResponse, isTest *bool) (*model.VkApiSavePhotoResponse, error) {
	/*	span := jaeger.GetSpan(ctx, "SavePhoto")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	url := a.cfg.API.VkAPI.SavePhoto
	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"url": url, "loadPhotoRes": loadPhotoRes, "isTest": isTest}, "SavePhoto")

	formData := map[string]string{
		"v":            "5.131",
		"access_token": a.cfg.API.VkAPI.AccessToken,
		"photo":        *loadPhotoRes.Photo,
		"server":       fmt.Sprintf("%v", *loadPhotoRes.Server),
		"hash":         *loadPhotoRes.Hash,
	}
	if isTest != nil && *isTest {
		formData["group_id"] = fmt.Sprintf("%v", a.cfg.API.VkAPI.TestGroupId)
	} else {
		formData["group_id"] = fmt.Sprintf("%v", a.cfg.API.VkAPI.GroupId)
	}

	body, err := a.client.R().
		SetContext(ctx).
		SetHeader("Content-Type", "application/json").
		SetFormData(formData).
		Post(url)

	if err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"url": url, "loadPhotoRes": loadPhotoRes, "isTest": isTest}, "Network error while dialing VkApi", err)
		return nil, errs.NetworkError
	}
	if body.StatusCode() != 200 {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"url": url, "loadPhotoRes": loadPhotoRes, "isTest": isTest, "body": string(body.Body()), "status": body.Status()}, "VkApi returned non OK status", err)
		return nil, errs.NonOkStatusCode
	}
	//logger.TracefWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"url": url, "loadPhotoRes": loadPhotoRes, "isTest": isTest}, "VkApi responded in %v seconds", body.Time().Seconds())

	var response model.VkApiSavePhotoResponse
	if err = json.Unmarshal(body.Body(), &response); err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"url": url, "loadPhotoRes": loadPhotoRes, "isTest": isTest, "body": string(body.Body())}, "Can't parse VkApi response", err)
		return nil, errs.InvalidResponse
	}

	return &response, nil
}
