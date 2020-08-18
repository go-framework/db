package helper

import (
	"context"

	"github.com/jinzhu/gorm"

	"github.com/go-framework/db"
)

func Update(ctx context.Context, in *gorm.DB, data interface{}, conditions *db.Conditions) (*gorm.DB, int64, error) {
	_db := ParseContext(ctx, in)
	_db = ParseQuery(_db, conditions)
	if _, ok := _db.Get(TableNameKey); !ok {
		_db = _db.Model(data)
	}
	_db = _db.Updates(data)
	return _db, _db.RowsAffected, _db.Error
}
