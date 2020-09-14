package redis

import (
	"context"
	"errors"

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

// HashGetList list redis key object not hash field
func HashGetList(ctx context.Context, rdb *redis.Client, list interface{}, conditions *db.Conditions) (pagination *db.Pagination, err error) {
	if conditions == nil {
		return pagination, errors.New("parameter: *db.Conditions is nil")
	}
	if len(conditions.Match) == 0 {
		return pagination, errors.New("match condition is not set")
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
		return pagination, errors.New("list have not impl MapUnmarshaler interface")
	}

	for {
		// scan match keys
		keys, cursor, err = rdb.Scan(ctx, cursor, conditions.Match, count).Result()
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
				return
			}
			// key is not exist
			if value == nil || len(value) == 0 {
				err = &NotExistError{
					Key: key,
				}
				errs = append(errs, err)
				continue
			}
			// map unmarshal
			err = mapUnmarshaler.UnmarshalMap(value)
			if err != nil {
				errs = append(errs, &RedisError{
					Key:     key,
					Operate: "UnmarshalMap",
					Err:     err.Error(),
				})
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
