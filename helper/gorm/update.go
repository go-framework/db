package helper

import (
	"context"

	"github.com/jinzhu/gorm"

	"github.com/go-framework/db"
)

func Update(ctx context.Context,in *gorm.DB, data interface{}, condition ...db.Condition) (*gorm.DB, int64, error) {
	_db := ParseContext(ctx, in)
	if len(condition) > 0 {
		_db = ParseQuery(_db, db.NewConditions(condition...))
	}
	if _, ok := _db.Get(TableNameKey); !ok {
		_db = _db.Model(data)
	}
	_db = _db.Updates(data)
	return _db, _db.RowsAffected, _db.Error
}
