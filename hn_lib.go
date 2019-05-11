package main

import (
	"strconv"
	"time"
)

const day = time.Minute * 60 * 24
const week = day * 7
const month = day * 31
const year = day * 365

func duration_format(ts int) string {

	since := time.Since(time.Unix(int64(ts), 0))

	second_diff := since.Seconds()

	if since < day {
		switch {
		case second_diff < 10:
			return "just now"
		case second_diff < 60:
			return strconv.Itoa(int(second_diff)) + " seconds ago"
		case second_diff < 120:
			return "1 minute ago"
		case second_diff < 3600:
			return strconv.Itoa(int(second_diff/60)) + " minutes ago"
		case second_diff < 7200:
			return "1 hour ago"
		case second_diff < 86400:
			return strconv.Itoa(int(second_diff/3600)) + " hours ago"
		}
	}
	if since < 2*day {
		return "Yesterday"
	}
	if since < week {
		return strconv.Itoa(int(second_diff/86400)) + " days ago"
	}
	if since < month {
		return strconv.Itoa(int(second_diff/86400/7)) + " week(s) ago"
	}
	if since < year {
		return strconv.Itoa(int(second_diff/86400/30)) + " month(s) ago"
	}
	return strconv.Itoa(int(second_diff/86400/365)) + " year(s) ago"
}
