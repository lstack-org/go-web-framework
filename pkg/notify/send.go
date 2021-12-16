package notify

import (
	"github.com/gin-gonic/gin"
	"github.com/lstack-org/go-web-framework/pkg/notify/third/ding"
	"k8s.io/klog/v2"
)

func SendWithGinCtx(ctx *gin.Context, stack string) {
	EmailNotify(ctx, stack)
	DingNotify(ctx, stack)
}

//EmailNotify 用于邮件通知程序panic
func EmailNotify(ctx *gin.Context, stack string) {
	request := ctx.Request
	subject, body, err := NewPanicHTMLEmail(request.Method, request.Host, request.URL.String(), stack)
	if err != nil {
		klog.Error(err)
	}
	m := NewMail(subject)
	if m.Validate() {
		err := m.Send(body)
		if err != nil {
			klog.Error(err)
		}
	}
}

//DingNotify 用于钉钉通知panic
func DingNotify(ctx *gin.Context, stack string) {
	request := ctx.Request
	client := ding.NewClient(nil)
	body := NewPanicDing(request.Method, request.Host, request.URL.String(), stack)
	d := NewDingDing(client)
	if d.Validate() {
		err := d.Send(body)
		if err != nil {
			klog.Error(err)
		}
	}
}
