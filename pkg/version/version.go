package version

import (
	"github.com/gin-gonic/gin"
	"k8s.io/klog/v2"
	"net/http"
)

var (
	AppVersion string
	GoVersion  string
	BuildTime  string
	GitCommit  string

	v versionInfo
)

type versionInfo struct {
	AppVersion string `json:"appVersion"`
	GoVersion  string `json:"goVersion"`
	BuildTime  string `json:"buildTime"`
	GitCommit  string `json:"gitCommit"`
}

func init() {
	if AppVersion != "" {
		v.AppVersion = AppVersion
		klog.Infof("app version: %s", AppVersion)
	}

	if GoVersion != "" {
		v.GoVersion = GoVersion
		klog.Infof("go version: %s", GoVersion)
	}

	if BuildTime != "" {
		v.BuildTime = BuildTime
		klog.Infof("build time: %s", BuildTime)
	}

	if GitCommit != "" {
		v.GitCommit = GitCommit
		klog.Infof("git commit: %s", GitCommit)
	}
}

func Version(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, v)
}
