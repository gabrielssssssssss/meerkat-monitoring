package runner

import (
	"context"
	"sync"

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
	domains := make(chan string, 1000000)

	var wg sync.WaitGroup

	if err := r.GitScanner(ctx, domains, &wg); err != nil {
		return err
	}

	if len(r.config.CtLogs) > 0 && len(r.options.DomainsFile) == 0 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			r.MonitoringTransparency(ctx, r.config.CtLogs, domains)
		}()
	} else if len(r.options.DomainsFile) > 0 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = r.LoadDomains(r.options.DomainsFile, domains)
			close(domains)
		}()
	}

	go func() {
		r.Scheduler(ctx)
	}()

	wg.Wait()

	return nil
}
