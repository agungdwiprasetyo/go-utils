package parser

import (
	"fmt"
	"regexp"
	"time"
)

var timeFormats = []string{"1/2/2006", "1/2/2006 15:4:5", "2006", "2006-1", "2006-1-2", "2006-1-2 15", "2006-1-2 15:4", "2006-1-2 15:4:5", "1-2", "15:4:5", "15:4", "15", "15:4:5 Jan 2, 2006 MST", "2006-01-02 15:04:05.999999999 -0700 MST"}
var hasTimeRegexp = regexp.MustCompile(`(\s+|^\s*)\d{1,2}((:\d{1,2})*|((:\d{1,2}){2}\.(\d{3}|\d{6}|\d{9})))\s*$`)
var onlyTimeRegexp = regexp.MustCompile(`^\s*\d{1,2}((:\d{1,2})*|((:\d{1,2}){2}\.(\d{3}|\d{6}|\d{9})))\s*$`)

func parseTimeWithFormat(str string) (t time.Time, err error) {
	for _, format := range timeFormats {
		t, err = time.Parse(format, str)
		if err == nil {
			return
		}
	}
	err = fmt.Errorf("Can't parse string as time: %s", str)
	return
}

// ParseTime parse string to time
func ParseTime(strs ...string) (t time.Time, err error) {
	now := time.Now()
	var (
		setCurrentTime  bool
		parseTime       []int
		currentTime     = []int{now.Nanosecond(), now.Second(), now.Minute(), now.Hour(), now.Day(), int(now.Month()), now.Year()}
		currentLocation = now.Location()
		onlyTimeInStr   = true
	)

	for _, str := range strs {
		hasTimeInStr := hasTimeRegexp.MatchString(str)
		onlyTimeInStr = hasTimeInStr && onlyTimeInStr && onlyTimeRegexp.MatchString(str)
		if t, err = parseTimeWithFormat(str); err == nil {
			location := t.Location()
			if location.String() == "UTC" {
				location = currentLocation
			}

			parseTime = []int{t.Nanosecond(), t.Second(), t.Minute(), t.Hour(), t.Day(), int(t.Month()), t.Year()}

			for i, v := range parseTime {
				// Don't reset hour, minute, second if current time str including time
				if hasTimeInStr && i <= 3 {
					continue
				}

				// If value is zero, replace it with current time
				if v == 0 {
					if setCurrentTime {
						parseTime[i] = currentTime[i]
					}
				} else {
					setCurrentTime = true
				}

				// if current time only includes time, should change day, month to current time
				if onlyTimeInStr {
					if i == 4 || i == 5 {
						parseTime[i] = currentTime[i]
						continue
					}
				}
			}

			t = time.Date(parseTime[6], time.Month(parseTime[5]), parseTime[4], parseTime[3], parseTime[2], parseTime[1], parseTime[0], location)
			currentTime = []int{t.Nanosecond(), t.Second(), t.Minute(), t.Hour(), t.Day(), int(t.Month()), t.Year()}
		}
	}

	if err != nil {
		t = time.Date(1980, time.Month(1), 1, 0, 0, 0, 0, currentLocation)
	}

	return
}
