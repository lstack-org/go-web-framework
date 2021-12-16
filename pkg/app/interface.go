package app

import (
	"github.com/lstack-org/go-web-framework/pkg/req"
	"github.com/lstack-org/go-web-framework/pkg/res"
)

type Interface interface {
	req.Interface
	Validate() res.Interface
	Action() res.Interface
	Run(Interface) res.Interface
}
