package sumday

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//TimeMatch is a regex that matches lines such as: " - 1230 - 3: category: what I did during this time"
const TimeMatch = "\\s*?-\\s?(\\d+)\\s?-\\s?(\\d+)\\s?:\\s?(.*?):"

type TimeOfDay struct {
	TimeString string
	Err        error
	Hour       int
	Minutes    int
	Time       time.Time
}

type Task struct {
	Start    *TimeOfDay
	End      *TimeOfDay
	Category string
}

func NewTimeOfDay(time string) *TimeOfDay {
	tod := &TimeOfDay{}
	tod.SetTime(time)
	return tod
}

func (t *TimeOfDay) SetTime(timestring string) *TimeOfDay {
	t.TimeString = timestring
	itime, err := strconv.Atoi(timestring)
	if err != nil {
		t.Err = err
		return t
	}

	h, m := "0", "0"
	s := strconv.Itoa(itime)
	switch len(timestring) {
	case 1, 2:
		h = s
	case 3:
		h = s[0:1]
		m = s[1:]

	case 4:
		h = s[0:2]
		m = s[2:]
	}
	hh, _ := strconv.Atoi(h)
	mm, _ := strconv.Atoi(m)
	t.Hour = hh
	t.Minutes = mm
	t.Time = time.Date(2042, time.January, 1, hh, mm, 0, 0, time.UTC)
	return t
}

func MakeTimes(start, end string) (*TimeOfDay, *TimeOfDay) {
	s := NewTimeOfDay(start)
	e := NewTimeOfDay(end)
	if s.Hour > e.Hour {
		d, _ := time.ParseDuration("12h")
		e.Time = e.Time.Add(d)
	}
	return s, e
}

//ParseLine Takes a line such as  ' - 2 - 3: RM: Weekly RM Meeting' and returns a Task representing start, end, and category
//Categories are normalized by being lowercased and having all non-alpha chars removed

func ParseLine(input string) Task {
	re := regexp.MustCompile(TimeMatch)
	matches := re.FindStringSubmatch(input)

	if len(matches) == 0 {
		return Task{NewTimeOfDay("0:00"), NewTimeOfDay("0:00"), "UNKNOWN CATEGORY"}
	}

	s, e := MakeTimes(strings.TrimSpace(matches[1]), strings.TrimSpace(matches[2]))
	return Task{s, e, Normalize(matches[3])}
}

//ParseAllLines takes a block of lines for a given day (see example below) and returns a map of summed durations keyed on category
//eg: {'cfgov': 6.5, 'RM': 2.25}
/*
Example time block (See timesheet_test.go for full example):

 - 9 - 915: Team: Email facilities about parking and whomever about monitor
...
  - 2 - 3: RM: Weekly RM Meeting
           - Production patching... Qu got skipped so we're doing today
  - 3 - 330: RM: Sally 1:1
...

*/
func SumDay(input string) map[string]float64 {
	hours := make(map[string]float64)
	re := regexp.MustCompile(TimeMatch)

	for _, v := range strings.Split(input, "\n") {
		if strings.TrimSpace(v) == "" {
			continue
		} else if re.MatchString(v) {
			task := ParseLine(v)
			hours[task.Category] += task.End.Time.Sub(task.Start.Time).Hours()
		} else {
			fmt.Println("Could not parse line! ", v)
		}
	}
	return hours
}

//Normalize normalizes a category string by lowercasing and removing all non-alpha characters
func Normalize(category string) string {
	re := regexp.MustCompile("[^a-zA-Z0-9\\s]")
	clean := re.ReplaceAllString(category, "")
	return strings.ToLower(clean)
}
