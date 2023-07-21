package gateway

import (
	"github.com/aliworkshop/configlib"
	"github.com/aliworkshop/loggerlib/logger"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type RouterGroupModel interface {
	// SetMonitoringHandler set monitoring handler before handling routes
	// to monitor requests
	SetMonitoringHandler(monitoring MonitoringModel)

	Group(relativePath string) RouterGroupModel
	READ(path string, handlers ...HandlerEngine)
	CREATE(path string, handlers ...HandlerEngine)
	UPDATE(path string, handlers ...HandlerEngine)
	DELETE(path string, handlers ...HandlerEngine)
	STATIC(path string)
	// middleware
	Middleware(handlers ...HandlerEngine)
	SetupMiddlewares(registry configlib.Registry,
		logger logger.Logger, languageBundle *i18n.Bundle)
}

func RegisterRouters(model RouterGroupModel, path string, action ActionType, handlers ...HandlerEngine) {
	switch action {
	case ActionRead:
		model.READ(path, handlers...)
	case ActionCreate:
		model.CREATE(path, handlers...)
	case ActionUpdate:
		model.UPDATE(path, handlers...)
	case ActionDelete:
		model.DELETE(path, handlers...)
	}
}
