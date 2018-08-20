package algorithm

import (
	"testing"

	"github.com/agungdwiprasetyo/go-utils/debug"
)

func TestSearchInSlice(t *testing.T) {
	data := []struct {
		ID   int
		Name string
	}{
		{ID: 4, Name: "ha"},
		{ID: 2, Name: "hi"},
		{ID: 100, Name: "hu"},
		{ID: 23, Name: "he"},
		{ID: 188, Name: "ho"},
		{ID: 67, Name: "ol"},
	}

	type args struct {
		slice     interface{}
		value     interface{}
		fieldName string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Test found #1",
			args{
				slice:     data,
				value:     188,
				fieldName: "ID",
			},
			4,
		},
		{
			"Test found #2",
			args{
				slice:     data,
				value:     67,
				fieldName: "ID",
			},
			5,
		},
		{
			"Test not found #3",
			args{
				slice:     data,
				value:     66,
				fieldName: "ID",
			},
			-1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SearchInSlice(tt.args.slice, tt.args.value, tt.args.fieldName)
			debug.PrintGreen("Got", got)
			if got != tt.want {
				t.Errorf("\x1b[31;1mSearchInSlice() = %v, want %v\x1b[0m", got, tt.want)
			}
		})
	}
}
