package main

import (
	"os"

	"github.com/whyhilde/goget/src"
)

func main() {
	cfg := src.ParseArgs()

	src.Download(cfg.URL, cfg.Output)

	os.Exit(0)
}
