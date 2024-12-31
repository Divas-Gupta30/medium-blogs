//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package wire

import (
	"awesomeProject2/wireTest/service"
	"github.com/google/wire"
)

func InitialiseNotificationService(conf1 *service.DbConf, conf2 *service.Config) *service.NotificationService {
	wire.Build(service.NotificationServiceSet)
	return &service.NotificationService{}
}
