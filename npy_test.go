package npy

import (
	"testing"
	"math"
)

func TestNpyDense(t *testing.T) {
	r, c, data, err := Read("data/dense.npy")
	if err != nil {
		t.Errorf("Error %v reading dense.py", err)
	}
	if r != 100 || c != 100 || math.Floor(data[0]*1e24) != 1614457 {
		t.Errorf("Error reading dense.py")
	}
}
