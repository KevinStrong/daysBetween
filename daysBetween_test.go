package daysbetween_test

import (
	"testing"
	"time"

	"github.com/KevinStrong/daysbetween"

	"github.com/matryer/is"
)

func Test_Get_WithNoValues_ReturnsDaysBetweenUnixEpocAndToday(t *testing.T) {
	expect := is.New(t)
	defaultStart := time.Unix(0, 0)
	defaultEnd := time.Now()
	want := CalculateRange(defaultStart, defaultEnd)

	got := daysbetween.Get()

	expect.Equal(got, want)
}

func Test_daysBetweenCustomStartAndToday(t *testing.T) {
	expect := is.New(t)

	customStart := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
	defaultEnd := time.Now()
	want := CalculateRange(customStart, defaultEnd)

	got := daysbetween.Get(daysbetween.SetStart(customStart))

	expect.Equal(got, want)
}

func Test_daysBetween_WithACustomRange_ReturnsDaysBetweenTheRange(t *testing.T) {
	expect := is.New(t)

	customStart := time.Date(2019, 5, 1, 0, 0, 0, 0, time.UTC)
	customEnd := time.Date(2019, 6, 1, 0, 0, 0, 0, time.UTC)
	want := 31

	got := daysbetween.Get(daysbetween.SetStart(customStart), daysbetween.SetEnd(customEnd))

	expect.Equal(got, want)
}

func CalculateRange(start time.Time, end time.Time) int {
	// 	Calculate the expected result.  This may fail if run exactly at midnight
	want := int(end.Sub(start).Hours()) / 24
	return want
}
