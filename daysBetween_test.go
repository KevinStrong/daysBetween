package daysBetween_test

import "testing"

func Test_getDaysFromEpoc(t *testing.T) {
	count := getsDaysFromEpoc()
	// todo this test only works on the day I made it
	if count != 19119 {
		t.Fail()
	}
}
