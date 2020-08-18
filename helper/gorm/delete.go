package helper

import (
	"context"

	"github.com/jinzhu/gorm"

	"github.com/go-framework/db"
)

func Delete(ctx context.Context, in *gorm.DB, data interface{}, conditions *db.Conditions) (_db *gorm.DB, err error) {
	_db = ParseContext(ctx, in)
	_db = ParseQuery(_db, conditions)

	if _, ok := _db.Get(TableNameKey); !ok {
		_db = _db.Model(data)
	}
	err = _db.Delete(data).Error
	return
}
