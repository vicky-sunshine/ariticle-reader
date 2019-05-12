package hn

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

var apiBase = "https://hacker-news.firebaseio.com/v0"

func ItemFormat(item Item) string {
	u, _ := url.Parse(item.URL)
	return fmt.Sprintf("%8d  %v (%v)\n          %d points by %v %v ago | %v comments",
		item.ID, item.Title, u.Hostname(),
		item.Score, item.By, DurationFormat(item.Time), item.Descendants)
}

func TopItems() ([]int, error) {
	resp, err := http.Get(fmt.Sprintf("%s/topstories.json", apiBase))
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
	return ids, nil
}

func GetItem(id int) (Item, error) {
	var item Item
	resp, err := http.Get(fmt.Sprintf("%s/item/%d.json", apiBase, id))
	if err != nil {
		return item, err
	}
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&item)
	if err != nil {
		return item, err
	}
	return item, nil
}

type Item struct {
	By          string `json:"by"`
	Descendants int    `json:"descendants"`
	ID          int    `json:"id"`
	Kids        []int  `json:"kids"`
	Score       int    `json:"score"`
	Time        int64  `json:"time"`
	Title       string `json:"title"`
	Type        string `json:"type"`

	// Only one of these should exist
	Text string `json:"text"`
	URL  string `json:"url"`
}
