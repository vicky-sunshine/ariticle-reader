// Package rssfetch helps fetch rss/atom format data
package rssfetch

import (
	"fmt"
	"net/http"

	"github.com/mmcdole/gofeed"
)

// HTTPError wraps error info for http
type HTTPError struct {
	StatusCode int
	Status     string
}

// Error used to implements error interaface
func (err HTTPError) Error() string {
	return fmt.Sprintf("http error: %s", err.Status)
}

// Fetch is used to fetch and parse rss data from given url
// Ref for returns *gofeed.Feed type: https://github.com/mmcdole/gofeed/blob/master/feed.go
func Fetch(url string) (feed *gofeed.Feed, err error) {
	fp := gofeed.NewParser()

	// make client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Headers
	// Fake user agent for web crawling
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.186 Safari/537.36")

	// Fetch Request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp != nil {
		defer func() {
			ce := resp.Body.Close()
			if ce != nil {
				err = ce
			}
		}()
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, HTTPError{
			StatusCode: resp.StatusCode,
			Status:     resp.Status,
		}
	}

	feed, err = fp.Parse(resp.Body)
	if err != nil {
		return feed, err
	}
	return feed, nil
}
