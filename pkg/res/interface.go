package res

import (
	"github.com/lstack-org/go-web-framework/pkg/code"
)

type Interface interface {
	//业务错误码
	Code() int
	//业务校验
	Check() error
	//错误缓存
	ErrorSave(err error)
	//错误处理
	ErrorHandle(err error) code.Code
	//无错误响应处理
	SucceedRes(ctx code.Header)
	//基于ServiceCode进行错误响应处理
	ErrorRes(ctx code.Header, serviceCode code.Code)
}

//CustomizeErrorHandler 自定义错误处理
type CustomizeErrorHandler func(err error) code.Code
