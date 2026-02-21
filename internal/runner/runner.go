package runner

import (
	"context"

	"github.com/gabrielssssssssss/meerkat-monitoring/config"
	"github.com/gabrielssssssssss/meerkat-monitoring/internal/service"
	"github.com/gabrielssssssssss/meerkat-monitoring/pkg/githarvest"
	"github.com/gabrielssssssssss/meerkat-monitoring/pkg/transparency"
)

type Runner struct {
	options             *Options
	config              *config.Config
	hitService          service.HitService
	transparencyService service.TransparencyService
	gitHarvest          *githarvest.Client
	transparency        *transparency.Client
}

func NewRunner(
	options *Options,
	config *config.Config,
	hitService service.HitService,
	transparencyService service.TransparencyService,
	gitHarvest *githarvest.Client,
	transparency *transparency.Client) *Runner {
	return &Runner{
		options:             options,
		config:              config,
		hitService:          hitService,
		transparencyService: transparencyService,
		gitHarvest:          gitHarvest,
		transparency:        transparency,
	}
}

func (r *Runner) RunScanner() error {
	ctx, cancel := context.WithCancel(context.Background())
	return r.RunScannerWithCtx(ctx, cancel)
}

func (r *Runner) RunScannerWithCtx(ctx context.Context, cancel context.CancelFunc) error {
	return r.MonitoringScanner(r.config.CtLogs)
}
