package runner

import (
	"github.com/t101804/xorhunt/options"
	"github.com/t101804/xorhunt/pkg/resolve"
)

type Runner struct {
	options        *options.GlobalOptions
	resolverClient *resolve.Resolver
}

func NewRunner(options *options.GlobalOptions) (*Runner, error) {

	return nil, nil
}
