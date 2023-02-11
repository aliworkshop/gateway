package handlerlib

import (
	"github.com/aliworkshop/configlib"
	"github.com/aliworkshop/loggerlib/logger"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type RouterGroupModel interface {
	// SetMonitoringHandler set monitoring handler before handling routes
	// to monitor requests
	SetMonitoringHandler(monitoring MonitoringModel)
	//
	Group(relativePath string) RouterGroupModel
	READ(path string, handlers ...HandlerModel)
	CREATE(path string, handlers ...HandlerModel)
	UPDATE(path string, handlers ...HandlerModel)
	DELETE(path string, handlers ...HandlerModel)
	// middleware
	Middleware(handlers ...HandlerModel)
	SetupMiddlewares(registry configlib.Registry,
		logger logger.Logger, languageBundle *i18n.Bundle)
}

func RegisterRouters(model RouterGroupModel, path string, action ActionType, handlers ...HandlerModel) {
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
