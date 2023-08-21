package serviceInterface

import (
	"context"
)

//go:generate mockgen -source=iEmailService.go -destination=../serviceMock/emailService.go -package=serviceMock

type IEmailService interface {
	SendEmail(ctx context.Context, email string, body string) error
}
