package middleware

import (
	"os"

	"github.com/nongdenchet/covidform/endpoint"

	kitlog "github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

func NewLoggingMiddleware(s endpoint.IdenticontService) LoggingMiddleware {
	logger := kitlog.NewLogfmtLogger(os.Stderr)

	return LoggingMiddleware{Logger: logger, Next: s}
}

func NewIntrumentationMiddleware(s endpoint.IdenticontService) InstrumentingMiddleware {
	fieldKeys := []string{"method", "error"}

	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "identicon",
		Subsystem: "my_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)

	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "identicon",
		Subsystem: "my_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)

	return InstrumentingMiddleware{RequestCount: requestCount, RequestLatency: requestLatency, Next: s}
}
