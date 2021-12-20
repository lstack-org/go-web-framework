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
	//CanPage 用于判断是否可分页
	CanPage() bool
}

//SearchAble 用于模糊查询
type SearchAble interface {
	//Key 表示用于模糊查询的键
	Key() string
	//Value 表示模糊查询的值
	Value() interface{}
	//CanSearch 用于判断是否可模糊查询
	CanSearch() bool
}

//SortAble 用于查询结果排序
type SortAble interface {
	//SortKey 表示用于排序的键
	SortKey() string
	//IsAsc 表示是否是升序排序
	IsAsc() bool
	//CanSort 用于判断是否可排序
	CanSort() bool
}
