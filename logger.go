// Copyright 2019 minigo Author. All Rights Reserved.
// License that can be found in the LICENSE file.

package minigo

import (
	"fmt"
	"time"
)

func Logger() HandlerFunc {
	return func(ctx *Context) {
		start := time.Now()

		ctx.Continue()

		end := time.Now()
		path := ctx.Path()
		latency := end.Sub(start)
		clientIP := ctx.ClientIP()
		method := ctx.Method()
		status := ctx.getApiStatus()

		content := fmt.Sprintf("[MINI] %v | %s | %13v | %15s | %s | %s",
			time2Str(end), colorForStatus(status), latency, clientIP, colorForMethod(method), path)

		fmt.Println(content)
	}
}

func colorForStatus(code int) string {
	switch {
	case code >= 200 && code < 300:
		return fmt.Sprintf("\033[32m%d\033[0m", code)
	case code >= 300 && code < 400:
		return fmt.Sprintf("\033[37m%d\033[0m", code)
	case code >= 400 && code < 500:
		return fmt.Sprintf("\033[33m%d\033[0m", code)
	default:
		return fmt.Sprintf("\033[31m%d\033[0m", code)
	}
}

func colorForMethod(method string) string {
	return fmt.Sprintf("\033[34m%s\033[0m", method)
}

func logPlain(args ...interface{}) {
	fmt.Println(fmt.Sprintf("[%s][Application]%s", timeStr(), fmt.Sprint(args...)))
}

func logInfo(s string, args ...interface{}) {
	fmt.Println(fmt.Sprintf("[%s][\033[32m[%s]\033[0m]", timeStr(), fmt.Sprint(args...)))
}

func logRequest(s string, args ...interface{}) {
	fmt.Println(fmt.Sprintf("[%s][\033[32m[Application Error]\033[0m][%s]%s", timeStr(), s, fmt.Sprint(args...)))
}

func logPrintRoute(httpMethod, absPath string, handlers HandlerFuncChain) {
	handlerNum := len(handlers)
	handlerName := nameOfFunc(handlers[handlerNum-1])
	fmt.Println(fmt.Sprintf("%-6s %-25s --> %s (%d handlers)", httpMethod, absPath, handlerName, handlerNum))
}
func logPanic(args ...interface{}) {
	fmt.Println(fmt.Sprintf("[%s][\033[31m[Application Panic]\033[0m]%s", timeStr(), fmt.Sprint(args...)))
}
