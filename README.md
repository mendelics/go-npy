# npy
Read in .npy files produced by numpy into Go

###Usage
```
import (
  "github.com/cquotient/npy"
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

### Build Status
[![Build Status](https://travis-ci.org/cquotient/npy.svg?branch=master)](https://travis-ci.org/cquotient/npy)