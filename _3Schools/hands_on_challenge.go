package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {	
	urls := gen2020URLs()
	size, err := downloadsSize(urls)
	if err != nil {
		log.Fatal(err)
	}
	sizeGB := float64(size) / (1 << 30)
	fmt.Printf("size = %.2fGB\n", sizeGB)
}

func downloadSize(url string) (int, error) {
	resp, err := http.Head(url)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("Bad Status: %d %s",
			resp.StatusCode, resp.Status)
	}
	return strconv.Atoi(resp.Header.Get("Content-Length"))
}

func downloadsSize(urls []string) (int, error) {
	var dlSize int64
	var g errgroup.Group

	for _, url := range urls {
		url := url // https://golang.org/doc/faq#closures_and_goroutines
		g.Go(func() error {
			log.Print(url)
			size, err := downloadSize(url)
			if err != nil {
				return err
			}
			atomic.AddInt64(&dlSize, int64(size))
			return nil
		})

	}
	if err := g.Wait(); err != nil {
		return 0, err
	}
	return int(dlSize), nil
}

func gen2020URLs() []string {
	var urls []string
	urlTemplate := "https://www1.nyc.gov/site/tlc/about/tlc-trip-record-data.page"
	for _, vendor := range []string{"yellow", "green"} {
		for month := 1; month <= 12; month++ {
			url := fmt.Sprintf(urlTemplate,
				vendor, month)
			urls = append(urls, url)
		}
	}
	return urls
}
