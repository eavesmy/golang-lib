package compress

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"
)

func UnGzip(data []byte) []byte {
	buf := bytes.NewBuffer(data)
	r, err := gzip.NewReader(buf)

	if err != nil {
		fmt.Println(err, 123)
		return []byte{}
	}

	ret, _ := ioutil.ReadAll(r)

	return ret
}
