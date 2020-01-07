package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTryCatch(t *testing.T) {
	t.Run("Test Catch Panic", func(t *testing.T) {
		TryCatch(
			func() {
				panic("panic!!")
			},
			func(err error) {
				assert.NotNil(t, err)
				assert.Equal(t, err.Error(), "panic!!")
			},
		)
	})
	t.Run("Test Catch Panic Nil Pointer", func(t *testing.T) {
		TryCatch(
			func() {
				var a *struct {
					s string
				}
				fmt.Println(a.s)
			},
			func(err error) {
				assert.NotNil(t, err)
				assert.Contains(t, err.Error(), "invalid memory address or nil pointer dereference")
			},
		)
	})
	t.Run("Test Catch Panic index out of range", func(t *testing.T) {
		TryCatch(
			func() {
				var a []string
				fmt.Println(a[10])
			},
			func(err error) {
				assert.NotNil(t, err)
				assert.Contains(t, err.Error(), "index out of range")
			},
		)
	})
	t.Run("Test Catch Panic interface conversion", func(t *testing.T) {
		TryCatch(
			func() {
				var a interface{}
				a = 10
				fmt.Println(a.(string))
			},
			func(err error) {
				assert.NotNil(t, err)
				assert.Contains(t, err.Error(), "interface conversion")
			},
		)
	})
}
