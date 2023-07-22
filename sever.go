package gateway

import (
	"github.com/aliworkshop/errorslib"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

type ServerModel interface {
	SetMonitoringHandler(monitoring MonitoringModel)
	AddMonitoring(m *Monitoring) (prometheus.Collector, errorslib.ErrorModel)
	StartMonitoring()
	Middleware(handlers ...Handler)
	SetupMiddlewares()
	SetController(controller Controller)
	NewRouterGroup(path string) RouterGroupModel
	Shutdown(timeout time.Duration) error
	Run(...string) error
}
