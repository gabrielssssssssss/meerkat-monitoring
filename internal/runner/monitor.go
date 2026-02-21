package runner

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/gabrielssssssssss/meerkat-monitoring/internal/models"
	"github.com/gabrielssssssssss/meerkat-monitoring/pkg/transparency"
)

func (r *Runner) MonitoringScanner(sources []string) error {
	return r.MonitoringScannerWithCtx(context.Background(), sources)
}

func (r *Runner) MonitoringScannerWithCtx(ctx context.Context, sources []string) error {
	domains := make(chan string, 1000000)

	for i := 0; i < r.options.Threads; i++ {
		go func(int) {
			for domain := range domains {
				url := "https://" + domain

				isExposed, err := r.gitHarvest.IsGitExposed(url)
				fmt.Println(url, isExposed)

				if err != nil || (!isExposed) {
					continue
				}

				var validTokens []string

				if isExposed {
					for _, path := range r.config.GitPath {
						tokens, err := r.gitHarvest.ExtractTokens(url, path)
						if err != nil {
							continue
						}

						validTokens = append(validTokens, tokens...)
					}
				}

				if len(validTokens) > 0 {
					for _, token := range validTokens {
						isValid, err := r.gitHarvest.IsValidToken(token)
						if err != nil || (!isValid) {
							continue
						}

						tokenInfo, err := r.gitHarvest.GetTokenInfo(token)
						if err != nil {
							continue
						}

						err = r.hitService.Create(&models.Hit{
							URL:        url,
							Token:      token,
							UserGithub: *tokenInfo,
							CreatedAt:  time.Now(),
						})

						if err != nil {
							continue
						}
					}
				}
			}
		}(i)
	}

	r.MonitoringTransparency(sources, domains)

	return nil
}

func (r *Runner) MonitoringTransparency(sources []string, ch chan string) error {
	return r.MonitoringTransparencyWithCtx(context.Background(), sources, ch)
}

func (r *Runner) MonitoringTransparencyWithCtx(ctx context.Context, sources []string, ch chan string) error {
	var (
		sourceInf = make(map[string]int64)
		mu        sync.Mutex
	)

	fmt.Println(sources)
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

					cleanDomain := strings.Replace(domain, "*.", "", 1)
					err = r.transparencyService.Create(&models.Transparency{
						Domain:    cleanDomain,
						CreatedAt: time.Now(),
					})

					if err != nil {
						continue
					}

					ch <- cleanDomain
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
