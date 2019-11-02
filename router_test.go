// Copyright 2019 minigo Author. All Rights Reserved.
// License that can be found in the LICENSE file.

package minigo

import (
	"testing"
)

func TestRouter(t *testing.T) {
	api := getTestApi()

	AssertEqual(t, 4, len(api.routes))

	var routeCountMap = map[string]int {
		"/user/info": 4,
		"/user/favorite/list": 4,
		"/order/list": 5,
		"/order/detail": 5,
	}

	for routeName := range api.routes {
		if _, ok := routeCountMap[routeName]; !ok {
			t.Error("route no exist")
		}
		var routeInfo = api.routes[routeName]
		AssertEqual(t, routeCountMap[routeName], len(routeInfo.handlers))
	}
}
