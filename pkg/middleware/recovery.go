package middleware

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lstack-org/go-web-framework/pkg/notify"
)

//Recovery gin中间件、执行路由时发送panic，进行recover，并发送一封通知邮件
func Recovery() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		buf := new(bytes.Buffer)
		defer func() {
			stack := buf.String()
			if stack != "" {
				fmt.Println(stack)
				go notify.SendWithGinCtx(ctx, stack)

			}
		}()

		gin.RecoveryWithWriter(buf)(ctx)
	}
}
