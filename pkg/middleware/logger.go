package middleware

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"k8s.io/klog/v2"
	"reflect"
	"runtime"
)

var (
	skipPaths = []string{
		"/health",
	}

	skipHandlerFncsMap map[string]struct{}
)

//AddLogSkipPaths 用于添加自定义不打印日志的接口
func AddLogSkipPaths(paths ...string) {
	skipPaths = append(skipPaths, paths...)
}

//AddLogSkipHandlerFuncs 用于添加自定义不打印日志的gin.HandlerFunc
func AddLogSkipHandlerFuncs(fncs ...gin.HandlerFunc) {
	for _, fnc := range fncs {
		name := runtime.FuncForPC(reflect.ValueOf(fnc).Pointer()).Name()
		skipHandlerFncsMap[name] = struct{}{}
	}
}

func getSkipPaths() []string {
	return skipPaths
}

func getSkipHandlerFncsMap() map[string]struct{} {
	return skipHandlerFncsMap
}

//Logger 用于gin请求调用时，输出请求体，响应体的日志中间件
func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			requestBody          []byte
			bodyStr              string
			buf                  = new(bytes.Buffer)
			customResponseWriter = &customResponseWriter{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
		)

		if ctx.Request.Body != nil {
			requestBody, _ = ioutil.ReadAll(ctx.Request.Body)
			bodyStr = string(requestBody)
		}

		ctx.Request.Body = ioutil.NopCloser(bytes.NewReader(requestBody))
		ctx.Writer = customResponseWriter

		//强制打印颜色（默认Output为终端输出时才会打印）
		gin.ForceConsoleColor()
		gin.LoggerWithConfig(gin.LoggerConfig{
			Formatter: nil,
			Output:    buf,
			SkipPaths: getSkipPaths(),
		})(ctx)

		log := buf.String()
		if log == "" {
			return
		}

		handlerFncsMap := getSkipHandlerFncsMap()
		if len(handlerFncsMap) > 0 {
			name := ctx.HandlerName()
			if _, ok := handlerFncsMap[name]; ok {
				fmt.Print(log)
				return
			}
		}

		if ctx.Request.Header != nil {
			if klog.V(4).Enabled() {
				log = fmt.Sprintf("%sRequest Header: %v\n", log, ctx.Request.Header)
			}
		}

		if bodyStr != "" {
			if klog.V(2).Enabled() {
				log = fmt.Sprintf("%sRequest Body: %s\n", log, bodyStr)
			}
		}

		resHeader := ctx.Writer.Header()
		if resHeader != nil {
			if klog.V(5).Enabled() {
				log = fmt.Sprintf("%sResponse Header: %v\n", log, resHeader)
			}
		}

		responseStr := customResponseWriter.body.String()
		if responseStr != "" {
			if klog.V(3).Enabled() {
				log = fmt.Sprintf("%sResponse Body: %s\n", log, responseStr)
			}
		}
		fmt.Print(log)
	}
}

type customResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w customResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w customResponseWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}
