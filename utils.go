// Copyright 2019 minigo Author. All Rights Reserved.
// License that can be found in the LICENSE file.

package minigo

import (
	"encoding/json"
	"strings"
	"time"
	"runtime"
	"reflect"
)

func dealSlash(path string) string {
	size := len(path)

	if size == 0 {
		return path
	}

	if path[0] != '/' {
		path = "/" + path
	}

	if path[size - 1] == '/' {
		path = strings.TrimRight(path, "/")
	}

	return path
}

func joinPath(basePath, path string) string {
	return dealSlash(basePath) + dealSlash(path)
}

func nameOfFunc(f interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}

func timeStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func time2Str(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// just for test
func toJson(s interface{}) string {
	jsons, errs := json.Marshal(s)
	if errs != nil {
		return ""
	}
	return string(jsons)
}
