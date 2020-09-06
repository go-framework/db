package db

//go:generate easyjson -all conditions.go

// Conditions
// swagger:model Conditions
type Conditions struct {
	// Parse this Conditions flag.
	Parsed bool `json:"-"`

	// Limit record
	//
	// In: query
	// Type: integer
	Limit int `json:"limit,omitempty" form:"limit"`
	// Offset record
	//
	// In: query
	// Type: integer
	Offset int `json:"offset,omitempty" form:"offset"`
	// Cursor
	//
	// if exist Cursor set back.
	//
	// In: query
	Cursor interface{} `json:"cursor,omitempty" form:"cursor"`
	// Select fields
	//
	// In: query
	// Type: array
	// Unique: true
	// CollectionFormat: multi
	Fields []string `json:"fields,omitempty" form:"fields"`
	// Sort order
	//
	// In: query
	// Type: array
	// Unique: true
	// CollectionFormat: multi
	Order []string `json:"order,omitempty" form:"order"`
	// Group fields
	Group []string `json:"group" form:"group"`
	// Having
	Having []interface{} `json:"having" form:"having"`
	// And list
	And []interface{} `json:"and" form:"and"`
	// Or list
	Or []interface{} `json:"or" form:"or"`
	// Not list
	Not []interface{} `json:"not" form:"not"`
}

func NewConditions(condition ...Condition) *Conditions {
	conditions := Conditions{}
	for _, item := range condition {
		item.apply(&conditions)
	}
	return &conditions
}

func NewNilConditions(condition ...Condition) *Conditions {
	if len(condition) == 0 {
		return nil
	}
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
		Offset: conditions.Offset,
	}
	return &p
}

// 获取int64类型的游标
func (conditions *Conditions) GetInt64Cursor() int64 {
	return GetInt64Cursor(conditions.Cursor)
}

// 获取字符串类型的游标
func (conditions *Conditions) GetStringCursor() string {
	return GetStringCursor(conditions.Cursor)
}
