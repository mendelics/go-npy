// Package npy provides functions to read in Python numpy data files into Go
package npy

import (
	"bufio"
	"encoding/binary"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	npyHdrLen = 6
	npyMarker = 0x93
	npyHdr    = "NUMPY"
)

func checkFormat(buf []byte) bool {
	if buf[0] != npyMarker {
		return false
	}
	for i := 1; i < npyHdrLen; i++ {
		if npyHdr[i-1] != buf[i] {
			return false
		}
	}
	return true
}

/*
Read returns the number of rows, columns and the data array corresponding to a dense float64 numpy
matrix stored in the input file
*/
func Read(fname string) (rows int, cols int, data []float64, err error) {
	fi, err := os.Open(fname)
	if err != nil {
		log.Print(err)
		return 0, 0, nil, err
	}
	defer func() {
		if err := fi.Close(); err != nil {
			log.Print(err)
		}
	}()
	r := bufio.NewReader(fi)
	magicbuf := make([]byte, npyHdrLen+4)
	_, err = r.Read(magicbuf)
	if err != nil {
		log.Print(err)
		return 0, 0, nil, err
	}
	if !checkFormat(magicbuf) {
		log.Printf("File is not an npy file %v", magicbuf)
		return 0, 0, nil, err
	}
	hdrLen, _ := binary.Uvarint(magicbuf[8:9])
	// log.Printf("File %s is an npy file of version %x.%x with hdrLen %v\n", fname, magicbuf[6], magicbuf[7], hdrLen)

	extraBytes := (npyHdrLen + 4 + hdrLen) % 16
	if extraBytes > 0 {
		extraBytes = 16 - extraBytes
	} else {
		extraBytes = 0
	}

	hdrBuf := make([]byte, hdrLen+extraBytes)
	_, err = r.Read(hdrBuf)
	// log.Printf("Read %d bytes\n", n)
	if err != nil {
		log.Print(err)
		return 0, 0, nil, err
	}
	hdrStr := strings.TrimSpace(string(hdrBuf))
	shape := strings.Split(hdrStr[strings.Index(hdrStr, "(")+1:+strings.Index(hdrStr, ")")], ",")
	rows, _ = strconv.Atoi(strings.TrimSpace(shape[0]))
	cols, _ = strconv.Atoi(strings.TrimSpace(shape[1]))

	// log.Printf("Matrix shape: %d X %d, Data size:%v bytes\n", rows, cols, rows*cols*8)

	qdata := make([]byte, 8)
	data = make([]float64, rows*cols)
	for i := int(0); i < (rows * cols); i++ {
		_, err = r.Read(qdata)
		if err != nil {
			log.Print(err)
			return 0, 0, nil, err
		}
		data[i] = math.Float64frombits(binary.LittleEndian.Uint64(qdata))
	}
	return rows, cols, data, nil
}
