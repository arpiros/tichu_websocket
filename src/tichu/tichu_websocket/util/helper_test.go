package util

import (
	"testing"
)

func TestRandomRange(t *testing.T) {
	for i := 0; i < 10; i++ {
		getValue := RandomRange(0, 13)
		println(getValue)
	}
}
