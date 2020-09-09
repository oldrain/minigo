# Minigo
A fast, tiny, api framework for go.

It is heavily influenced by Gin.
## Overview
* Router, group router
* Request context
* Handlers
* Just support json input and output
* Just suport POST
* Validate
## Installation
> step 1.

    go get -u github.com/oldrain/minigo

> step 2.

    import "github.com/oldrain/minigo"

## Quick start
> File:

    example.go

> Code:

```go
package main

import (
    "github.com/oldrain/minigo"
)

func main() {
    api := minigo.Default()

    api.Post("/hey", func(ctx *minigo.Context) {
        ctx.JSON("Hi~")
    })

    _ = api.Run(":9527")
}
```

> Request:

    curl -X POST -H "Content-Type: application/json" http://127.0.0.1:9527/hey

> Response:

    {"code":200,"msg":"Success","data":"Hi~","ts":"2019-01-01 00:00:00"}

## Examples
### Using Router and POST
```go
func main() {
    api := minigo.Default()

    // user
    user := api.Group("/user")
    user.Post("/info", func(ctx *minigo.Context) {
        ctx.JSON("user.info")
    })

    // user.favorite
    favorite := user.Group("/favorite")
    favorite.Post("/list", func(ctx *minigo.Context) {
        ctx.JSON("user.favorite.list")
    })

     _ = api.Run(":9527")
}
```
### Using handlers
```go
func main() {
    api := minigo.Default()

    // user
    user := api.Group("/user", before)
    user.Post("/info", userInfo)

    // user.favorite
    favorite := user.Group("/favorite", before)
    favorite.Post("/list", userFavoriteList)
    favorite.LastUse(after)

     _ = api.Run(":9527")
}

func before(ctx *minigo.Context) {
    // do something
}

func after(ctx *minigo.Context) {
    // logging
}

func userInfo(ctx *minigo.Context) {
    ctx.JSON("user.info")
}

func userFavoriteList(ctx *minigo.Context) {
    ctx.JSON("user.favorite.list")
}

```
### Input && Validate && Output
```go
type UserInfoIn struct {
	Id int `json:"id" regexp:"^[2-4]$" tips:"Invalid id"`
	Name string `json:"name" regexp:"^.{2,8}$" tips:"Invalid name"`
	Email string `json:"email" regexp:"^.+@.+\\..+$" tips:"Invalid email"`
}

type UserInfoOut struct {
	User *UserInfoIn `json:"user"`
}

func main() {
	api := minigo.Default()

	// user
	user := api.Group("/user")
	user.Post("/info", userInfo)

	_ = api.Run(":9527")
}

func userInfo(ctx *minigo.Context) {
	userInfoIn := new(UserInfoIn)

	_ = ctx.BindJSON(&userInfoIn)

	validate := minigo.NewValidate()
	err := validate.Do(userInfoIn)

	if err != nil {
		ctx.Error(10000, err.Error())
		return
	}

	userInfoOut := new(UserInfoOut)
	userInfoOut.User = userInfoIn

	ctx.JSON(userInfoOut)
}
```
### Others
#### Customization handler
```go
func before(ctx *minigo.Context) {
    // do something
    // if ok ctx.Continue()
    // if err ctx.Abort() or ctx.AbortWithError()
}
```
#### Request header
```go
func userInfo(ctx *minigo.Context) {
    token := ctx.GetInHeader("token")
    ...
}
```
#### Context get and set
```go
func before(ctx *minigo.Context) {
    ctx.set("key", "value")
    ...
}

func userInfo(ctx *minigo.Context) {
    value := ctx.getString("key")
    ...
}
```
## Testing
> File:

    example.go

> CMD:

    hey -n 100000 -c 10 -m POST -H "Content-Type: application/json" http://127.0.0.1:9527/hey

> Result:

```
Summary:
  Total:	4.8541 secs
  Slowest:	0.0128 secs
  Fastest:	0.0001 secs
  Average:	0.0005 secs
  Requests/sec:	20601.1037

  Total data:	6800000 bytes
  Size/request:	68 bytes

Response time histogram:
  0.000 [1]	|
  0.001 [98798]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.003 [1063]	|
  0.004 [102]	|
  0.005 [21]	|
  0.006 [7]	|
  0.008 [5]	|
  0.009 [1]	|
  0.010 [0]	|
  0.012 [1]	|
  0.013 [1]	|


Latency distribution:
  10% in 0.0002 secs
  25% in 0.0003 secs
  50% in 0.0004 secs
  75% in 0.0006 secs
  90% in 0.0008 secs
  95% in 0.0009 secs
  99% in 0.0015 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.0128 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0000 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0044 secs
  resp wait:	0.0004 secs, 0.0001 secs, 0.0128 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0059 secs

Status code distribution:
  [200]	100000 responses
```
