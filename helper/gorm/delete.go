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
	if _, ok := _db.Get(TableNameKey); !ok {
		_db = _db.Model(data)
	}
	err = _db.Delete(data).Error
	return
}
