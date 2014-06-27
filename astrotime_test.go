package astrotime

import (
	"testing"
	"time"
)

// Fixed location used for testing.
const TEST_LAT = 37.3894
const TEST_LONG = 122.0819

func compareApproxTimes(value, expected time.Time) bool {
	approx := time.Minute

	return value.After(expected.Add(-1*approx)) && value.Before(expected.Add(approx))
}

func TestCalcSunrise(t *testing.T) {

	expectedSunrise := time.Date(2014, time.June, 12, 05, 47, 00, 0, time.Local)

	// Test well before sunrise.
	now := time.Date(2014, time.June, 12, 2, 57, 12, 0, time.Local)
	sunrise := CalcSunrise(now, TEST_LAT, TEST_LONG)
	if !compareApproxTimes(sunrise, expectedSunrise) {
		t.Error("Unexpected sunrise result:", sunrise)
	}

	// Test 1ns after sunrise.
	now = expectedSunrise.Add(time.Nanosecond)
	sunrise = CalcSunrise(now, TEST_LAT, TEST_LONG)
	if !compareApproxTimes(sunrise, expectedSunrise) {
		t.Error("Unexpected sunrise result:", sunrise)
	}

	// Test well after sunrise.
	now = time.Date(2014, time.June, 12, 10, 57, 12, 0, time.Local)
	sunrise = CalcSunrise(now, TEST_LAT, TEST_LONG)
	if !compareApproxTimes(sunrise, expectedSunrise) {
		t.Error("Unexpected sunrise result:", sunrise)
	}

	// Test late in day.
	now = time.Date(2014, time.June, 12, 22, 57, 12, 0, time.Local)
	sunrise = CalcSunrise(now, TEST_LAT, TEST_LONG)
	if !compareApproxTimes(sunrise, expectedSunrise) {
		t.Error("Unexpected sunrise result:", sunrise)
	}
}

func TestNextSunrise(t *testing.T) {

	expectedSunriseToday := time.Date(2014, time.June, 12, 05, 47, 00, 0, time.Local)
	expectedSunriseTomorrow := time.Date(2014, time.June, 13, 05, 46, 58, 0, time.Local)

	// Test well before sunrise.
	now := time.Date(2014, time.June, 12, 2, 57, 12, 0, time.Local)
	sunrise := NextSunrise(now, TEST_LAT, TEST_LONG)
	if !compareApproxTimes(sunrise, expectedSunriseToday) {
		t.Error("Unexpected sunrise result:", sunrise)
	}

	// Test 1ns before sunrise.
	now = expectedSunriseToday.Add(-1 * time.Nanosecond)
	sunrise = NextSunrise(now, TEST_LAT, TEST_LONG)
	if !compareApproxTimes(sunrise, expectedSunriseToday) {
		t.Error("Unexpected sunrise result:", sunrise)
	}

	// Test 1ns after sunrise.
	now = expectedSunriseToday.Add(time.Nanosecond)
	sunrise = NextSunrise(now, TEST_LAT, TEST_LONG)
	if !compareApproxTimes(sunrise, expectedSunriseTomorrow) {
		t.Error("Unexpected sunrise result:", sunrise)
	}

	// Test well after sunrise.
	now = time.Date(2014, time.June, 12, 10, 57, 12, 0, time.Local)
	sunrise = NextSunrise(now, TEST_LAT, TEST_LONG)
	if !compareApproxTimes(sunrise, expectedSunriseTomorrow) {
		t.Error("Unexpected sunrise result:", sunrise)
	}

	// Test late in day.
	now = time.Date(2014, time.June, 12, 22, 57, 12, 0, time.Local)
	sunrise = NextSunrise(now, TEST_LAT, TEST_LONG)
	if !compareApproxTimes(sunrise, expectedSunriseTomorrow) {
		t.Error("Unexpected sunrise result:", sunrise)
	}
}

func TestCalcSunset(t *testing.T) {

	expectedSunset := time.Date(2014, time.June, 12, 20, 29, 33, 0, time.Local)

	// Test well before sunset.
	now := time.Date(2014, time.June, 12, 2, 57, 12, 0, time.Local)
	sunset := CalcSunset(now, TEST_LAT, TEST_LONG)
	if !compareApproxTimes(sunset, expectedSunset) {
		t.Error("Unexpected sunset result:", sunset)
	}

	// Test 1ns after sunset.
	now = expectedSunset.Add(time.Nanosecond)
	sunset = CalcSunset(now, TEST_LAT, TEST_LONG)
	if !compareApproxTimes(sunset, expectedSunset) {
		t.Error("Unexpected sunset result:", sunset)
	}

	// Test well after sunset.
	now = time.Date(2014, time.June, 12, 10, 57, 12, 0, time.Local)
	sunset = CalcSunset(now, TEST_LAT, TEST_LONG)
	if !compareApproxTimes(sunset, expectedSunset) {
		t.Error("Unexpected sunset result:", sunset)
	}

	// Test late in day.
	now = time.Date(2014, time.June, 12, 22, 57, 12, 0, time.Local)
	sunset = CalcSunset(now, TEST_LAT, TEST_LONG)
	if !compareApproxTimes(sunset, expectedSunset) {
		t.Error("Unexpected sunset result:", sunset)
	}
}

func TestNextSunset(t *testing.T) {

	expectedSunsetToday := time.Date(2014, time.June, 12, 20, 29, 42, 0, time.Local)
	expectedSunsetTomorrow := time.Date(2014, time.June, 13, 20, 30, 18, 0, time.Local)

	// Test well before sunset.
	now := time.Date(2014, time.June, 12, 2, 57, 12, 0, time.Local)
	sunset := NextSunset(now, TEST_LAT, TEST_LONG)
	if !compareApproxTimes(sunset, expectedSunsetToday) {
		t.Error("Unexpected sunset result:", sunset)
	}

	// Test 1ns before sunset.
	now = expectedSunsetToday.Add(-1 * time.Nanosecond)
	sunset = NextSunset(now, TEST_LAT, TEST_LONG)
	if !compareApproxTimes(sunset, expectedSunsetToday) {
		t.Error("Unexpected sunset result:", sunset)
	}

	// Test 1ns after sunset.
	now = expectedSunsetToday.Add(time.Nanosecond)
	sunset = NextSunset(now, TEST_LAT, TEST_LONG)
	if !compareApproxTimes(sunset, expectedSunsetTomorrow) {
		t.Error("Unexpected sunset result:", sunset)
	}

	// Test well after sunset.
	now = time.Date(2014, time.June, 12, 22, 57, 12, 0, time.Local)
	sunset = NextSunset(now, TEST_LAT, TEST_LONG)
	if !compareApproxTimes(sunset, expectedSunsetTomorrow) {
		t.Error("Unexpected sunset result:", sunset)
	}
}
