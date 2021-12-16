package notify

import (
	"context"
	"fmt"
	"github.com/lstack-org/go-web-framework/pkg/notify/templates"
	"github.com/lstack-org/go-web-framework/pkg/notify/third/ding"
)

var (
	dOpts dingOpts
	_     Interface = &dingDing{}
)

func InitDingNotify(dingHook string) {
	dOpts.dingHook = dingHook
}

func NewDingDing(client ding.Client) Interface {
	return &dingDing{client: client, dingOpts: dOpts}
}

type dingOpts struct {
	dingHook string
}

type dingDing struct {
	dingOpts
	client ding.Client
}

func (d *dingDing) Validate() bool {
	return d.dingHook != ""
}

func (d *dingDing) Send(body string) error {
	return d.client.Notify(context.TODO(), ding.Message{
		Type: "text",
		Text: ding.Text{
			Content: body,
		},
		Token: ding.AccessToken{
			AccessToken: d.dingHook,
		},
	})
}

func NewPanicDing(method, host, uri, stack string) (body string) {
	url := fmt.Sprintf("%s: %s:%s", method, host, uri)
	return fmt.Sprintf(templates.PanicDing, url, stack)
}
