package daysbetween

import "time"

type config struct {
	start time.Time
	end   time.Time
}

type Setup func(*config)

func SetStart(start time.Time) Setup {
	return func(c *config) {
		c.start = start
	}
}

func SetEnd(end time.Time) Setup {
	return func(c *config) {
		c.end = end
	}
}

func Get(setups ...Setup) int {
	today := time.Now()
	epoc := time.Unix(0, 0)
	c := config{start: epoc, end: today}

	for i := range setups {
		setups[i](&c)
	}

	return daysBetween(c)
}

func daysBetween(c config) int {
	days := int(c.end.Sub(c.start).Hours()) / 24

	return days
}
