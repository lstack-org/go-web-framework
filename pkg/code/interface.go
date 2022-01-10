package code

import "github.com/gin-gonic/gin"

type Code interface {
	error
	//BusinessStatus 表示业务错误码
	BusinessStatus() int
	//HttpStatus 表示http错误码
	HttpStatus() int
	//GetMsg 用于获取错误码描述信息，参数ctx用于获取描述信息语言类型如中文、英文等
	GetMsg(ctx *gin.Context) string
	//MergeObj 用于合并错误信息
	MergeObj(msg interface{}) Code
}
