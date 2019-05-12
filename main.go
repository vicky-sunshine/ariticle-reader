package main

import (
	"fmt"
	"hn-reader/article"
	"hn-reader/hackernews"
)

func main() {
	hnr := hackernews.NewHackerNewsReader()
	list, _ := hnr.TopArticles(10)
	for _, v := range list {
		ar, _ := hnr.GetArticle(v)
		fmt.Println(article.Summerized(ar))
	}
}
