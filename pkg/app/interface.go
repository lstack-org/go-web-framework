package app

import (
	"context"
	"github.com/lstack-org/go-web-framework/pkg/req"
	"github.com/lstack-org/go-web-framework/pkg/res"
	"time"
)

type Interface interface {
	context.Context
	req.Interface
	//Timeout 配置全局上下文的超时时间
	Timeout() time.Duration
	Validate() res.Interface
	Action() res.Interface
	Run(Interface) res.Interface
}
