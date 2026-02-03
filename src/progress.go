package src

import (
	"fmt"
	"io"
	"time"
)

type ProgressReader struct {
	io.Reader
	Total     int64
	Current   int64
	StartTime time.Time
}

func (pr *ProgressReader) printProgress() {
	if pr.Total <= 0 {
		fmt.Printf("\rDownloaded: %s", formatBytes(pr.Current))
		return
	}
}
