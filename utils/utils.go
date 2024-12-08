package utils

import (
	"fmt"
	"time"
)

type Timer struct {
	start time.Time
	total time.Duration
}

func (t *Timer) StartTimer() {
	t.start = time.Now()
}

func (t *Timer) PrintDuration() {
	fmt.Printf("Run time: %f seconds\n", time.Since(t.start).Seconds())
	t.total += time.Since(t.start)
}

func (t *Timer) PrintTotalDuration() {
	fmt.Printf("Total run time %f seconds\n", t.total.Seconds())
}
