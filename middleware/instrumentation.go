package middleware

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kit/kit/metrics"
	"github.com/nongdenchet/covidform/endpoint"
)

type InstrumentingMiddleware struct {
	RequestCount   metrics.Counter
	RequestLatency metrics.Histogram
	Next           endpoint.IdenticontService
}

func (mw InstrumentingMiddleware) Generate(ctx context.Context, text string, size int) (output string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "encrypt", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.Generate(ctx, text, size)
	return
}
