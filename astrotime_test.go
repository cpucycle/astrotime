package astrotime

import (
	"fmt"
	"math"
	"testing"
	"time"
)

type location struct {
	timezone  string
	latitude  float64
	longitude float64
}

type times struct {
	date    string
	sunrise string
	sunset  string
	dawn    string
	dusk    string
}

type test struct {
	times    times
	location location
	error    time.Duration
}

// test data from http://www.timeanddate.com/worldclock/sunrise.html

var (
	SYDNEY = location{
		timezone:  "Australia/Sydney",
		latitude:  -33.86,
		longitude: -151.20,
	}
	STOCKHOLM = location{
		timezone:  "Europe/Stockholm",
		latitude:  59.33,
		longitude: -18.067,
	}
	NEWYORK = location{
		timezone:  "America/New_York",
		latitude:  40.642,
		longitude: 74.017,
	}
	NOVEMBER = "2014-11-01"
	JULY     = "2015-07-01"
	MINUTE   = time.Minute
	DATA     = []*test{
		&test{times{NOVEMBER, "05:55", "19:23", "05:29", "19:49"}, SYDNEY, MINUTE},
		&test{times{NOVEMBER, "07:07", "15:55", "06:23", "16:39"}, STOCKHOLM, MINUTE},
		&test{times{NOVEMBER, "07:26", "17:52", "06:58", "18:21"}, NEWYORK, MINUTE},
		&test{times{JULY, "07:01", "16:57", "06:33", "17:25"}, SYDNEY, MINUTE},
		&test{times{JULY, "03:37", "22:06", "02:09", "23:33"}, STOCKHOLM, MINUTE},
		&test{times{JULY, "05:29", "20:31", "04:55", "21:04"}, NEWYORK, MINUTE},
	}
)

func TestData(t *testing.T) {
	for _, datum := range DATA {
		if loc, err := time.LoadLocation(datum.location.timezone); err != nil {
			t.Fatalf("bad location %s", datum.location.timezone)
		} else {
			formatted := fmt.Sprintf("%s 00:00:00", datum.times.date)
			midnight, err := time.ParseInLocation("2006-01-02 15:04:05", formatted, loc)
			if err != nil {
				t.Fatalf("failed to parse date/timestamp '%s': %s", formatted, err)
			}

			checkDate := func(description string, timeOfDay string, calculated time.Time) {
				formatted := fmt.Sprintf("%s %s", datum.times.date, timeOfDay)
				timestamp, err := time.ParseInLocation("2006-01-02 15:04", formatted, loc)
				if err != nil {
					t.Fatalf("failed to parse date/timestamp '%s': %s", formatted, err)
				}
				actualError := math.Abs((float64)(timestamp.Sub(calculated)))
				if actualError > float64(datum.error) {
					t.Errorf("calculated -> %v, wanted -> %v %d -> (wanted < %d). location=%s date=%s", calculated, timestamp, actualError, datum.error, datum.location.timezone, datum.times.date)
				}
			}

			checkDate("dawn", datum.times.dawn, NextDawn(midnight, datum.location.latitude, datum.location.longitude, CIVIL_DAWN))
			checkDate("sunrise", datum.times.sunrise, NextSunrise(midnight, datum.location.latitude, datum.location.longitude))
			checkDate("sunset", datum.times.sunset, NextSunset(midnight, datum.location.latitude, datum.location.longitude))
			checkDate("dusk", datum.times.dusk, NextDusk(midnight, datum.location.latitude, datum.location.longitude, CIVIL_DUSK))
		}

	}
}
