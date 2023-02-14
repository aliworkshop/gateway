package handlerlib

import (
	"github.com/aliworkshop/errorslib"
	"github.com/aliworkshop/loggerlib/logger"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

type ServerModel interface {
	SetMonitoringHandler(monitoring MonitoringModel)
	AddMonitoring(m *Monitoring) (prometheus.Collector, errorslib.ErrorModel)
	StartMonitoring()
	Middleware(handlers ...HandlerModel)
	SetupMiddlewares(logger logger.Logger, languageBundle *i18n.Bundle)
	NewRouterGroup(path string) RouterGroupModel
	Shutdown(timeout time.Duration) error
	Run(...string) error
}
