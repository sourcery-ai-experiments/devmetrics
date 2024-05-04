package agent

import (
	"math/rand"
	"reflect"
	"runtime"
	"strconv"
	"time"

	"github.com/rybalka1/devmetrics/internal/memstorage"
	"github.com/rybalka1/devmetrics/internal/metrics"
)

func (agent Agent) GetMetrics() {
	if agent.metrics == nil {
		agent.metrics = make(map[string]metrics.MyMetrics)
	}
	m := runtime.MemStats{}
	runtime.ReadMemStats(&m)
	values := reflect.ValueOf(m)
	for _, name := range usedMemStats {
		if values.FieldByName(name).IsValid() {
			if values.FieldByName(name).CanInt() {
				agent.metrics[name] = metrics.MyMetrics{
					Value:    strconv.FormatInt(values.FieldByName(name).Int(), 10),
					SendType: metrics.Gauge,
				}
				continue
			}
			if values.FieldByName(name).CanUint() {
				agent.metrics[name] = metrics.MyMetrics{
					Value:    strconv.FormatUint(values.FieldByName(name).Uint(), 10),
					SendType: metrics.Gauge,
				}
				continue
			}
			if values.FieldByName(name).CanFloat() {
				agent.metrics[name] = metrics.MyMetrics{
					Value:    strconv.FormatFloat(values.FieldByName(name).Float(), 'f', -1, 64),
					SendType: metrics.Gauge,
				}
				continue
			}
		}
	}

	metric, ok := agent.metrics["PollCount"]
	if !ok {
		agent.metrics["PollCount"] = metrics.MyMetrics{
			Value:    "1",
			SendType: metrics.Counter,
		}
	} else {
		metric.AddVal(1)
		agent.metrics["PollCount"] = metric
	}

	r := rand.New(rand.NewSource(time.Now().Unix()))
	randVal := float64(r.Intn(1000)) + r.Float64()
	agent.metrics["RandomValue"] = metrics.MyMetrics{
		Value:    strconv.FormatFloat(randVal, 'f', -1, 64),
		SendType: metrics.Gauge,
	}
}

func (agent Agent) CollectMetrics() {
	agent.store = memstorage.NewMemStorage()
	m := runtime.MemStats{}
	runtime.ReadMemStats(&m)
	values := reflect.ValueOf(m)
	for _, name := range usedMemStats {
		if values.FieldByName(name).IsValid() {
			if values.FieldByName(name).CanInt() {
				val := values.FieldByName(name).Int()
				agent.store.UpdateMetric(&metrics.Metrics{
					ID:    name,
					MType: metrics.Counter,
					Delta: &val,
					Value: nil,
				})
			}
			if values.FieldByName(name).CanFloat() {
				val := values.FieldByName(name).Float()
				agent.store.UpdateMetric(&metrics.Metrics{
					ID:    name,
					MType: metrics.Gauge,
					Delta: nil,
					Value: &val,
				})
			}
		}
	}

	metric, ok := agent.metrics["PollCount"]
	if !ok {
		agent.metrics["PollCount"] = metrics.MyMetrics{
			Value:    "1",
			SendType: metrics.Counter,
		}
	} else {
		metric.AddVal(1)
		agent.metrics["PollCount"] = metric
	}

	r := rand.New(rand.NewSource(time.Now().Unix()))
	randVal := float64(r.Intn(1000)) + r.Float64()
	agent.metrics["RandomValue"] = metrics.MyMetrics{
		Value:    strconv.FormatFloat(randVal, 'f', -1, 64),
		SendType: metrics.Gauge,
	}
}