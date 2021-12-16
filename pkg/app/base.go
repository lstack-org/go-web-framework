package app

import (
	"fmt"
	"github.com/lstack-org/go-web-framework/pkg/code"
	"github.com/lstack-org/go-web-framework/pkg/notify"
	"github.com/lstack-org/go-web-framework/pkg/req"
	"github.com/lstack-org/go-web-framework/pkg/res"
	"runtime"
)

var _ Interface = &Base{}

type Base struct {
	req.IamToken
}

func (b *Base) Validate() res.Interface {
	return res.Succeed()
}

func (b *Base) Action() res.Interface {
	return res.Succeed()
}

func (b *Base) Run(api Interface) (response res.Interface) {
	defer func() {
		if err := recover(); err != nil {
			printStackAndNotify(api)
			response = res.ErrorMsgsRes(code.Error, err)
		}
	}()
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
