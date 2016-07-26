package models

import (
	"fmt"
	"testing"
)

func TestPayState(t *testing.T) {
	s := fmt.Sprintf("%s", PayState(0))
	if s != payStateDesc[0] {
		t.Errorf("%s not equal to %s", s, payStateDesc[0])
	}
}
