package runner

import (
	"context"
	"sync"
	"time"

	"github.com/gabrielssssssssss/meerkat-monitoring/internal/models"
	"github.com/rs/zerolog/log"
)

func (r *Runner) GitScanner(ctx context.Context, domains chan string, wg *sync.WaitGroup) error {
	return r.GitScannerWithCtx(ctx, domains, wg)
}

func (r *Runner) GitScannerWithCtx(ctx context.Context, domains chan string, wg *sync.WaitGroup) error {
	for i := 0; i < r.options.Threads; i++ {
		wg.Add(1)
		go func(int) {
			for domain := range domains {
				url := "https://" + domain

				isExposed, _ := r.gitHarvest.IsGitExposed(url)
				if !isExposed {
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
						found, _ := r.hitService.FindByToken(token)
						if found != nil {
							log.Info().
								Str("token", token).
								Str("on", url).
								Msg("Duplicated token found")

							continue
						}

						log.Info().
							Str("token", token).
							Str("on", url).
							Msg("New token found")

						if r.hitService.Create(&models.Hit{
							URL:       url,
							Token:     token,
							CreatedAt: time.Now(),
						}) != nil {
							continue
						}
					}
				}
			}
		}(i)
	}

	return nil
}
