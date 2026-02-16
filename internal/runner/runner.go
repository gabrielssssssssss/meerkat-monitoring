package runner

import (
	"context"

	"github.com/gabrielssssssssss/meerkat-monitoring/config"
	"github.com/gabrielssssssssss/meerkat-monitoring/internal/service"
)

type Runner struct {
	options    *Options
	config     *config.Config
	hitService *service.HitService
}

func NewRunner(options *Options, config *config.Config, hitService *service.HitService) *Runner {
	return &Runner{
		options:    options,
		config:     config,
		hitService: hitService,
	}
}

func (r *Runner) RunScanner() error {
	ctx, cancel := context.WithCancel(context.Background())
	return r.RunScannerWithCtx(ctx, cancel)
}

func (r *Runner) RunScannerWithCtx(ctx context.Context, cancel context.CancelFunc) error {
	return nil
}
