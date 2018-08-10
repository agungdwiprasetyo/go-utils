package shuffle

import (
	"math/rand"
	"reflect"
	"sort"
)

type (
	Interface interface {
		Len() int
		Swap(i, j int)
	}

	Int64Slice []int64

	Shuffler rand.Rand
)

func (p Int64Slice) Len() int {
	return len(p)
}
func (p Int64Slice) Less(i, j int) bool {
	return p[i] < p[j]
}
func (p Int64Slice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func SortInt64s(a []int64) {
	sort.Sort(Int64Slice(a))
}

func Slice(slice interface{}) {
	rv := reflect.ValueOf(slice)
	swap := reflect.Swapper(slice)
	n := rv.Len()
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		swap(i, j)
	}
}

func Shuffle(data Interface) {
	n := data.Len()
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		data.Swap(i, j)
	}
}

func Ints(a []int) {
	Shuffle(sort.IntSlice(a))
}

func Int64s(a []int64) {
	Shuffle(Int64Slice(a))
}

func Float64s(a []float64) {
	Shuffle(sort.Float64Slice(a))
}

func Strings(a []string) {
	Shuffle(sort.StringSlice(a))
}

// For Shuffler method

func New(src rand.Source) *Shuffler {
	return (*Shuffler)(rand.New(src))
}

func (s *Shuffler) Slice(slice Interface) {
	rv := reflect.ValueOf(slice)
	swap := reflect.Swapper(slice)
	n := rv.Len()
	for i := n - 1; i >= 0; i-- {
		j := (*rand.Rand)(s).Intn(i + 1)
		swap(i, j)
	}
}

func (s *Shuffler) Shuffle(data Interface) {
	n := data.Len()
	for i := n - 1; i >= 0; i-- {
		j := (*rand.Rand)(s).Intn(i + 1)
		data.Swap(i, j)
	}
}

func (s *Shuffler) Ints(a []int) {
	s.Shuffle(sort.IntSlice(a))
}

func (s *Shuffler) Float64s(a []float64) {
	s.Shuffle(sort.Float64Slice(a))
}

func (s *Shuffler) Strings(a []string) {
	s.Shuffle(sort.StringSlice(a))
}
