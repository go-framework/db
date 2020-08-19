package db

// Pagination
type Pagination struct {
	Total  uint        `json:"total"`
	Count  uint        `json:"count"`
	Offset int         `json:"offset"`
	Cursor interface{} `json:"cursor,omitempty"`
}

func NewPagination() *Pagination {
	p := Pagination{}
	return &p
}
