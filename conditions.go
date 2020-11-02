package db

//go:generate easyjson -all conditions.go

// Conditions
// swagger:model Conditions
type Conditions struct {
	// parse this Conditions flag.
	parsed bool `json:"-" form:"-"`

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
	// Match
	Match string `json:"match,omitempty" form:"match"`
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
	//
	// In: query
	// Type: array
	// Unique: true
	// CollectionFormat: multi
	Group []string `json:"group,omitempty" form:"group"`
	// Having
	Having []interface{} `json:"having,omitempty" form:"having"`
	// And list
	And []interface{} `json:"and,omitempty" form:"and"`
	// Or list
	Or []interface{} `json:"or,omitempty" form:"or"`
	// Not list
	Not []interface{} `json:"not,omitempty" form:"not"`
	// Between
	Between [2]interface{} `json:"between,omitempty" form:"between"`
}

func GetDefaultConditions() *Conditions {
	conditions := Conditions{
		Limit:  20,
		Offset: 0,
	}
	return &conditions
}

func NewConditions(condition ...Condition) *Conditions {
	conditions := GetDefaultConditions()
	for _, item := range condition {
		item.apply(conditions)
	}
	return conditions
}

func NewNilConditions(condition ...Condition) *Conditions {
	if len(condition) == 0 {
		return nil
	}
	conditions := GetDefaultConditions()
	for _, item := range condition {
		item.apply(conditions)
	}
	return conditions
}

func (conditions *Conditions) SetParsed(parsed bool) {
	conditions.parsed = parsed
}

func (conditions Conditions) GetParsed() bool {
	return conditions.parsed
}

func (conditions *Conditions) NewPagination() *Pagination {
	p := Pagination{
		Offset: conditions.Offset,
	}
	return &p
}

func (conditions *Conditions) GetInt64Cursor() int64 {
	return GetInt64Cursor(conditions.Cursor)
}

func (conditions *Conditions) GetUint64Cursor() uint64 {
	return GetUint64Cursor(conditions.Cursor)
}

func (conditions *Conditions) GetStringCursor() string {
	return GetStringCursor(conditions.Cursor)
}
