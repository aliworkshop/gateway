package gateway

import (
	errors "github.com/aliworkshop/error"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

type ServerModel interface {
	SetMonitoringHandler(monitoring MonitoringModel)
	AddMonitoring(m *Monitoring) (prometheus.Collector, errors.ErrorModel)
	StartMonitoring()
	Middleware(handlers ...Handler)
	SetupMiddlewares()
	SetController(controller Controller)
	NewRouterGroup(path string) RouterGroupModel
	Shutdown(timeout time.Duration) error
	Run(...string) error
}
