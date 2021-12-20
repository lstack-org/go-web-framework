package req

var _ PageAble = Common{}
var _ SearchAble = Common{}
var _ SortAble = Common{}

type Common struct {
	*Paging
	*Search
	*Sort
}

var _ PageAble = &Paging{}

type Paging struct {
	Page     int `json:"page" form:"page" validate:"required_with=PageSize,omitempty,gt=0"`
	PageSize int `json:"pageSize" form:"pageSize" validate:"required_with=Page,omitempty,gt=0"`
}

func (p *Paging) CanPage() bool {
	if p == nil {
		return false
	}
	return p.Page > 0 && p.PageSize > 0
}

func (p *Paging) PNumber() int {
	return p.Page
}

func (p *Paging) PSize() int {
	return p.PageSize
}

var _ SearchAble = &Search{}

type Search struct {
	SearchKey   string `json:"searchKey" form:"searchKey" validate:"required_with=SearchValue,omitempty"`
	SearchValue string `json:"searchValue" form:"searchValue" validate:"required_with=SearchKey,omitempty"`
}

func (s *Search) CanSearch() bool {
	if s == nil {
		return false
	}
	return s.SearchKey != "" && s.SearchValue != ""
}

func (s *Search) Key() string {
	return s.SearchKey
}

func (s *Search) Value() interface{} {
	return s.SearchValue
}

var _ SortAble = &Sort{}

type Sort struct {
	Asc       bool   `json:"asc" form:"asc" validate:"omitempty"`
	SortField string `json:"sortField" form:"sortField" validate:"required_with=Asc"`
}

func (s *Sort) CanSort() bool {
	if s == nil {
		return false
	}
	return s.SortField != ""
}

func (s *Sort) SortKey() string {
	return s.SortField
}

func (s *Sort) IsAsc() bool {
	return s.Asc
}

//BatchReq 用于批量操作，Infos数组中一般为某资源的id，常见的适用场景：批量删除资源
type BatchReq struct {
	Infos []string `json:"infos" validate:"required,min=1"`
}
