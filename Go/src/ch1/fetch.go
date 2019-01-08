// Fetch prints the content found at a URL.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	const (
		http_prefix = "http://"
	)

	for _, url := range os.Args[1:] {

		if !strings.HasPrefix(url, http_prefix) {
			url = http_prefix + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s:%v\n", url, err)
			os.Exit(1)
		}

		fmt.Printf("\nHTTP status code: %v\n", resp.Status)
		resp.Body.Close()
	}
}
