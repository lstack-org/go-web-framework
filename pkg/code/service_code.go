package code

import (
	"fmt"
	"net/http"
	"strings"
)

var _ Code = ServiceCode{}

//ServiceCode 封装错误码，和对应的错误信息
type ServiceCode struct {
	HttpCode     int
	BusinessCode int
	EnglishMsg   string
	ChineseMsg   string
	Format       bool
}

func (s ServiceCode) BusinessStatus() int {
	return s.BusinessCode
}

func (s ServiceCode) HttpStatus() int {
	return s.HttpCode
}

//MergeObj 合并错误信息
func (s ServiceCode) MergeObj(msgs ...interface{}) Code {
	if len(msgs) == 0 {
		return s
	}
	if s.Format {
		s.EnglishMsg = fmt.Sprintf(s.EnglishMsg, msgs...)
		s.ChineseMsg = fmt.Sprintf(s.ChineseMsg, msgs...)
	} else {
		var mergedMsgs []string
		for _, msg := range msgs {
			mergedMsgs = append(mergedMsgs, fmt.Sprintf("%v", msg))
		}
		join := strings.Join(mergedMsgs, ",")
		s.EnglishMsg = fmt.Sprintf("%s,%v", s.EnglishMsg, join)
		s.ChineseMsg = fmt.Sprintf("%s,%v", s.ChineseMsg, join)
	}
	return s
}

//GetMsg 根据配置，返回中文错误或英文错误
func (s ServiceCode) GetMsg(ctx Header) string {
	if ctx == nil {
		return s.EnglishMsg
	}
	switch ctx.GetHeader(AcceptLanguageHeader) {
	case AcceptLanguageZh:
		return s.ChineseMsg
	default:
		return s.EnglishMsg
	}
}

func (s ServiceCode) Error() string {
	return s.EnglishMsg
}

var (
	Success = ServiceCode{
		HttpCode:     http.StatusOK,
		BusinessCode: SuccessCode,
		EnglishMsg:   "succeed",
		ChineseMsg:   "操作成功",
	}

	Error = ServiceCode{
		HttpCode:     http.StatusOK,
		BusinessCode: ErrorCode,
		EnglishMsg:   "Error",
		ChineseMsg:   "错误",
	}

	BindError = ServiceCode{
		HttpCode:     http.StatusOK,
		BusinessCode: ErrorCode,
		EnglishMsg:   "Params error",
		ChineseMsg:   "参数错误",
	}

	CheckTokenError = ServiceCode{
		HttpCode:     http.StatusUnauthorized,
		BusinessCode: Unauthorized,
		EnglishMsg:   "Token error",
		ChineseMsg:   "token解析错误",
	}

	DingNotifyError = ServiceCode{
		HttpCode:     http.StatusOK,
		BusinessCode: DingNotifyFailed,
		EnglishMsg:   "ding notify failed",
		ChineseMsg:   "钉钉通知失败",
	}
)
