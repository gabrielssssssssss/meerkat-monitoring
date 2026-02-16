package runner

import (
	"fmt"
	"os"

	"github.com/projectdiscovery/goflags"
)

type Options struct {
	Cfg         string
	NoColor     bool
	Silent      bool
	Threads     int
	Timeout     int
	Domain      string
	DomainsFile string
}

func ParseOptions() *Options {
	options := &Options{}

	flagSet := goflags.NewFlagSet()
	flagSet.SetDescription(`Meerkat is a GIT exposures monitoring that discovers private credentials.`)

	flagSet.CreateGroup("input", "Input",
		flagSet.StringVarP(&options.Domain, "domain", "d", "", "domains to find subdomains for"),
		flagSet.StringVarP(&options.DomainsFile, "list", "dL", "", "file containing list of domains for credentials discovery"),
	)

	flagSet.CreateGroup("debug", "Debug",
		flagSet.BoolVar(&options.Silent, "silent", false, "show only valid tokens in output"),
		flagSet.BoolVarP(&options.NoColor, "no-color", "nc", false, "disable color in output"),
	)

	flagSet.CreateGroup("optimization", "Optimization",
		flagSet.IntVar(&options.Timeout, "timeout", 30, "seconds to wait before timing out"),
		flagSet.IntVar(&options.Threads, "t", 10, "number of concurrent goroutines for resolving (-active only)"),
	)

	flagSet.CreateGroup("config", "Config",
		flagSet.StringVarP(&options.Cfg, "conf", "c", "", "environnement variable (.yaml)"),
	)

	if err := flagSet.Parse(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return options
}
