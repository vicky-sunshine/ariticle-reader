package hn

import (
	"fmt"
	"time"
)

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
		return fmt.Sprintf("%v days ago", delta/timeDay)
	case delta < timeMonth:
		return fmt.Sprintf("%v week(s) ago", delta/timeWeek)
	case delta < timeYear:
		return fmt.Sprintf("%v month(s) ago", delta/timeMonth)
	default:
		return fmt.Sprintf("%v year(s) ago", delta/timeYear)
	}
}
