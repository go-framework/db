package helper

import (
	"github.com/jinzhu/gorm"

	"github.com/go-framework/db"
)

func Update(in *gorm.DB, data interface{}, condition ...db.Condition) (*gorm.DB, int64, error) {
	var out = ParseQuery(in, db.NewConditions(condition...))
	out = out.Model(data)
	out = out.Updates(data)
	return out, out.RowsAffected, out.Error
}
