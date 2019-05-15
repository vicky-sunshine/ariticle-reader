// Package article define the main interface for artread projects
package article

import (
	"fmt"
	"time"
)

// Article is used for specifying requirement field of an article
type Article interface {
	GetID() string
	GetTitle() string
	GetAuthor() string
	GetTimestamp() int64
}

// Reader is an interface which defines required function to fetch articles from certain website
type Reader interface {
	// List current top articles from target website
	TopArticles(number int) ([]string, error)
	// Get article with specified ID from target website
	GetArticle(id string) (Article, error)
}

// Summarized function is used to one-line summarize ariticle when listing aritcles
func Summarized(ar Article) string {
	return fmt.Sprintf("%v  %v by %v %v",
		ar.GetID(), ar.GetTitle(), ar.GetAuthor(), durationFormat(ar.GetTimestamp()))
}

const (
	timeDay   = 24 * time.Hour
	timeWeek  = 7 * timeDay
	timeMonth = 30 * timeDay
	timeYear  = 365 * timeDay
)

// durationFormat makes unix timestamp to duration from now
// ex: "just now", "yesterday", "2 month(s) ago", ..etc
func durationFormat(ts int64) string {
	switch delta := time.Since(time.Unix(ts, 0)); {
	case delta < 10*time.Second:
		return "just now"
	case delta < time.Minute:
		return fmt.Sprintf("%v seonds ago", int(delta.Seconds()))
	case delta < time.Hour:
		return fmt.Sprintf("%v minutes ago", int(delta.Minutes()))
	case delta < timeDay:
		return fmt.Sprintf("%v hours ago", int(delta.Hours()))
	case delta < 2*timeDay:
		return "yesterday"
	case delta < timeWeek:
		return fmt.Sprintf("%v days ago", int(delta/timeDay))
	case delta < timeMonth:
		return fmt.Sprintf("%v week(s) ago", int(delta/timeWeek))
	case delta < timeYear:
		return fmt.Sprintf("%v month(s) ago", int(delta/timeMonth))
	default:
		return fmt.Sprintf("%v year(s) ago", int(delta/timeYear))
	}
}
