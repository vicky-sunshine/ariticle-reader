package reddit

import (
	"fmt"
	"hn-reader/rssfetch"
	"strings"
	"time"
)

type RedditArticle struct {
	ID     string
	Time   int64
	Title  string
	Author string
}

func (rda RedditArticle) GetID() string {
	return rda.ID
}
func (rda RedditArticle) GetTitle() string {
	return rda.Title
}
func (rda RedditArticle) GetAuthor() string {
	return rda.Author
}
func (rda RedditArticle) GetTimestamp() int64 {
	return rda.Time
}

type RedditReader struct {
	apiBase string
}

func NewRedditReader(apiBase string) *RedditReader {
	rdr := &RedditReader{}
	rdr.apiBase = apiBase
	return rdr
}

func (rdr *RedditReader) GetArticle(id string) (RedditArticle, error) {
	feed, err := rssfetch.Fetch(fmt.Sprintf("%s/comments/%s/.rss", rdr.apiBase, id))

	var rds RedditArticle
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

func (rdr *RedditReader) TopArticles(number int) ([]string, error) {
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
