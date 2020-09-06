package db

import (
	"fmt"
	"strconv"
)

// 获取int64类型的游标
func GetInt64Cursor(cursor interface{}) int64 {
	switch v := cursor.(type) {
	case int64:
		return v
	case int32:
		return int64(v)
	case int16:
		return int64(v)
	case int8:
		return int64(v)
	case int:
		return int64(v)
	case uint64:
		return int64(v)
	case uint32:
		return int64(v)
	case uint16:
		return int64(v)
	case uint8:
		return int64(v)
	case uint:
		return int64(v)
	case float32:
		return int64(v)
	case float64:
		return int64(v)
	default:
		str := fmt.Sprintf("%v", cursor)
		i64, _ := strconv.ParseInt(str, 10, 64)
		return i64
	}
}

// 获取字符串类型的游标
func GetStringCursor(cursor interface{}) string {
	switch s := cursor.(type) {
	case string:
		return s
	default:
		return fmt.Sprintf("%s", cursor)
	}
}
