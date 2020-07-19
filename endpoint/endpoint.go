package endpoint

import (
	"context"
)

type IdenticontService interface {
	Generate(context.Context, string, int) (string, error)
}
