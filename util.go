package footballdata

import (
	"fmt"
	"time"
)

const (
	day  = 24 * time.Hour
	week = 7 * day

	timeFrameLayout = "2006-01-02"
)

func durationToTimeFrame(d time.Duration) (r string) {
	if d < 0 {
		r = "p"
		d = -d
	} else if d > 0 {
		r = "n"
	}
	r += fmt.Sprint(d.Hours() / 24)
	return
}
