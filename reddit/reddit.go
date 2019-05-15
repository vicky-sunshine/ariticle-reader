// Package reddit provide the client and function for reading articles from reddit rss
// Ref:
// - reddit rss: https://www.reddit.com/wiki/rss
// - atom feed format: https://en.wikipedia.org/wiki/Atom_(Web_standard)
package reddit

import (
	"artread/rssfetch"
	"fmt"
	"strings"
	"time"
)

// RdtArticle define the field of reddit post, and
// implements the Article interface
type RdtArticle struct {
	ID     string
	Time   int64
	Title  string
	Author string
}

func (rda RdtArticle) GetID() string {
	return rda.ID
}
func (rda RdtArticle) GetTitle() string {
	return rda.Title
}
func (rda RdtArticle) GetAuthor() string {
	return rda.Author
}
func (rda RdtArticle) GetTimestamp() int64 {
	return rda.Time
}

// RdtReader implements Reader interface to read reddit rss
// You can take artread/rssfetch as tool to fetch
type RdtReader struct {
	apiBase string
}

func NewRdtReader(apiBase string) *RdtReader {
	rdr := &RdtReader{}
	rdr.apiBase = apiBase
	return rdr
}

func (rdr *RdtReader) GetArticle(id string) (RdtArticle, error) {
	feed, err := rssfetch.Fetch(fmt.Sprintf("%s/comments/%s/.rss", rdr.apiBase, id))

	var rds RdtArticle
	if err != nil {
		return rds, err
	}

	rds.ID = id
	rds.Title = feed.Items[0].Title
	rds.Author = feed.Items[0].Author.Name

	t, err := time.Parse(time.RFC3339, feed.Items[0].Updated)
	if err != nil {
		return rds, err
	}
	rds.Time = t.Unix()

	return rds, nil
}

func (rdr *RdtReader) TopArticles(number int) ([]string, error) {
	feed, err := rssfetch.Fetch(fmt.Sprintf("%s/.rss", rdr.apiBase))
	if err != nil {
		return nil, err
	}
	var idStrs []string
	for _, v := range feed.Items[:number] {
		spID := strings.Split(v.GUID, "_")
		if len(spID) < 2 {
			return idStrs, fmt.Errorf("ID format error %v", v.GUID)
		}
		idStrs = append(idStrs, spID[1])
	}
	return idStrs, nil
}
