package utils

import (
	"fmt"
	"time"
)

// Timer is struct for calculate performance
type Timer struct {
	Description string
	StartTime   time.Time
}

// NewTimer for create new object timer
func NewTimer(description string) *Timer {
	return &Timer{Description: description, StartTime: time.Now()}
}

func (timer *Timer) Elapsed() time.Duration {
	return time.Since(timer.StartTime)
}

func (timer *Timer) Print() {
	elapsed := time.Since(timer.StartTime)
	fmt.Printf("%s %s: %v\n", "Time elapsed", timer.Description, elapsed)
}
