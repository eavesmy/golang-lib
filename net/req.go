package net

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func NewRequest(method, path string, buf []byte) (res *http.Response, err error) {

	var buffer io.Reader
	if buf != nil {
		buffer = bytes.NewBuffer(buf)
	}

	req, err := http.NewRequest(method, path, buffer)
	c := &http.Client{}
	res, err = c.Do(req)

	if err != nil {
		fmt.Println("request error: ", err)
		return
	}

	defer res.Body.Close()

	return
}
