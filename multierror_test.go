package utils

import (
	"fmt"
	"testing"
)

func TestMultiError(t *testing.T) {
	multiError := NewMultiError()

	multiError.Append("err1", fmt.Errorf("error 1"))
	if multiError.IsNil() == true {
		t.Errorf("should not nil, got: %v", multiError.IsNil())
	}
	errMap := multiError.ToMap()
	if len(errMap) != 1 {
		t.Errorf("unexpected, got: %d", len(errMap))
	}
	if multiError.Error() != "err1: error 1" {
		t.Errorf("unexpected, got: %s", multiError.Error())
	}
}
