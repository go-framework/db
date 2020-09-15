package redis

import (
	"context"
	"errors"
	"strings"

	"github.com/go-redis/redis/v8"

	"github.com/go-framework/db"
)

func HashUpdateMap(ctx context.Context, rdb *redis.Client, key string, data map[string]interface{}) error {
	expire, ok := GetExpireFromContext(ctx)
	_, err := rdb.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		err := pipe.HMSet(ctx, key, data).Err()
		if err != nil {
			return err
		}
		if ok {
			err = pipe.Expire(ctx, key, expire).Err()
		}
		return err
	})
	return err
}

func HashUpdate(ctx context.Context, rdb *redis.Client, key string, data MapMarshaler) error {
	value, err := data.MarshalMap()
	if err != nil {
		return err
	}
	return HashUpdateMap(ctx, rdb, key, value)
}

func HashGetOne(ctx context.Context, rdb *redis.Client, key string, data MapUnmarshaler) error {
	value, err := rdb.HGetAll(ctx, key).Result()
	if err != nil {
		return err
	}
	if value == nil || len(value) == 0 {
		return &NotExistError{
			Key: key,
		}
	}
	return data.UnmarshalMap(value)
}

func ScanKeys(ctx context.Context, rdb *redis.Client, cursor uint64, match string, count int64) ([]string, uint64, error) {
	return rdb.Scan(ctx, cursor, match, count).Result()
}

// HashGetList list redis key object not hash field
func HashGetList(ctx context.Context, rdb *redis.Client, list interface{}, conditions *db.Conditions) (pagination *db.Pagination, err error) {
	if list == nil {
		return pagination, errors.New("parameter: list is nil")
	}
	if conditions == nil {
		return pagination, errors.New("parameter: *db.Conditions is nil")
	}
	if len(conditions.Match) == 0 {
		return pagination, errors.New("conditions: match field is not set")
	}
	// set default
	if conditions.Limit <= 0 {
		conditions.Limit = 10
	}
	var (
		keys    []string
		value   map[string]string
		errs    = make(RedisErrors, 0)
		getKeys = make(map[string]struct{}) // for remove duplication
	)
	pagination = conditions.NewPagination()
	cursor := conditions.GetUint64Cursor()
	count := int64(conditions.Limit)

	// just support UnmarshalMap
	mapUnmarshaler, ok := list.(MapUnmarshaler)
	if !ok {
		return pagination, errors.New("list have not implement MapUnmarshaler interface")
	}

	var (
		scanKeyFunc = ScanKeys
	)

	// get ScanKeyFunc from context
	if f, ok := GetScanKeyFuncFromContext(ctx); ok {
		scanKeyFunc = f
	}

	for {
		// scan match keys
		keys, cursor, err = scanKeyFunc(ctx, rdb, cursor, conditions.Match, count)
		if err != nil {
			return
		}
		for _, key := range keys {
			// remove duplication
			if _, ok = getKeys[key]; ok {
				continue
			}
			getKeys[key] = struct{}{}
			// get all hash fields
			value, err = rdb.HGetAll(ctx, key).Result()
			if err != nil {
				// wrong type error
				if strings.Contains(err.Error(), "WRONGTYPE") {
					errs = append(errs, &RedisError{
						Key:     key,
						Operate: "HGetAll",
						Err:     err.Error(),
					})
					continue
				}
				return
			}
			// map unmarshal
			err = mapUnmarshaler.UnmarshalMap(value)
			if err != nil {
				errs = append(errs, &RedisError{
					Key:     key,
					Operate: "UnmarshalMap",
					Err:     err.Error(),
				})
				continue
			}
		}
		if cursor == 0 {
			break
		}
		if len(getKeys) >= conditions.Limit {
			break
		}
	}
	// update pagination
	pagination.Cursor = cursor
	pagination.Count = uint(len(getKeys))
	return pagination, errs.Nil()
}
