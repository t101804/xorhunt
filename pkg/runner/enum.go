package runner

import (
	"context"
	"io"
	"time"

	"github.com/t101804/xorhunt/pkg/logger"
	"github.com/t101804/xorhunt/pkg/resolve"
)

func (r *Runner) EnumurateSingleWithCtx(ctx context.Context, single string, writers []io.Writer) error {
	logger.Info().Msgf("Starting for %s", single)
	var resolutionPool *resolve.ResolutionPool
	now := time.Now()
	passiveResults := r.passiveAgent.EnumerateSubdomainsWithCtx(ctx, domain, r.options.Proxy, r.options.RateLimit, r.options.Timeout, time.Duration(r.options.MaxEnumerationTime)*time.Minute, passive.WithCustomRateLimit(r.rateLimit))

	return nil
}
