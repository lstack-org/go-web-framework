package req

import "github.com/gin-gonic/gin"

type Interface interface {
	GetCtx() *gin.Context
	SetCtx(ctx *gin.Context)
}

type Binder func(ctx *gin.Context, obj interface{}) error

//CustomizeBinder 自定义参数绑定
type CustomizeBinder interface {
	Bind(ctx *gin.Context) error
}

//PageAble 用于分页查询
type PageAble interface {
	//PNumber 表示第几页
	PNumber() int
	//PSize 表示每页长度
	PSize() int
}

//SearchAble 用于模糊查询
type SearchAble interface {
	//Key 表示用于模糊查询的键
	Key() string
	//Value 表示模糊查询的值
	Value() interface{}
	//ToMap 用于将SearchAble转换为一个map[string]interface{}
	ToMap() map[string]interface{}
}

//SortAble 用于查询结果排序
type SortAble interface {
	//Key 表示用于排序的键
	Key() string
	//IsAsc 表示是否是升序排序
	IsAsc() bool
}
