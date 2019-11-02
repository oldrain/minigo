// Copyright 2019 minigo Author. All Rights Reserved.
// License that can be found in the LICENSE file.

package minigo

import "math"

const (
	maxHandleSize int8 = math.MaxUint8 / 2

	reqMethodPost = "POST"

	contentTypeHeader = "Content-Type"
	inputContentTypeJson = "application/json"
	outputContentTypeJson = "application/json; charset=utf-8"

	codeOk = 200
	code400 = 400
	code404 = 404
	code405 = 405
	code500 = 500

	msgOk = "Success"
	msg400 = "Request error"
	msg404 = "Not Found"
	msg500 = "Server error"
)
