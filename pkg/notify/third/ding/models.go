package ding

type Message struct {
	Type  string      `json:"msgtype" validate:"required"`
	Text  Text        `json:"text,omitempty" validate:"required"`
	Token AccessToken `json:"token"`
}

type Text struct {
	Content string `json:"content" validate:"required"`
}

type Res struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
}
