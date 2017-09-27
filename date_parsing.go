package gtd_bot

import (
	"errors"
	"regexp"
	"strconv"
	"time"
)

func parseDateInput(input string) (time.Time, error) {
	dateRegexp, err := regexp.Compile("(\\d{4})-(\\d{2})-(\\d{2})")
	if err != nil {
		panic("bad regular expression")
	}
	matches := dateRegexp.FindAllStringSubmatch(input, 1)
	if len(matches) == 0 {
		return time.Now(), errors.New("bad string")
	}
	year, _ := strconv.Atoi(matches[0][1])
	month, _ := strconv.Atoi(matches[0][2])
	day, _ := strconv.Atoi(matches[0][3])
	return time.Date(year, time.Month(month), day, 23, 59, 59, 0, time.UTC), nil
}

func endOfToday() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())
}

func nextDay(t time.Time) time.Time {
	return t.Add(time.Duration(24) * time.Hour)
}

func endOfTomorrow() time.Time {
	return nextDay(endOfToday())
}

func endOfWeek() time.Time {
	endOfDay := endOfToday()
	for endOfDay.Weekday() != time.Sunday {
		endOfDay = nextDay(endOfDay)
	}
	return endOfDay
}
