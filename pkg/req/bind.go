package req

import (
	"github.com/gin-gonic/gin"
	"github.com/lstack-org/go-web-framework/pkg/code"
	"github.com/lstack-org/go-web-framework/pkg/res"
)

var (
	defaultBinders = []Binder{BindJSON, BindQuery, BindHeader, BindUri}
)

//Bind 用户请求参数绑定
func Bind(ctx *gin.Context, object Interface, binders ...Binder) error {
	if customizeBinder, ok := object.(CustomizeBinder); ok {
		err := customizeBinder.Bind(ctx)
		if err != nil {
			res.Res(ctx, res.ErrorMsgsRes(code.BindError, err))
			return err
		}
	}else {
		if len(binders) == 0 {
			binders = defaultBinders
		}

		for _, binder := range binders {
			if binder != nil {
				err := binder(ctx, object)
				if err != nil {
					res.Res(ctx, res.ErrorMsgsRes(code.BindError, err))
					return err
				}
			}
		}
	}

	if customizeValidater, ok := object.(CustomizeArgsValidater); ok {
		if customizeValidater.SkipArgsValidate() {
			object.SetCtx(ctx)
			return nil
		}
	}

	if err := Val.Struct(object); err != nil {
		res.Res(ctx, res.ErrorMsgsRes(code.BindError, translate(ctx, err)))
		return err
	}

	object.SetCtx(ctx)
	return nil
}

//BindJSON 用于绑定请求体
func BindJSON(ctx *gin.Context, obj interface{}) error {
	err := ctx.ShouldBindJSON(obj)
	if err != nil {
		if err.Error() != "EOF" {
			return err
		}
	}
	return nil
}

//BindHeader 用于绑定请求头
func BindHeader(ctx *gin.Context, obj interface{}) error {
	return ctx.ShouldBindHeader(obj)
}

//BindQuery 用于版本query参数
func BindQuery(ctx *gin.Context, obj interface{}) error {
	return ctx.ShouldBindQuery(obj)
}

//BindUri 用于版本路径参数
func BindUri(ctx *gin.Context, obj interface{}) error {
	return ctx.ShouldBindUri(obj)
}
