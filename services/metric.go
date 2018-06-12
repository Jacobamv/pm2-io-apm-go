package services

import "github.com/keymetrics/pm2-io-apm-go/structures"

var metrics []*structures.Metric

func AddMetric(metric *structures.Metric) {
	metrics = append(metrics, metric)
}

func GetMetricsAsMap() map[string]*structures.Metric {
	m := make(map[string]*structures.Metric, len(metrics))
	for _, metric := range metrics {
		metric.Get()
		m[metric.Name] = metric
	}
	return m
}
