package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/lstack-org/go-web-framework/pkg/code"
	"github.com/lstack-org/go-web-framework/pkg/req"
	"github.com/lstack-org/go-web-framework/pkg/res"
	"net/url"
)

//ParseToken 用户token解析
func ParseToken(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	errHandle := func(errMsg string) {
		ctx.Abort()
		res.Res(ctx, res.ErrorMsgsRes(code.CheckTokenError, errMsg))
	}
	if token == "" {
		errHandle("token not found")
		return
	}

	tokenCp := token
	token, err := url.PathUnescape(token)
	if err != nil {
		errHandle("token pathUnescape failed")
		return
	}
	var iamToken req.IamToken
	err = json.Unmarshal([]byte(token), &iamToken)
	if err != nil {
		errHandle("token marshal failed")
		return
	}

	iamToken.Token = tokenCp
	ctx.Set("token", iamToken)
	ctx.Next()
}
