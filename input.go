// Copyright 2019 minigo Author. All Rights Reserved.
// License that can be found in the LICENSE file.

package minigo

import (
	"io/ioutil"
	"net/http"
)

type Input struct {
	Request *http.Request
}

func (in *Input) GetHeader(key string) string {
	return in.Request.Header.Get(key)
}

func (in *Input) SetHeader(key, value string) {
	in.Request.Header.Set(key, value)
}

func (in *Input) GetBody() string {
	body, _ := ioutil.ReadAll(in.Request.Body)
	return string(body)
}
