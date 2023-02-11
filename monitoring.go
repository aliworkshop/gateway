package handlerlib

type MonitoringModel interface {
	OnRequestStart(request RequestModel)
	OnRequestEnd(request RequestModel)
}

type MonitoringHandler interface {
	Handler
	MonitoringModel
}

var DefaultMonitoring MonitoringModel = new(emptyMonitoringHandler)

type emptyMonitoringHandler struct {
}

func (h *emptyMonitoringHandler) OnRequestStart(request RequestModel) {
}

func (h *emptyMonitoringHandler) OnRequestEnd(request RequestModel) {
}
