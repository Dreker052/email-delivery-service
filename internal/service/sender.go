package service

import (
	"github.com/Dreker052/email-delivery-service.git/internal/config"
	"gopkg.in/gomail.v2"
)

type Sender struct {
	dialer *gomail.Dialer
}

func NewSender(cfg *config.Config) *Sender {
	return &Sender{
		dialer: gomail.NewDialer(cfg.SMTPHost, cfg.SMTPPort, cfg.SMTPUser, cfg.SMTPPass),
	}
}

func (s *Sender) Send(to, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "no-reply@productivity.app")
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	return s.dialer.DialAndSend(m)
}
