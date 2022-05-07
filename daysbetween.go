package daysbetween

import "time"

func DaysAfterEpoc() int {
	today := time.Now()
	epoc := time.Unix(0, 0)

	days := int(today.Sub(epoc).Hours()) / 24

	return days
}
