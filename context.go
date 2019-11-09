// Copyright 2019 minigo Author. All Rights Reserved.
// License that can be found in the LICENSE file.

package minigo

import (
	"encoding/json"
	"net"
	"strings"
	"bytes"
)

type Context struct {
	Input  Input
	Output Output

	KV map[string]interface{}

	handlers  HandlerFuncChain
	handleIdx int8
}

func (ctx *Context) reset() {
	ctx.KV = nil
	ctx.handlers = nil
	ctx.handleIdx = -1
}

func (ctx *Context) Continue() {
	ctx.handleIdx++
	count := int8(len(ctx.handlers))
	for ; ctx.handleIdx < count; ctx.handleIdx++ {
		ctx.handlers[ctx.handleIdx](ctx)
	}
}

func (ctx *Context) Abort() {
	ctx.handleIdx = maxHandleSize - 1
}

func (ctx *Context) AbortWithError(code int, msg string) {
	ctx.Error(code, msg)
	ctx.Abort()
}

// input
func (ctx *Context) BindJSON(obj interface{}) (err error) {
	decoder := json.NewDecoder(bytes.NewBufferString(ctx.GetInBody()))
	err = decoder.Decode(obj)
	return
}

func (ctx *Context) GetInHeader(key string) string {
	return ctx.Input.GetHeader(key)
}

func (ctx *Context) SetInHeader(key, value string) {
	ctx.Input.SetHeader(key, value)
}

func (ctx *Context) GetInBody() string {
	return ctx.Input.GetBody()
}

func (ctx *Context) Method() string {
	return ctx.Input.Request.Method
}

func (ctx *Context) Path() string {
	return dealSlash(ctx.Input.Request.URL.Path)
}

func (ctx *Context) ClientIP() string {
	ip := ctx.GetInHeader("X-Forwarded-For")
	if idx := strings.IndexByte(ip, ','); idx >= 0 {
		ip = ip[0:idx]
	}

	ip = strings.TrimSpace(ip)
	if len(ip) > 0 {
		return ip
	}

	ip = strings.TrimSpace(ctx.GetInHeader("X-Real-Ip"))
	if len(ip) > 0 {
		return ip
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(ctx.Input.Request.RemoteAddr)); err == nil {
		return ip
	}

	return ""
}

// end

// output
func (ctx *Context) JSON(data interface{}) {
	oJson := OJson{
		Code: codeOk,
		Msg:  msgOk,
		Data: data,
		Ts:   timeStr(),
	}
	ctx.setApiStatus(codeOk)
	err := ctx.Output.RenderJson(oJson)
	if err != nil {
		panic(err)
	}
}

func (ctx *Context) Error(code int, msg string) {
	oJson := OJson{
		Code: code,
		Msg:  msg,
		Data: nil,
		Ts:   timeStr(),
	}
	ctx.setApiStatus(code)
	err := ctx.Output.RenderJson(oJson)
	if err != nil {
		panic(err)
	}
}

func (ctx *Context) GetOutHeader(key string) string {
	return ctx.Output.GetHeader(key)
}

func (ctx *Context) SetOutHeader(key, value string) {
	ctx.Output.SetHeader(key, value)
}

func (ctx *Context) GetOutBody() string {
	return ctx.Output.getBody()
}

func (ctx *Context) SetOutBody(body string) {
	ctx.Output.setBody(body)
}

func (ctx *Context) getApiStatus() int {
	return ctx.Output.getApiStatus()
}

func (ctx *Context) setApiStatus(status int) {
	ctx.Output.setApiStatus(status)
}

// end

// KV
func (ctx *Context) Set(key string, value interface{}) {
	if ctx.KV == nil {
		ctx.KV = make(map[string]interface{})
	}
	ctx.KV[key] = value
}

func (ctx *Context) Get(key string) (value interface{}, ok bool) {
	value, ok = ctx.KV[key]
	return
}

func (ctx *Context) GetString(key string) (s string) {
	if value, ok := ctx.Get(key); ok && value != nil {
		s, _ = value.(string)
	}
	return
}

func (ctx *Context) GetBool(key string) (b bool) {
	if value, ok := ctx.Get(key); ok && value != nil {
		b, _ = value.(bool)
	}
	return
}

func (ctx *Context) GetInt(key string) (i int) {
	if value, ok := ctx.Get(key); ok && value != nil {
		i, _ = value.(int)
	}
	return
}

func (ctx *Context) GetInt32(key string) (i int32) {
	if value, ok := ctx.Get(key); ok && value != nil {
		i, _ = value.(int32)
	}
	return
}

func (ctx *Context) GetInt64(key string) (i int64) {
	if value, ok := ctx.Get(key); ok && value != nil {
		i, _ = value.(int64)
	}
	return
}

func (ctx *Context) GetUint(key string) (i uint) {
	if value, ok := ctx.Get(key); ok && value != nil {
		i, _ = value.(uint)
	}
	return
}

func (ctx *Context) GetUint32(key string) (i uint32) {
	if value, ok := ctx.Get(key); ok && value != nil {
		i, _ = value.(uint32)
	}
	return
}

func (ctx *Context) GetUint64(key string) (i uint64) {
	if value, ok := ctx.Get(key); ok && value != nil {
		i, _ = value.(uint64)
	}
	return
}

func (ctx *Context) GetFloat32(key string) (i float32) {
	if value, ok := ctx.Get(key); ok && value != nil {
		i, _ = value.(float32)
	}
	return
}

func (ctx *Context) GetFloat64(key string) (i float64) {
	if value, ok := ctx.Get(key); ok && value != nil {
		i, _ = value.(float64)
	}
	return
}

// end
