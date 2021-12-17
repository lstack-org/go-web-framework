package res

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lstack-org/go-web-framework/pkg/code"
	k8sErrors "k8s.io/apimachinery/pkg/util/errors"
	"net/http"
	"strings"
)

func Res(ctx *gin.Context, res Interface) {
	err := res.Check()
	if err != nil {
		serviceCode := res.ErrorHandle(err)
		res.ErrorRes(ctx, serviceCode)
		ctx.JSON(serviceCode.HttpCode, res)
	} else {
		res.SucceedRes(ctx)
		ctx.JSON(http.StatusOK, res)
	}
}

//Response 响应请求基本数据结构
type Response struct {
	Err    error       `json:"-"`
	Status int         `json:"status"`
	ResMsg string      `json:"resMsg"`
	Data   interface{} `json:"data"`
}

//Code 返回错误码
func (r *Response) Code() int {
	return r.Status
}

//SucceedRes 请求成功响应处理
func (r *Response) SucceedRes(ctx *gin.Context) {
	r.Status = code.Success.BusinessCode
	r.ResMsg = code.Success.GetMsg(ctx)
}

//ErrorHandle 根据错误，返回错误码，默认返回e.Error
func (r *Response) ErrorHandle(err error) code.ServiceCode {
	switch c := err.(type) {
	case code.ServiceCode:
		return c
	default:
		return code.Error.MergeObj(err.Error())
	}
}

//ErrorRes 根据serviceCode，进行错误响应
func (r *Response) ErrorRes(ctx *gin.Context, serviceCode code.ServiceCode) {
	r.Status = serviceCode.BusinessCode
	r.ResMsg = serviceCode.GetMsg(ctx)
}

//Check 检查返回的响应结果
func (r *Response) Check() error {
	if r.Err != nil {
		return r.Err
	}
	if r.Status != code.SuccessCode {
		r.Err = errors.New(r.ResMsg)
		return r.Err
	}
	return nil
}

//ErrorSave 保存一个错误
func (r *Response) ErrorSave(err error) {
	r.Err = err
}

func Succeed() Interface {
	return &Response{}
}

func SucceedRes(data interface{}) Interface {
	return &Response{
		Data: data,
	}
}

func ListRes(total int, data interface{}) Interface {
	return SucceedRes(ListData{
		Total: total,
		Items: data,
	})
}

func ErrCheckRes(data interface{}, err error) Interface {
	if err != nil {
		return ErrorRes(err)
	}
	return SucceedRes(data)
}

func ErrorRes(err error) Interface {
	s := &Response{}
	s.ErrorSave(err)
	return s
}

func ErrorsRes(code code.ServiceCode, errs ...error) Interface {
	err := k8sErrors.NewAggregate(errs)
	if err != nil {
		return ErrorMsgsRes(code, err.Error())
	}
	return ErrorMsgsRes(code)
}

func ErrorMsgsRes(code code.ServiceCode, mergedMsg ...interface{}) Interface {
	if len(mergedMsg) > 0 {
		var msgs []string
		for _, msg := range mergedMsg {
			msgs = append(msgs, fmt.Sprintf("%v", msg))
		}
		code = code.MergeObj(strings.Join(msgs, ","))
	}
	return ErrorRes(code)
}
