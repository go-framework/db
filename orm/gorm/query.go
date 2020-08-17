package gorm

import (
	"database/sql"

	"github.com/jinzhu/gorm"

	"github.com/go-framework/db"
)

func Count(db *gorm.DB, conditions *db.Conditions, list interface{}) (_db *gorm.DB, total uint, err error) {
	_db = parseQuery(db, conditions)

	_db = _db.Model(&list)
	err = _db.Limit(-1).Offset(-1).Count(&total).Error
	if err == sql.ErrNoRows || total == 0 {
		err = nil
		return
	}
	return
}

func List(db *gorm.DB, conditions *db.Conditions, list interface{}) (_db *gorm.DB, total uint, err error) {
	_db, total, err = Count(db, conditions, list)
	if err != nil {
		return
	}
	err = _db.Find(&list).Error
	return
}

func Find(db *gorm.DB, conditions *db.Conditions, object interface{}) (_db *gorm.DB, err error) {
	_db = parseQuery(db, conditions)

	_db = _db.Model(&object)
	err = _db.Find(&object).Error
	return
}
