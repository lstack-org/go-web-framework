package notify

type Interface interface {
	//Validate 用于校验相关参数，true时参数正常
	Validate() bool
	//Send 发送通知
	Send(body string) error
}
