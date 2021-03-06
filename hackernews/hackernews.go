// Package hackernews provide the client and function for reading articles from hackernews
// Ref: https://github.com/HackerNews/API
package hackernews

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Article define the field of hacker news item, and
// implements the Article interface
// Ref: https://github.com/HackerNews/API#items
type Article struct {
	By          string `json:"by"`
	Descendants int    `json:"descendants"`
	ID          int    `json:"id"`
	Kids        []int  `json:"kids"`
	Score       int    `json:"score"`
	Time        int64  `json:"time"`
	Title       string `json:"title"`

	// Only one of these should exist
	Text string `json:"text"`
	URL  string `json:"url"`
}

// GetID implements Reader interface GetID funtion
func (hna Article) GetID() string {
	return strconv.Itoa(hna.ID)
}

// GetTitle implements Reader interface GetTitle function
func (hna Article) GetTitle() string {
	return hna.Title
}

// GetAuthor implements Reader interface GetAuthor function
func (hna Article) GetAuthor() string {
	return hna.By
}

// GetTimestamp implements Reader interface GetTimestamp function
func (hna Article) GetTimestamp() int64 {
	return hna.Time
}

// Reader implements Reader interface to read hackernews website
type Reader struct {
	apiBase string
}

// NewReader initialize an Reader instance
func NewReader(apiBase string) *Reader {
	hnr := &Reader{}
	hnr.apiBase = apiBase
	return hnr
}

// GetArticle returns a certain Article with given id
// Ref: https://github.com/HackerNews/API#items
func (hnr *Reader) GetArticle(id string) (Article, error) {
	var ar Article
	resp, err := http.Get(fmt.Sprintf("%s/item/%s.json", hnr.apiBase, id))
	if err != nil {
		return ar, err
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&ar)
	if err != nil {
		return ar, err
	}
	return ar, nil
}

// TopArticles returns a list of top ariticles ids
// Ref: https://github.com/HackerNews/API#new-top-and-best-stories
func (hnr *Reader) TopArticles(number int) ([]string, error) {
	resp, err := http.Get(fmt.Sprintf("%s/topstories.json", hnr.apiBase))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var ids []int
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&ids)
	if err != nil {
		return nil, err
	}

	var idStrs []string
	for _, v := range ids[:number] {
		idStrs = append(idStrs, strconv.Itoa(v))
	}

	return idStrs, nil
}
