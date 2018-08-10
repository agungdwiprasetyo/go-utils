package utils

import (
	"fmt"
	"time"
)

// Timer is struct for calculate performance
type Timer struct {
	Desc      string
	StartTime time.Time
}

// NewTimer for create new object timer
func NewTimer(desc string) *Timer {
	return &Timer{Desc: desc, StartTime: time.Now()}
}

func (timer *Timer) Elapsed() time.Duration {
	return time.Since(timer.StartTime)
}

func (timer *Timer) Print() {
	elapsed := time.Since(timer.StartTime)
	fmt.Printf("%s %s: %v\n", "Time elapsed", timer.Desc, elapsed)
}
