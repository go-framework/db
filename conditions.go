package db

// Conditions
type Conditions struct {
	Parsed bool `json:"-"`

	Limit   int           `json:"limit"`
	Offset  int           `json:"offset"`
	Fields  []string      `json:"fields"`
	Order   []string      `json:"order"`
	GroupBy []string      `json:"groupBy"`
	Having  []interface{} `json:"having"`
	And     []interface{} `json:"and"`
	Or      []interface{} `json:"or"`
	Not     []interface{} `json:"not"`
}

func NewConditions(condition ...Condition) *Conditions {
	conditions := Conditions{}
	for _, item := range condition {
		item.apply(&conditions)
	}
	return &conditions
}

func (conditions *Conditions) SetParsed(parsed bool) {
	conditions.Parsed = parsed
}

func (conditions *Conditions) NewPagination() *Pagination {
	p := Pagination{
		Offset: uint(conditions.Offset),
	}
	if p.Offset < 0 {
		p.Offset = 0
	}
	return &p
}
