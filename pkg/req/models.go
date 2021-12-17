package req

//Paging 用于分页查询
type Paging struct {
	Page     int `json:"page" form:"page" validate:"omitempty,gt=0"`
	PageSize int `json:"pageSize" form:"pageSize" validate:"required_with=Page,gt=0"`
}

//Search 用于模糊查询
type Search struct {
	SearchKey   string `json:"searchKey" form:"searchKey" validate:"omitempty"`
	SearchValue string `json:"searchValue" form:"searchValue" validate:"required_with=SearchKey"`
}

//BatchReq 用于批量操作，Infos数组中一般为某资源的id，常见的适用场景：批量删除资源
type BatchReq struct {
	Infos []string `json:"infos" validate:"required,min=1"`
}
