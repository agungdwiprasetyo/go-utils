package utils

import (
	"fmt"
	"testing"

	"github.com/agungdwiprasetyo/go-utils/debug"
)

func TestParseDateString(t *testing.T) {
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
			if got := ParseDateString(tt.args.date); got != tt.want {
				str := debug.StringRed(fmt.Sprintf("ParseDateString() = %v, want %v", got, tt.want))
				t.Errorf(str)
			}
		})
	}
}
