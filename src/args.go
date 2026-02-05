package src

import (
	"flag"
	"fmt"
	"os"
)

type Config struct {
	URL     string
	Output  string
	Verbose bool
	Help    bool
}

func ParseArgs() *Config {
	cfg := &Config{}

	flag.StringVar(&cfg.Output, "o", "output", "output filename")
	flag.StringVar(&cfg.Output, "output", "output", "output filename")

	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Fprintf(os.Stderr, "Error: URL required\n")
		os.Exit(1)
	}

	cfg.URL = args[0]

	return cfg
}
