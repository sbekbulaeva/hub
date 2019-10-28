package util

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"
)

func IsGzipData(data []byte) bool {
	return len(data) > 2 &&
		data[0] == '\x1f' && data[1] == '\x8b'
}

func Gunzip(compressed []byte) ([]byte, error) {
	gunzip, err := gzip.NewReader(bytes.NewBuffer(compressed))
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(gunzip)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func Gzip(data []byte) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	z := gzip.NewWriter(buf)
	wrote, err := z.Write(data)
	err2 := z.Close()
	if err != nil || wrote != len(data) || err2 != nil {
		if err == nil && err2 != nil {
			err = err2
		}
		return nil, fmt.Errorf("%v; wrote %d of %d bytes", err, wrote, len(data))
	}
	return buf.Bytes(), nil
}
