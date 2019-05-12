package hn

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type HNArticle struct {
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

func (hna HNArticle) GetID() string{
	return strconv.Itoa(hna.ID)
}
func (hna HNArticle) GetTitle() string{
	return hna.Title
}
func (hna HNArticle) GetAuthor() string{
	return hna.By
}
func (hna HNArticle) GetTimestamp() int64{
	return hna.Time
}

type HackerNewsReader struct {
	apiBase string
}

func NewHackerNewsReader() *HackerNewsReader {
	hnr := &HackerNewsReader{}
	hnr.apiBase = "https://hacker-news.firebaseio.com/v0"
	return hnr
}

func (hnr *HackerNewsReader) GetArticle(id string) (HNArticle, error) {
	var ar HNArticle
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

func (hnr *HackerNewsReader) TopArticles(number int) ([]HNArticle, error) {
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

	var ars []HNArticle
	for _, v := range ids[:number] {
		ar, err := hnr.GetArticle(strconv.Itoa(v))
		if err != nil {
			return ars, err
		}
		ars = append(ars, ar)
	}
	return ars, nil
}

