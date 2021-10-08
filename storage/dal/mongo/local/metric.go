package local

import (
	"sync"
	"time"

	"github.com/wxc/cmdb/common/metrics"

	"github.com/prometheus/client_golang/prometheus"
)

var mtc *mongoMetric
var once = sync.Once{}

func initMongoMetric() {
	once.Do(func() {
		mtc = new(mongoMetric)

		mtc.totalOperCount = prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: metrics.Namespace,
			Subsystem: "mongo",
			Name:      "total_operate_count",
			Help:      "the total operate count with mongodb",
		}, []string{"collection", "operation"})
		metrics.Register().MustRegister(mtc.totalOperCount)

		mtc.totalErrorCount = prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: metrics.Namespace,
			Subsystem: "mongo",
			Name:      "total_error_count",
			Help:      "the total operate error count with mongodb",
		}, []string{"collection", "operation"})
		metrics.Register().MustRegister(mtc.totalErrorCount)

		mtc.operDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: metrics.Namespace,
			Subsystem: "mongo",
			Name:      "operate_duration_seconds",
			Help:      "the cost second duration with one mongodb operation",
			Buckets:   []float64{0.02, 0.04, 0.06, 0.08, 0.1, 0.3, 0.5, 0.7, 1, 5, 10, 20, 30, 60},
		}, []string{"collection", "operation"})
		metrics.Register().MustRegister(mtc.operDuration)
	})
}

type oper string

const (
	findOper        oper = "find"
	insertOper      oper = "insert"
	updateOper      oper = "update"
	upsertOper      oper = "upsert"
	deleteOper      oper = "delete"
	distinctOper    oper = "distinct"
	aggregateOper   oper = "aggregate"
	countOper       oper = "count"
	columnOper      oper = "column"
	indexCreateOper oper = "create_index"
	indexDropOper   oper = "drop_index"
)

type mongoMetric struct {
	// record the total operation countOper with mongodb
	totalOperCount *prometheus.CounterVec
	// record the total error operation countOper with mongodb.
	totalErrorCount *prometheus.CounterVec
	// record the operate duration with mongodb
	operDuration *prometheus.HistogramVec
}

func (m *mongoMetric) collectOperCount(collection string, operation oper) {
	if m == nil {
		return
	}

	m.totalOperCount.With(prometheus.Labels{
		"collection": collection,
		"operation":  string(operation),
	}).Inc()
}

func (m *mongoMetric) collectErrorCount(collection string, operation oper) {
	if m == nil {
		return
	}

	m.totalErrorCount.With(prometheus.Labels{
		"collection": collection,
		"operation":  string(operation),
	}).Inc()
}

func (m *mongoMetric) collectOperDuration(collection string, operation oper, duration time.Duration) {
	if m == nil {
		return
	}

	m.operDuration.With(prometheus.Labels{
		"collection": collection,
		"operation":  string(operation),
	}).Observe(duration.Seconds())
}
