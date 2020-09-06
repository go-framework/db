package helper

import (
	"context"

	"github.com/jinzhu/gorm"
)

func Create(ctx context.Context, db *gorm.DB, data interface{}) (_db *gorm.DB, err error) {
	_db = ParseContext(ctx, db)

	if _, ok := _db.Get(TableNameKey); !ok {
		_db = _db.Model(data)
	}

	err = _db.Create(data).Error
	return
}
