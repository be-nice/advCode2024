package utils

import (
	"fmt"
	"time"
)

type Timer struct {
	start time.Time
}

func (t *Timer) StartTimer() {
	t.start = time.Now()
}

func (t *Timer) PrintDuration() {
	fmt.Printf("Run time: %f\n", time.Since(t.start).Seconds())
}
