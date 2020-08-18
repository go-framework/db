package helper

import (
	"context"

	"github.com/jinzhu/gorm"

	"github.com/go-framework/db"
)

func Delete(ctx context.Context, in *gorm.DB, data interface{}, condition ...db.Condition) (_db *gorm.DB, err error) {
	_db = ParseContext(ctx, in)
	if len(condition) > 0 {
		_db = ParseQuery(_db, db.NewConditions(condition...))
	}
	err = _db.Model(data).Delete(data).Error
	return
}
