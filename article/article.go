package article

import (
	"fmt"
	"time"
)

type Article interface {
	GetID() string
	GetTitle() string
	GetAuthor() string
	GetTimestamp() int64
}

type ArticleReader interface {
	TopArticles(number int) ([]string, error)
	GetArticle(id string) (Article, error)
}

func Summerized(ar Article) string {
	return fmt.Sprintf("%v  %v by %v %v",
		ar.GetID(), ar.GetTitle(), ar.GetAuthor(), DurationFormat(ar.GetTimestamp()))
}

const (
	timeDay   = 24 * time.Hour
	timeWeek  = 7 * timeDay
	timeMonth = 30 * timeDay
	timeYear  = 365 * timeDay
)

func DurationFormat(ts int64) string {
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
		return "Yesterday"
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
