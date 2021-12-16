package ding

import (
	"context"
	"github.com/lstack-org/go-web-framework/pkg/code"
	"github.com/lstack-org/utils/pkg/rest"
	"net/http"
)

const (
	dingDingAddr = "https://oapi.dingtalk.com"
)

var DefaultClient = NewClient(nil)

type Client interface {
	Notify(ctx context.Context, body Message) error
}

func NewClient(customize *http.Client) Client {
	client, _ := rest.NewRESTClientEasy("ding", dingDingAddr, customize)
	return &clientImpl{client: client}
}

type clientImpl struct {
	client *rest.RESTClient
}

func (c *clientImpl) Notify(ctx context.Context, body Message) error {
	res := &Res{}
	err := c.client.
		Post().
		AbsPath("/robot/send").
		Params(&AccessToken{
			AccessToken: body.Token.AccessToken,
		}).
		Body(body).
		DoInto(ctx, res)
	if err != nil {
		return err
	}
	if res.ErrCode != 0 {
		return code.DingNotifyError.MergeObj(res.ErrMsg)
	}
	return nil
}
