package parser

import (
	"fmt"
	"testing"
	"time"

	"github.com/agungdwiprasetyo/go-utils/debug"
)

func TestParseDateFormat(t *testing.T) {
	type args struct {
		date string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Testcase #1",
			args{
				date: "2017-10-11 13:00:54",
			},
			"11 October 2017",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseDateFormat(tt.args.date); got != tt.want {
				str := debug.StringRed(fmt.Sprintf("ParseDateFormat() = %v, want %v", got, tt.want))
				t.Errorf(str)
			}
		})
	}
}

func TestParseDateString(t *testing.T) {
	now, _ := ParseTime("2018-08-10")
	type args struct {
		date time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Testcase #1",
			args{
				date: now,
			},
			"10-08-2018",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseDateString(tt.args.date); got != tt.want {
				str := debug.StringRed(fmt.Sprintf("ParseDateString(%v) = %v, want %v", tt.args.date, got, tt.want))
				t.Errorf(str)
			}
		})
	}
}
