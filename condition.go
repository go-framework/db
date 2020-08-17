package db

// Condition
type Condition interface {
	apply(*Conditions)
}

// condition func
type conditionFunc func(*Conditions)

func (f conditionFunc) apply(condition *Conditions) {
	f(condition)
}

func WithConditions(conditions *Conditions) Condition {
	return conditionFunc(func(conditions_ *Conditions) {
		conditions_ = conditions
	})
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
