package daysbetween_test

import "testing"
import "daysbetween"

func Test_daysAfterEpoc(t *testing.T) {
	count := daysbetween.DaysAfterEpoc()
	// todo this test only works on the day I made it
	if count != 19119 {
		t.Fail()
	}
}
