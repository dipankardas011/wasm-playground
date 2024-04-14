package main

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"sync"
)

func main() {
	slog.Info("data", "uid", os.Getuid())

	wg := new(sync.WaitGroup)
	urls := []string{
		"https://www.google.com",
		"https://www.youtube.com",
		"https://www.facebook.com",
		"https://ksctl.com",
	}
	ch := make(chan string, len(urls))

	for _, URL := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()

			if err := fetchDataFromDiffUrl(url, ch); err != nil {
				slog.Error("failed", "url", url, "reason", err)
			}
		}(URL)
	}

	// go func() {
	wg.Wait()
	close(ch)
	// }()

	for res := range ch {
		if len(res) != 0 {
			slog.Info(res)
		}
	}
}

func fetchDataFromDiffUrl(url string, ch chan<- string) (err error) {
	var resp *http.Response
	resp, err = http.Get(url)
	if err != nil {
		ch <- ""
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		ch <- ""
		return
	}

	ch <- fmt.Sprintf("url: %s -> Status: %s; Body: %s", url, resp.Status, body[:50])
	return
}
