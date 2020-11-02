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

func Match(match string) Condition {
	return conditionFunc(func(conditions *Conditions) {
		conditions.Match = match
	})
}

func Cursor(cursor interface{}) Condition {
	return conditionFunc(func(conditions *Conditions) {
		conditions.Cursor = cursor
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

func Between(start, end interface{}) Condition {
	return conditionFunc(func(conditions *Conditions) {
		conditions.Between = [2]interface{}{start, end}
	})
}

func Parsed(parsed bool) Condition {
	return conditionFunc(func(conditions *Conditions) {
		conditions.parsed = parsed
	})
}

func WithCondition(conditions *Conditions) Condition {
	return conditionFunc(func(c *Conditions) {
		if conditions == nil || c == nil {
			return
		}
		c.parsed = conditions.parsed
		c.Limit = conditions.Limit
		c.Offset = conditions.Offset
		c.Match = conditions.Match
		c.Cursor = conditions.Cursor
		c.Fields = conditions.Fields
		c.Order = conditions.Order
		c.Group = conditions.Group
		c.Having = conditions.Having
		c.And = conditions.And
		c.Or = conditions.Or
		c.Not = conditions.Not
		c.Between = conditions.Between
	})
}
