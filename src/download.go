package src

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Download(url string, output string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	out, err := os.Create(output)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	defer out.Close()

	io.Copy(out, resp.Body)
}
