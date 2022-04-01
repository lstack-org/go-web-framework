package app

import (
	"context"
	"fmt"
	"github.com/lstack-org/go-web-framework/pkg/code"
	"github.com/lstack-org/go-web-framework/pkg/notify"
	"github.com/lstack-org/go-web-framework/pkg/req"
	"github.com/lstack-org/go-web-framework/pkg/res"
	"runtime"
	"strings"
	"time"
)

var (
	_ Interface = &Base{}
	//defaultCtxTimeout 默认的上下文超时时间
	defaultCtxTimeout = 15 * time.Second
)

const (
	//ctxTimeoutErr 表示上下文超时后的错误信息
	ctxTimeoutErr = "context deadline exceeded"
)

//InitCtxTimeout 自定义上下文超时时间
func InitCtxTimeout(timeout time.Duration) {
	defaultCtxTimeout = timeout
}

type Base struct {
	req.IamToken
	context.Context
}

func (b *Base) Timeout() time.Duration {
	return defaultCtxTimeout
}

func (b *Base) Validate() res.Interface {
	return res.Succeed()
}

func (b *Base) Action() res.Interface {
	return res.Succeed()
}

func (b *Base) Run(api Interface) (response res.Interface) {
	timeoutCtx, cancelFunc := context.WithTimeout(context.TODO(), api.Timeout())
	defer func() {
		//全局panic recover
		if err := recover(); err != nil {
			printStackAndNotify(api)
			response = res.ErrorMsgsRes(code.Error, err)
		} else {
			if err := response.Check(); err != nil {
				if strings.Contains(err.Error(), ctxTimeoutErr) {
					response.ErrorSave(code.CtxTimeoutError)
				}
			}
		}
		cancelFunc()
	}()
	b.Context = timeoutCtx
	response = api.Validate()
	if response.Check() != nil {
		return
	}
	return api.Action()
}

func printStackAndNotify(api Interface) {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	stack := string(buf[:n])
	fmt.Printf("==> %s\n", stack)
	go notify.SendWithGinCtx(api.GetCtx(), stack)
}
