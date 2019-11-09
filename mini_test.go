package minigo

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"
)

type tUserInfoIn struct {
	Id int `json:"id" regexp:"^[2-4]$" tips:"Invalid id"`
	Name string `json:"name" regexp:"^.{2,8}$" tips:"Invalid name"`
	Email string `json:"email" regexp:"^.+@.+\\..+$" tips:"Invalid email"`
}

type tDataOut struct {
	Path string `json:"path"`
	Before string `json:"before"`
	After string `json:"after"`
}

func getTestApi() *Api {
	tApi := Default()

	// user
	tUser := tApi.Group("/user", tBefore)
	tUser.Post("/info", tUserInfo)

	// user.favorite
	tFavorite := tUser.Group("/favorite")
	tFavorite.Post("/list", tUserFavoriteList)

	// order
	tOrder := tApi.Group("/order")
	tOrder.Use(tBefore)
	tOrder.Post("/list", tOrderList)
	tOrder.Post("/detail", tOrderDetail)
	tOrder.LastUse(tAfter)

	return tApi
}

func newTestContext() *Context {
	return Default().newContext()
}

func getTestContext(api *Api) *Context {
	return api.newContext()
}

func newTestServer() *httptest.Server {
	return httptest.NewServer(getTestApi())
}

func getTestServer(api *Api) *httptest.Server {
	return httptest.NewServer(api)
}

func getTestUrl(url, path string) string {
	return fmt.Sprintf("%s%s", url, path)
}

func tBefore(ctx *Context) {
	ctx.Set("before", "append before")
}

func tAfter(ctx *Context) {
	ctx.Set("after", "append after")
	body := ctx.GetOutBody()
	dataOut := new(tDataOut)
	_ = json.Unmarshal([]byte(body), dataOut)
	dataOut.After = ctx.GetString("after")
	ctx.JSON(dataOut)
}

func tUserInfo(ctx *Context) {
	userInfo := new(tUserInfoIn)

	_ = ctx.BindJSON(&userInfo)

	validate := NewValidate()
	err := validate.Do(userInfo)

	if err != nil {
		ctx.Error(10000, err.Error())
		return
	}

	ctx.JSON(userInfo)
}

func tUserFavoriteList(ctx *Context) {
	rtn := tDataOut{
		Path: "user.favorite.list",
		Before: ctx.GetString("before"),
		After: ctx.GetString("after"),
	}
	ctx.JSON(rtn)
}

func tOrderList(ctx *Context) {
	rtn := tDataOut{
		Path: "order.list",
		Before: ctx.GetString("before"),
		After: ctx.GetString("after"),
	}
	ctx.SetOutBody(toJson(rtn))
}

func tOrderDetail(ctx *Context) {
	rtn := tDataOut{
		Path: "order.detail",
		Before: ctx.GetString("before"),
		After: ctx.GetString("after"),
	}
	ctx.SetOutBody(toJson(rtn))
}

func getTUserInfo() *tUserInfoIn {
	return &tUserInfoIn{
		Id: 3,
		Name: "you",
		Email: "you@gmail.com",
	}
}

func AssertEqual(t *testing.T, expected, actual interface{}) {
	if expected != actual {
		t.Error(fmt.Sprintf("expected: %v, actual: %v", expected, actual))
	}
}

func AssertError(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}
