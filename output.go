// Copyright 2019 minigo Author. All Rights Reserved.
// License that can be found in the LICENSE file.

package minigo

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Output struct {
	Response http.ResponseWriter

	// body
	body *bytes.Buffer

	// Api status
	apiStatus int
}

type OJson struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
	Ts   string      `json:"ts"`
}

func (out *Output) RenderJson(oJson OJson) error {
	out.writeContentType(outputContentTypeJson)
	jsonBytes, err := json.Marshal(oJson)
	if err != nil {
		return err
	}
	out.setBody(string(jsonBytes))
	_, err = out.Response.Write(jsonBytes)
	return err
}

func (out *Output) GetHeader(key string) string {
	return out.Response.Header().Get(key)
}

func (out *Output) SetHeader(key, value string) {
	out.Response.Header().Set(key, value)
}

func (out *Output) getBody() string {
	return out.body.String()
}

func (out *Output) setBody(body string) {
	out.body = bytes.NewBufferString(body)
}

func (out *Output) getApiStatus() int {
	return out.apiStatus
}

func (out *Output) setApiStatus(status int) {
	out.apiStatus = status
}

func (out *Output) writeContentType(value string) {
	out.Response.Header().Set(contentTypeHeader, value)
}
