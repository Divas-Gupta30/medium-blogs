package service

import (
	"fmt"
	"github.com/google/wire"
)

var NotificationServiceSet = wire.NewSet(
	NewSMTPClient,
	NewEmailService,
	NewNotificationService,
	NewDb,
)

type DbConf Config

type EmailService struct {
	SMTPClient *SMTPClient
}

func NewEmailService(smtpClient *SMTPClient) *EmailService {
	return &EmailService{SMTPClient: smtpClient}
}

type SMTPClient struct {
	Config *Config
}

func NewSMTPClient(conf *Config) *SMTPClient {
	return &SMTPClient{Config: conf}
}

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
}

type Db struct {
	Config *Config
}

func NewDb(conf *Config) *Db {
	return &Db{
		Config: conf,
	}
}

type NotificationService struct {
	Db           *Db
	EmailService *EmailService
}

func NewNotificationService(db *Db, emailSvc *EmailService) *NotificationService {
	return &NotificationService{
		Db:           db,
		EmailService: emailSvc,
	}
}

func (n *NotificationService) SendNotifications() {
	fmt.Println("Notifications send")
}
