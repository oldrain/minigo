// Copyright 2019 minigo Author. All Rights Reserved.
// License that can be found in the LICENSE file.

package minigo

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestContextBindJSON(t *testing.T) {
	ctx := newTestContext()

	inJson := `{"id":3,"name":"you","email":"you@gmail.com"}`

	ctx.Input.Request, _ = http.NewRequest(reqMethodPost, "/", bytes.NewBufferString(inJson))
	ctx.Input.SetHeader(contentTypeHeader, inputContentTypeJson)

	userInfo := new(tUserInfoIn)

	AssertError(t, ctx.BindJSON(&userInfo))
}

func TestContextJSON(t *testing.T) {
	ctx := newTestContext()

	ctx.Output.Response = httptest.NewRecorder()

	ctx.JSON(getTUserInfo())
}

func TestContextHeader(t *testing.T) {
	ctx := newTestContext()

	ctx.Input.Request, _ = http.NewRequest(reqMethodPost, "/", nil)
	ctx.SetInHeader(contentTypeHeader, inputContentTypeJson)

	AssertEqual(t, inputContentTypeJson, ctx.GetInHeader(contentTypeHeader))

	var tokenK = "token"
	var tokenV = "you token"

	ctx.Output.Response = httptest.NewRecorder()
	ctx.SetOutHeader(tokenK, tokenV)
	AssertEqual(t, tokenV, ctx.GetOutHeader(tokenK))
}

func TestContextGetSet(t *testing.T) {
	ctx := newTestContext()

	var k = "foo"
	var v = "bar"
	ctx.Set(k, v)

	value, err := ctx.Get(k)
	AssertEqual(t, true, err)
	AssertEqual(t, v, value)

	ctx.Set("string", "this is a string")
	ctx.Set("int", 22)
	ctx.Set("int32", int32(-22))
	ctx.Set("int64", int64(2222222222222222))
	ctx.Set("uint", uint(22))
	ctx.Set("uint32", uint32(22))
	ctx.Set("uint64", uint64(22))
	ctx.Set("float32", float32(2.2))
	ctx.Set("float64", 2.2)
	var a interface{} = 1
	ctx.Set("intInterface", a)

	b, _ := ctx.Get("intInterface")
	AssertEqual(t, b.(int), 1)

	AssertEqual(t, ctx.GetString("string"), "this is a string")
	AssertEqual(t, ctx.GetInt32("int32"), int32(-22))
	AssertEqual(t, ctx.GetInt64("int64"), int64(2222222222222222))
	AssertEqual(t, ctx.GetUint("uint"), uint(22))
	AssertEqual(t, ctx.GetUint32("uint32"), uint32(22))
	AssertEqual(t, ctx.GetUint64("uint64"), uint64(22))
	AssertEqual(t, ctx.GetFloat32("float32"), float32(2.2))
	AssertEqual(t, ctx.GetFloat64("float64"), 2.2)
}
