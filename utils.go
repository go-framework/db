package db

import (
	"fmt"
	"strconv"
)

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
	case string:
		i64, _ := strconv.ParseInt(v, 10, 64)
		return i64
	default:
		str := fmt.Sprintf("%v", cursor)
		i64, _ := strconv.ParseInt(str, 10, 64)
		return i64
	}
}

func GetUint64Cursor(cursor interface{}) uint64 {
	switch v := cursor.(type) {
	case int64:
		return uint64(v)
	case int32:
		return uint64(v)
	case int16:
		return uint64(v)
	case int8:
		return uint64(v)
	case int:
		return uint64(v)
	case uint64:
		return v
	case uint32:
		return uint64(v)
	case uint16:
		return uint64(v)
	case uint8:
		return uint64(v)
	case uint:
		return uint64(v)
	case float32:
		return uint64(v)
	case float64:
		return uint64(v)
	case string:
		u64, _ := strconv.ParseUint(v, 10, 64)
		return u64
	default:
		str := fmt.Sprintf("%v", cursor)
		u64, _ := strconv.ParseUint(str, 10, 64)
		return u64
	}
}

func GetStringCursor(cursor interface{}) string {
	switch v := cursor.(type) {
	case string:
		return v
	case int64:
		return strconv.FormatInt(v, 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case int8:
		return strconv.FormatInt(int64(v), 10)
	case int:
		return strconv.FormatInt(int64(v), 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint16:
		return strconv.FormatUint(uint64(v), 10)
	case uint8:
		return strconv.FormatUint(uint64(v), 10)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	default:
		return fmt.Sprintf("%v", cursor)
	}
}
