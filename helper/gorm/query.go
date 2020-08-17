package helper

import (
	"context"
	"database/sql"

	"github.com/jinzhu/gorm"

	"github.com/go-framework/db"
)

func Count(ctx context.Context, db *gorm.DB, conditions *db.Conditions, list interface{}) (_db *gorm.DB, total uint, err error) {
	_db = ParseContext(ctx, db)
	_db = ParseQuery(_db, conditions)

	err = _db.Model(&list).Limit(-1).Offset(-1).Count(&total).Error
	if err == sql.ErrNoRows || total == 0 {
		err = nil
		return
	}
	return
}

func List(ctx context.Context, db *gorm.DB, conditions *db.Conditions, list interface{}) (_db *gorm.DB, total uint, err error) {
	_db = ParseContext(ctx, db)
	_db = ParseQuery(_db, conditions)

	_db = _db.Model(&list)
	err = _db.Limit(-1).Offset(-1).Count(&total).Error
	if err == sql.ErrNoRows || total == 0 {
		err = nil
		return
	}

	err = _db.Find(&list).Error
	return
}

func Find(ctx context.Context, db *gorm.DB, conditions *db.Conditions, object interface{}) (_db *gorm.DB, err error) {
	_db = ParseContext(ctx, db)
	_db = ParseQuery(_db, conditions)

	_db = _db.Model(object)
	err = _db.First(object).Error
	return
}