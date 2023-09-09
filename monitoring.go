package gateway

type MetricType string

const (
	CounterVec   MetricType = "counter_vec"
	Counter      MetricType = "counter"
	GaugeVec     MetricType = "gauge_vec"
	Gauge        MetricType = "gauge"
	HistogramVec MetricType = "histogram_vec"
	Histogram    MetricType = "histogram"
	SummaryVec   MetricType = "summary_vec"
	Summary      MetricType = "summary"
)

func (m MetricType) String() string {
	return string(m)
}

type Monitoring struct {
	Subsystem   string
	ID          string
	Name        string
	Description string
	Type        MetricType
	Args        []string
	Buckets     []float64
}
