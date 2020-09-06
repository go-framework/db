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

func Fields(fields ...string) Condition {
	return conditionFunc(func(conditions *Conditions) {
		conditions.Fields = fields
	})
}

func Order(order ...string) Condition {
	return conditionFunc(func(conditions *Conditions) {
		conditions.Order = order
	})
}

func Group(group ...string) Condition {
	return conditionFunc(func(conditions *Conditions) {
		conditions.Group = group
	})
}

func Having(having ...interface{}) Condition {
	return conditionFunc(func(conditions *Conditions) {
		conditions.Having = having
	})
}

func And(and ...interface{}) Condition {
	return conditionFunc(func(conditions *Conditions) {
		conditions.And = and
	})
}

func Or(or ...interface{}) Condition {
	return conditionFunc(func(conditions *Conditions) {
		conditions.Or = or
	})
}

func Not(not ...interface{}) Condition {
	return conditionFunc(func(conditions *Conditions) {
		conditions.Not = not
	})
}

func Parsed(parsed bool) Condition {
	return conditionFunc(func(conditions *Conditions) {
		conditions.Parsed = parsed
	})
}

func WithCondition(conditions *Conditions) Condition {
	return conditionFunc(func(c *Conditions) {
		if conditions == nil || c == nil {
			return
		}
		c.Parsed = conditions.Parsed
		c.Limit = conditions.Limit
		c.Offset = conditions.Offset
		c.Fields = conditions.Fields
		c.Order = conditions.Order
		c.Group = conditions.Group
		c.Having = conditions.Having
		c.And = conditions.And
		c.Or = conditions.Or
		c.Not = conditions.Not
	})
}
