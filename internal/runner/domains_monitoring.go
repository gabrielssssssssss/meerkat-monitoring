package runner

import (
	"context"
	"sync"
	"time"

	"github.com/gabrielssssssssss/meerkat-monitoring/internal/models"
	"github.com/gabrielssssssssss/meerkat-monitoring/pkg/transparency"
)

func (r *Runner) MonitoringTransparency(ctx context.Context, sources []string, ch chan string) error {
	return r.MonitoringTransparencyWithCtx(ctx, sources, ch)
}

func (r *Runner) MonitoringTransparencyWithCtx(ctx context.Context, sources []string, ch chan string) error {
	var (
		sourceInf = make(map[string]int64)
		mu        sync.Mutex
	)

	for _, source := range sources {
		tree, err := r.transparency.GetTreeSize(source)
		if err != nil {
			continue
		}

		sourceInf[source] = tree.TreeSize
	}

	for {
		var wg sync.WaitGroup

		for _, source := range sources {
			wg.Add(1)

			go func(source string) {
				defer wg.Done()

				mu.Lock()
				start := sourceInf[source]
				end := start + 20
				mu.Unlock()

				entries, err := r.transparency.GetEntries(source, start, end)
				if err != nil {
					return
				}

				for _, entrie := range entries.Entries {
					domain, err := transparency.ParseLeafInput(entrie.LeafInput)
					if err != nil {
						continue
					}

					if r.transparencyService.Create(&models.Transparency{
						Domain:    domain,
						CreatedAt: time.Now(),
					}) != nil {
						continue
					}

					ch <- domain
				}

				if len(entries.Entries) != 0 {
					mu.Lock()
					sourceInf[source] += 20
					mu.Unlock()
				}
			}(source)
		}

		wg.Wait()
	}
}
