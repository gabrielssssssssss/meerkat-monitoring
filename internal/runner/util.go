package runner

import (
	fileutil "github.com/projectdiscovery/utils/file"
	stringsutil "github.com/projectdiscovery/utils/strings"
)

// github.com/projectdiscovery/subfinder/blob/dev/pkg/runner/util.go
func (r *Runner) LoadDomains(file string, ch chan string) error {
	lines, err := fileutil.ReadFile(file)
	if err != nil {
		return err
	}

	for domain := range lines {
		domain = preprocessDomain(domain)
		if domain == "" {
			continue
		}

		ch <- domain
	}
	return nil
}

// github.com/projectdiscovery/subfinder/blob/dev/pkg/runner/util.go
func preprocessDomain(s string) string {
	return stringsutil.NormalizeWithOptions(s,
		stringsutil.NormalizeOptions{
			StripComments: true,
			TrimCutset:    "\n\t\"'` ",
			Lowercase:     true,
		},
	)
}
