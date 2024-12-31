package main

import (
	"github.com/Divas-Gupta30/medium-blogs/wire-go/service"
	"github.com/Divas-Gupta30/medium-blogs/wire-go/wire"
)

func main() {
	smtpClient := service.NewSMTPClient(&service.Config{
		Host: "localhost",
		Port: "8080",
	})
	emailSvc := service.NewEmailService(smtpClient)
	db := service.NewDb(&service.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "root",
	})

	notifSvc := service.NewNotificationService(db, emailSvc)
	notifSvc.SendNotifications()

	notifSvc2 := wire.InitialiseNotificationService(&service.DbConf{
		Host: "localhost",
		Port: "8080",
	}, &service.Config{
		Host: "localhost",
		Port: "8080",
	})
	notifSvc2.SendNotifications()
}
