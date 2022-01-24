package code

type Code interface {
	error
	//BusinessStatus 表示业务错误码
	BusinessStatus() int
	//HttpStatus 表示http错误码
	HttpStatus() int
	//GetMsg 用于获取错误码描述信息，参数ctx用于获取描述信息语言类型如中文、英文等
	GetMsg(ctx Header) string
	//MergeObj 用于合并错误信息
	MergeObj(msgs ...interface{}) Code
}

type Header interface {
	//GetHeader 获取请求头，key为header的键，返回其值
	GetHeader(key string) string
}
