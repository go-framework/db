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

// Condition
type Condition interface {
	apply(*Conditions)
}

// condition func
type conditionFunc func(*Conditions)

func (f conditionFunc) apply(condition *Conditions) {
	f(condition)
}

func WithCondition(condition ...Condition) *Conditions {
	conditions := Conditions{}
	for _, item := range condition {
		item.apply(&conditions)
	}
	return &conditions
}

func Limit(limit int) Condition {
	return conditionFunc(func(conditions *Conditions) {
		conditions.Limit = limit
	})
}

func Offset(offset int) Condition {
	return conditionFunc(func(conditions *Conditions) {
		conditions.Offset = offset
	})
}

func Fields(fields []string) Condition {
	return conditionFunc(func(conditions *Conditions) {
		conditions.Fields = fields
	})
}

func Order(order []string) Condition {
	return conditionFunc(func(conditions *Conditions) {
		conditions.Order = order
	})
}

func GroupBy(groupBy []string) Condition {
	return conditionFunc(func(conditions *Conditions) {
		conditions.GroupBy = groupBy
	})
}

func Having(having []interface{}) Condition {
	return conditionFunc(func(conditions *Conditions) {
		conditions.Having = having
	})
}

func And(and []interface{}) Condition {
	return conditionFunc(func(conditions *Conditions) {
		conditions.And = and
	})
}

func Or(or []interface{}) Condition {
	return conditionFunc(func(conditions *Conditions) {
		conditions.Or = or
	})
}

func Not(not []interface{}) Condition {
	return conditionFunc(func(conditions *Conditions) {
		conditions.Not = not
	})
}
