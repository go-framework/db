package helper

import (
	"context"

	"github.com/jinzhu/gorm"
)

func Create(ctx context.Context, db *gorm.DB, data interface{}) (_db *gorm.DB, err error) {
	_db = ParseContext(ctx, db)

	err = _db.Model(data).Create(data).Error
	return
}
