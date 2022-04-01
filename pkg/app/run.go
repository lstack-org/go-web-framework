package app

import (
	"github.com/gin-gonic/gin"
	"github.com/lstack-org/go-web-framework/pkg/req"
	"github.com/lstack-org/go-web-framework/pkg/res"
)

//Run 执行参数绑定，以及 Interface.Run方法
func Run(ctx *gin.Context, api Interface, binders ...req.Binder) {
	err := req.Bind(ctx, api, binders...)
	if err != nil {
		return
	}
	res.Res(ctx, api.Run(api))
}

//RunQuery 自带query参数绑定，一般用于get请求
func RunQuery(ctx *gin.Context, api Interface, binders ...req.Binder) {
	binders = append(binders, req.BindQuery)
	Run(ctx, api, binders...)
}

//RunJSON 自带JSON参数绑定,一般用于有请求体的请求
func RunJSON(ctx *gin.Context, api Interface, binders ...req.Binder) {
	binders = append(binders, req.BindJSON)
	Run(ctx, api, binders...)
}

//RunURI 自带uri参数绑定
func RunURI(ctx *gin.Context, api Interface, binders ...req.Binder) {
	binders = append(binders, req.BindUri)
	Run(ctx, api, binders...)
}
