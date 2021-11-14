package paginations

type Meta struct {
	Page      int   `json:"page"`
	Limit     int   `json:"limit"`
	ItemCount int64 `json:"itemCount"`
	PageCount int   `json:"pageCount"`
}

type Pagination struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type PaginationQuery struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
}

type PaginationOption struct {
	Page  int
	Limit int
	Data  interface{}
	Model interface{}
}
