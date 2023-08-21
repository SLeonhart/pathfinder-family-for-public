package emailService

import (
	"context"
	"crypto/tls"

	"pathfinder-family/config"

	"pathfinder-family/infrastructure/logger"

	"gopkg.in/gomail.v2"
)

type EmailService struct {
	cfg *config.Email
}

func NewEmailService(cfg *config.Config) *EmailService {
	return &EmailService{
		cfg: &cfg.Email,
	}
}

func (s *EmailService) SendEmail(ctx context.Context, email string, body string) error {
	/*	span := jaeger.GetSpan(ctx, "SendEmail")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	m := gomail.NewMessage()
	m.SetHeader("From", s.cfg.User)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Сброс пароля для сайта pathfinder.family")
	m.SetBody("text/html", body)

	d := gomail.NewDialer(s.cfg.Host, s.cfg.Port, s.cfg.User, s.cfg.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"body": body}, "Send email error", err)
		return err
	}
	return nil
}
