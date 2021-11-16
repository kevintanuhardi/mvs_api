package metrics

import (
	golangPrometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

/*
OrderDispatchCounter add counter for order dispatch

Example:

prometheus.OrderDispatchCounter.With(prometheus.AddLabel([]prometheus.CounterLabel{
	{Key: "target", Value: "wpd"},      // wpd or grosir pintar
	{Key: "status", Value: "success"}, // success or error
})).Inc()

*/
var OrderDispatchCounter = promauto.NewCounterVec(golangPrometheus.CounterOpts{
	Namespace: "merchant",
	Subsystem: "mbs",
	Name:      "order_dispatch_counter",
	Help:      "The total number of orders dispatched",
}, []string{"target", "status"})
