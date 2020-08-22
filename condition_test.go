package db

import (
	"reflect"
	"testing"
)

func TestWithCondition(t *testing.T) {
	conditions := NewConditions(Limit(10), And("id=1"))
	condition := WithCondition(conditions)
	var conditions2 = NewConditions(condition)
	if !reflect.DeepEqual(conditions, conditions2) {
		t.Errorf("WithCondition() failed\n\tgot %v\n\twant %v", conditions2, conditions)
	}
}
