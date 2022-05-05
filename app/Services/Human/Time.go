package Human

import (
	"fmt"
	"math"
	"sort"
	"time"
)

// Seconds-based time units
const (
	Day      = 24 * time.Hour
	Week     = 7 * Day
	Month    = 30 * Day
	Year     = 12 * Month
	LongTime = 37 * Year
)

// Time formats a time into a relative string.
//
// Time(someT) -> "3 weeks ago"
func Time(then time.Time) string {
	return RelTime(then, time.Now(), "назад", "назад")
}

// A RelTimeMagnitude struct contains a relative time point at which
// the relative format of time will switch to a new format string.  A
// slice of these in ascending order by their "D" field is passed to
// CustomRelTime to format durations.
//
// The Format field is a string that may contain a "%s" which will be
// replaced with the appropriate signed label (e.g. "ago" or "from
// now") and a "%d" that will be replaced by the quantity.
//
// The DivBy field is the amount of time the time difference must be
// divided by in order to display correctly.
//
// e.g. if D is 2*time.Minute and you want to display "%d minutes %s"
// DivBy should be time.Minute so whatever the duration is will be
// expressed in minutes.
type RelTimeMagnitude struct {
	D      time.Duration
	Format string
	DivBy  time.Duration
}

var defaultMagnitudes = []RelTimeMagnitude{
	{time.Second, "сейчас", time.Second},
	{2 * time.Second, "1 секунду %s", 1},
	{3 * time.Second, "2 секунды %s", 1},
	{4 * time.Second, "3 секунды %s", 1},
	{5 * time.Second, "4 секунды %s", 1},
	{time.Minute, "%d секунд %s", time.Second},
	{2 * time.Minute, "1 минуту %s", 1},
	{3 * time.Minute, "2 минуты %s", 1},
	{4 * time.Minute, "3 минуты %s", 1},
	{5 * time.Minute, "4 минуты %s", 1},
	{time.Hour, "%d минут %s", time.Minute},
	{2 * time.Hour, "1 час %s", 1},
	{3 * time.Hour, "2 часа %s", 1},
	{4 * time.Hour, "3 часа %s", 1},
	{5 * time.Hour, "4 часа %s", 1},
	{Day, "%d часов %s", time.Hour},
	{2 * Day, "1 день %s", 1},
	{3 * Day, "2 дня %s", 1},
	{4 * Day, "3 дня %s", 1},
	{5 * Day, "4 дня %s", 1},
	{Week, "%d дней %s", Day},
	{2 * Week, "1 неделю %s", 1},
	{Month, "%d недели %s", Week},
	{2 * Month, "1 месяц %s", 1},
	{3 * Month, "2 месяца %s", 1},
	{4 * Month, "3 месяца %s", 1},
	{5 * Month, "4 месяца %s", 1},
	{Year, "%d месяцев %s", Month},
	{18 * Month, "1 год %s", 1},
	{2 * Year, "2 года %s", 1},
	{3 * Year, "3 года %s", 1},
	{4 * Year, "4 года %s", 1},
	{LongTime, "%d лет %s", Year},
	{math.MaxInt64, "очень много времени %s", 1},
}

// RelTime formats a time into a relative string.
//
// It takes two times and two labels.  In addition to the generic time
// delta string (e.g. 5 minutes), the labels are used applied so that
// the label corresponding to the smaller time is applied.
//
// RelTime(timeInPast, timeInFuture, "earlier", "later") -> "3 weeks earlier"
func RelTime(a, b time.Time, albl, blbl string) string {
	//b = b.Add(3 * time.Hour) //GMT+3
	return CustomRelTime(a, b, albl, blbl, defaultMagnitudes)
}

// CustomRelTime formats a time into a relative string.
//
// It takes two times two labels and a table of relative time formats.
// In addition to the generic time delta string (e.g. 5 minutes), the
// labels are used applied so that the label corresponding to the
// smaller time is applied.
func CustomRelTime(a, b time.Time, albl, blbl string, magnitudes []RelTimeMagnitude) string {
	lbl := albl
	diff := b.Sub(a)

	if a.After(b) {
		lbl = blbl
		diff = a.Sub(b)
	}

	n := sort.Search(len(magnitudes), func(i int) bool {
		return magnitudes[i].D > diff
	})

	if n >= len(magnitudes) {
		n = len(magnitudes) - 1
	}
	mag := magnitudes[n]
	args := []interface{}{}
	escaped := false
	for _, ch := range mag.Format {
		if escaped {
			switch ch {
			case 's':
				args = append(args, lbl)
			case 'd':
				args = append(args, diff/mag.DivBy)
			}
			escaped = false
		} else {
			escaped = ch == '%'
		}
	}
	return fmt.Sprintf(mag.Format, args...)
}
