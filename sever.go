package handlerlib

import (
	"github.com/aliworkshop/loggerlib/logger"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"time"
)

type ServerModel interface {
	SetMonitoringHandler(monitoring MonitoringModel)
	Middleware(handlers ...HandlerModel)
	SetupMiddlewares(logger logger.Logger, languageBundle *i18n.Bundle)
	NewRouterGroup(path string) RouterGroupModel
	Shutdown(timeout time.Duration) error
	Run(...string) error
}
