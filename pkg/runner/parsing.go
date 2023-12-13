package runner

import (
	"context"

	mapsutil "github.com/projectdiscovery/utils/maps"
)

type Agent struct {
}
type Parse struct {
	RegexResult     string
	JsonParseResult string
}

type Result struct {
	Type   ResultType
	Source string
	Value  string
	Error  error
}
type ResultType int

type EnumerationOptions struct {
	customRateLimiter *mapsutil.SyncLockMap[string, uint]
}

type EnumerateOption func(opts *EnumerationOptions)

func (a *Agent) Parse(ctx context.Context, options ...EnumerateOption) {
	results := make(chan Result)
	go func() {
		defer close(results)
		var enumerateOptions EnumerationOptions
		for _, enumerateOption := range options {
			enumerateOption(&enumerateOptions)
		}
	}()
}
