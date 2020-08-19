package helper

import (
	"context"
	"strings"

	"github.com/jinzhu/gorm"

	"github.com/go-framework/db"
)

const (
	TableNameKey = "table"
)

func ParseQuery(db *gorm.DB, conditions *db.Conditions) *gorm.DB {
	if conditions == nil {
		return db
	}
	if conditions.Parsed {
		return db
	}
	defer func() {
		conditions.Parsed = true
	}()

	var _db *gorm.DB = db

	if conditions.Limit > 0 {
		_db = _db.Limit(conditions.Limit)
	}
	if conditions.Offset > 0 {
		_db = _db.Offset(conditions.Offset)
	}
	if len(conditions.Fields) > 0 {
		_db = _db.Select(conditions.Fields)
	}
	if len(conditions.Order) > 0 {
		_db = _db.Order(strings.Join(conditions.Order, ","))
	}
	if len(conditions.GroupBy) > 0 {
		_db = _db.Group(strings.Join(conditions.GroupBy, ","))
	}

	switch len(conditions.Having) {
	case 1:
		_db = _db.Having(conditions.Having[0])
	case 0:
	default:
		_db = _db.Having(conditions.Having[0], conditions.Having[1:]...)
	}

	switch len(conditions.And) {
	case 1:
		_db = _db.Where(conditions.And[0])
	case 0:
	default:
		_db = _db.Where(conditions.And[0], conditions.And[1:]...)
	}

	switch len(conditions.Or) {
	case 1:
		_db = _db.Or(conditions.Or[0])
	case 0:
	default:
		_db = _db.Or(conditions.Or[0], conditions.Or[1:]...)
	}

	switch len(conditions.Not) {
	case 1:
		_db = _db.Not(conditions.Not[0])
	case 0:
	default:
		_db = _db.Not(conditions.Not[0], conditions.Not[1:]...)
	}

	return _db
}

func ParseContext(ctx context.Context, db_ *gorm.DB) *gorm.DB {
	// debug
	debug := db.GetDebugFromContext(ctx)
	var _db = db_.LogMode(debug)
	// table
	table := db.GetTableFromContext(ctx)
	if len(table) > 0 {
		_db = _db.Table(table)
		_db.InstantSet(TableNameKey, table)
	}
	// unscoped
	unscoped := db.GetUnscopedFromContext(ctx)
	if unscoped {
		_db = _db.Unscoped()
	}
	return _db
}
