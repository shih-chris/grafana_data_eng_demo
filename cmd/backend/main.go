package main

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/shih-chris/grafana_data_eng_demo/pkg/query"
)

// Create Gauge
var (
	dsTypesInstalled = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "datasource_types_installed",
			Help: "Count of data sources installed aggregated by data source type",
		},
		[]string{
			"dsType",
		},
	)
)

func recordMetrics() {
	go func() {
		for {
			// Get MySQL Results
			dsTypeCounts := query.GetData()

			// Set gauge values
			for _, v := range dsTypeCounts {
				dsTypesInstalled.WithLabelValues(v.DsType).Set(v.DsCount)
			}

			// Pause for 15 seconds
			time.Sleep(15 * time.Second)

			// TODO: Figure out how to set counts to 0
			// 		 when they no longer show up in results

		}
	}()
}

func main() {
	recordMetrics()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
