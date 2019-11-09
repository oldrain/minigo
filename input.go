// Copyright 2019 minigo Author. All Rights Reserved.
// License that can be found in the LICENSE file.

package minigo

import (
	"io/ioutil"
	"net/http"
)

type Input struct {
	Request *http.Request

	// body
	body []byte
}

func (in *Input) GetHeader(key string) string {
	return in.Request.Header.Get(key)
}

func (in *Input) SetHeader(key, value string) {
	in.Request.Header.Set(key, value)
}

func (in *Input) GetBody() string {
	if len(in.body) > 0 {
		return string(in.body)
	}
	in.body, _ = ioutil.ReadAll(in.Request.Body)
	return string(in.body)
}
