package gateway

import (
	errors "github.com/aliworkshop/errors"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

type ServerModel interface {
	AddMonitoring(m *Monitoring) (prometheus.Collector, errors.ErrorModel)
	StartMonitoring()
	Middleware(handlers ...Handler)
	SetController(controller Controller)
	NewRouterGroup(path string) RouterGroupModel
	LoadHtml(path string)
	Shutdown(timeout time.Duration) error
	Run(...string) error
}
