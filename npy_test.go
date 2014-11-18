package npy

import (
	"testing"
)

func TestNpyDense(t *testing.T) {
	r, c, data, err := Read("data/dense.npy")
	if err != nil {
		t.Errorf("Error %v reading dense.py", err)
	}
	if r != 100 || c != 100 || data[0] != 1.61445790e-18 {
		t.Errorf("Error reading dense.py")
	}
}
