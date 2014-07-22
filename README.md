# go-npy
Read in .npy files produced by numpy into Go

[![Build Status](https://travis-ci.org/cquotient/go-npy.svg?branch=master)](https://travis-ci.org/cquotient/go-npy)
[![GoDoc](https://godoc.org/github.com/cquotient/go-npy?status.png)](https://godoc.org/github.com/cquotient/go-npy)

###Usage
```
import (
  npy "github.com/cquotient/go-npy"
  "github.com/gonum/matrix/mat64"
)
...

rows, cols, data, err := npy.NpyRead("data/dense.py")
x := mat64.NewDense(int(rows), int(cols), data)
```

###Caveats
- This is a hack, but it does work for us
- Only supports dense float64 matrices
- .npz files are not currently supported
- Assumes a single matrix in the file

###npy format documentation
https://github.com/numpy/numpy/blob/master/doc/neps/npy-format.rst
