package middleware

import (
	"context"
	"time"

	log "github.com/go-kit/kit/log"
	"github.com/nongdenchet/covidform/endpoint"
)

type LoggingMiddleware struct {
	Logger log.Logger
	Next   endpoint.IdenticontService
}

func (mw LoggingMiddleware) Generate(ctx context.Context, text string, size int) (output string, err error) {
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "generate",
			"size", size,
			"text", text,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.Next.Generate(ctx, text, size)
	return
}
