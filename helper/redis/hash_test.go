package redis_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"

	"github.com/go-framework/db"
	helper "github.com/go-framework/db/helper/redis"
)

type testT struct {
	Int     int64
	Uint64  uint64
	Float32 float32
	Float64 float64
	Bytes   []byte
	String  string
	Time    time.Time
}

func (v testT) MarshalMap() (map[string]interface{}, error) {
	var ret = make(map[string]interface{})
	ret["Int"] = v.Int
	ret["Uint64"] = v.Uint64
	ret["Float32"] = v.Float32
	ret["Float64"] = v.Float64
	ret["Bytes"] = v.Bytes
	ret["String"] = v.String
	ret["Time"] = v.Time
	return ret, nil
}

func (v *testT) UnmarshalMap(data map[string]string) (err error) {
	for key, value := range data {
		switch key {
		case "Int":
			v.Int, err = strconv.ParseInt(value, 10, 64)
			if err != nil {
				return err
			}
		case "Uint64":
			v.Uint64, err = strconv.ParseUint(value, 10, 64)
			if err != nil {
				return err
			}
		case "Float32":
			f64, err := strconv.ParseFloat(value, 32)
			if err != nil {
				return err
			}
			v.Float32 = float32(f64)
		case "Float64":
			v.Float64, err = strconv.ParseFloat(value, 64)
			if err != nil {
				return err
			}
		case "Bytes":
			v.Bytes = []byte(value)
		case "String":
			v.String = value
		case "Time":
			v.Time, err = time.Parse(time.RFC3339, value)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type testTList []*testT

func (l *testTList) UnmarshalMap(data map[string]string) error {
	v := testT{}
	err := v.UnmarshalMap(data)
	if err != nil {
		return err
	}
	*l = append(*l, &v)
	return nil
}

func TestHashUpdateMap(t *testing.T) {

	rdb := newRedisClient()

	type args struct {
		ctx  context.Context
		rdb  *redis.Client
		key  string
		data map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		after   func(*args) error
	}{
		{
			name: "Normal",
			args: args{
				ctx: context.TODO(),
				rdb: rdb,
				key: "TestHashUpdateMap-Normal",
				data: map[string]interface{}{
					"int":    0,
					"uint64": uint64(1),
					"int8":   int8(-1),
					"bytes":  []byte{1, 2, 3, 4, 5},
					"string": "string filed",
					"time":   time.Now(),
				},
			},
			wantErr: false,
			after: func(args *args) error {
				defer func() {
					args.rdb.Del(args.ctx, args.key)
				}()
				if value, err := args.rdb.HGet(args.ctx, args.key, "int").Int(); err != nil {
					return err
				} else {
					if value != args.data["int"].(int) {
						return fmt.Errorf("HGet int want %v got %v", args.data["int"], value)
					}
				}
				if value, err := args.rdb.HGet(args.ctx, args.key, "uint64").Uint64(); err != nil {
					return err
				} else {
					if value != args.data["uint64"].(uint64) {
						return fmt.Errorf("HGet uint64 want %v got %v", args.data["uint64"], value)
					}
				}
				if value, err := args.rdb.HGet(args.ctx, args.key, "int8").Int(); err != nil {
					return err
				} else {
					if value != int(args.data["int8"].(int8)) {
						return fmt.Errorf("HGet int8 want %v got %v", args.data["int8"], value)
					}
				}
				if value, err := args.rdb.HGet(args.ctx, args.key, "bytes").Bytes(); err != nil {
					return err
				} else {
					if !bytes.Equal(value, args.data["bytes"].([]byte)) {
						return fmt.Errorf("HGet bytes want %v got %v", args.data["bytes"], value)
					}
				}
				if value, err := args.rdb.HGet(args.ctx, args.key, "string").Result(); err != nil {
					return err
				} else {
					if value != args.data["string"].(string) {
						return fmt.Errorf("HGet string want %v got %v", args.data["string"], value)
					}
				}
				if value, err := args.rdb.HGet(args.ctx, args.key, "time").Time(); err != nil {
					return err
				} else {
					if value.UTC() != args.data["time"].(time.Time).UTC() {
						return fmt.Errorf("HGet time want %v got %v", args.data["time"], value)
					}
				}
				return nil
			},
		},
		{
			name: "Expire",
			args: args{
				ctx: helper.NewExpireContext(context.TODO(), time.Second*3),
				rdb: rdb,
				key: "TestHashUpdateMap-Expire",
				data: map[string]interface{}{
					"int":    0,
					"uint64": uint64(1),
					"int8":   int8(-1),
					"bytes":  []byte{1, 2, 3, 4, 5},
					"string": "string filed",
					"time":   time.Now(),
				},
			},
			wantErr: false,
			after: func(args *args) error {
				defer func() {
					args.rdb.Del(args.ctx, args.key)
				}()
				if value, err := args.rdb.HGet(args.ctx, args.key, "int").Int(); err != nil {
					return err
				} else {
					if value != args.data["int"].(int) {
						return fmt.Errorf("HGet int want %v got %v", args.data["int"], value)
					}
				}
				if value, err := args.rdb.HGet(args.ctx, args.key, "uint64").Uint64(); err != nil {
					return err
				} else {
					if value != args.data["uint64"].(uint64) {
						return fmt.Errorf("HGet uint64 want %v got %v", args.data["uint64"], value)
					}
				}
				if value, err := args.rdb.HGet(args.ctx, args.key, "int8").Int(); err != nil {
					return err
				} else {
					if value != int(args.data["int8"].(int8)) {
						return fmt.Errorf("HGet int8 want %v got %v", args.data["int8"], value)
					}
				}
				if value, err := args.rdb.HGet(args.ctx, args.key, "bytes").Bytes(); err != nil {
					return err
				} else {
					if !bytes.Equal(value, args.data["bytes"].([]byte)) {
						return fmt.Errorf("HGet bytes want %v got %v", args.data["bytes"], value)
					}
				}
				if value, err := args.rdb.HGet(args.ctx, args.key, "string").Result(); err != nil {
					return err
				} else {
					if value != args.data["string"].(string) {
						return fmt.Errorf("HGet string want %v got %v", args.data["string"], value)
					}
				}
				if value, err := args.rdb.HGet(args.ctx, args.key, "time").Time(); err != nil {
					return err
				} else {
					if value.UTC() != args.data["time"].(time.Time).UTC() {
						return fmt.Errorf("HGet time want %v got %v", args.data["time"], value)
					}
				}
				time.Sleep(time.Second * 3)
				if value := args.rdb.Exists(args.ctx, args.key).Val(); value != 0 {
					return fmt.Errorf("Should not exist key %v", args.key)
				}
				return nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := helper.HashUpdateMap(tt.args.ctx, tt.args.rdb, tt.args.key, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("HashUpdateMap() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.after != nil {
				err := tt.after(&tt.args)
				if err != nil {
					t.Errorf("HashUpdateMap() error = %v", err)
				}
			}
		})
	}
}

func TestHashUpdate(t *testing.T) {

	rdb := newRedisClient()

	type args struct {
		ctx  context.Context
		rdb  *redis.Client
		key  string
		data helper.MapMarshaler
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		after   func(*args) error
	}{
		{
			name: "Normal",
			args: args{
				ctx: context.TODO(),
				rdb: rdb,
				key: "TestHashUpdate-Normal",
				data: &testT{
					Int:     -10,
					Uint64:  10,
					Float32: -1.23456,
					Float64: 5.12345678,
					Bytes:   []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
					String:  "TestHashUpdate-Normal",
					Time:    time.Now(),
				},
			},
			wantErr: false,
			after: func(args *args) error {
				defer func() {
					args.rdb.Del(args.ctx, args.key)
				}()
				if value, err := args.rdb.HGet(args.ctx, args.key, "Int").Int64(); err != nil {
					return err
				} else {
					if value != args.data.(*testT).Int {
						return fmt.Errorf("HGet Int want %v got %v", args.data.(*testT).Int, value)
					}
				}
				if value, err := args.rdb.HGet(args.ctx, args.key, "Uint64").Uint64(); err != nil {
					return err
				} else {
					if value != args.data.(*testT).Uint64 {
						return fmt.Errorf("HGet Uint64 want %v got %v", args.data.(*testT).Uint64, value)
					}
				}
				if value, err := args.rdb.HGet(args.ctx, args.key, "Float32").Float32(); err != nil {
					return err
				} else {
					if value != args.data.(*testT).Float32 {
						return fmt.Errorf("HGet Float32 want %v got %v", args.data.(*testT).Float32, value)
					}
				}
				if value, err := args.rdb.HGet(args.ctx, args.key, "Float64").Float64(); err != nil {
					return err
				} else {
					if value != args.data.(*testT).Float64 {
						return fmt.Errorf("HGet Float64 want %v got %v", args.data.(*testT).Float64, value)
					}
				}
				if value, err := args.rdb.HGet(args.ctx, args.key, "Bytes").Bytes(); err != nil {
					return err
				} else {
					if !bytes.Equal(value, args.data.(*testT).Bytes) {
						return fmt.Errorf("HGet bytes want %v got %v", args.data.(*testT).Bytes, value)
					}
				}
				if value, err := args.rdb.HGet(args.ctx, args.key, "String").Result(); err != nil {
					return err
				} else {
					if value != args.data.(*testT).String {
						return fmt.Errorf("HGet String want %v got %v", args.data.(*testT).String, value)
					}
				}
				if value, err := args.rdb.HGet(args.ctx, args.key, "Time").Time(); err != nil {
					return err
				} else {
					if value.UTC() != args.data.(*testT).Time.UTC() {
						return fmt.Errorf("HGet time want %v got %v", args.data.(*testT).Time.UTC(), value.UTC())
					}
				}
				return nil
			},
		},
		{
			name: "Expire",
			args: args{
				ctx: helper.NewExpireContext(context.TODO(), time.Second*3),
				rdb: rdb,
				key: "TestHashUpdate-Expire",
				data: &testT{
					Int:     -10,
					Uint64:  10,
					Float32: -1.23456,
					Float64: 5.12345678,
					Bytes:   []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
					String:  "TestHashUpdate-Normal",
					Time:    time.Now(),
				},
			},
			wantErr: false,
			after: func(args *args) error {
				defer func() {
					args.rdb.Del(args.ctx, args.key)
				}()
				if value, err := args.rdb.HGet(args.ctx, args.key, "Int").Int64(); err != nil {
					return err
				} else {
					if value != args.data.(*testT).Int {
						return fmt.Errorf("HGet Int want %v got %v", args.data.(*testT).Int, value)
					}
				}
				if value, err := args.rdb.HGet(args.ctx, args.key, "Uint64").Uint64(); err != nil {
					return err
				} else {
					if value != args.data.(*testT).Uint64 {
						return fmt.Errorf("HGet Uint64 want %v got %v", args.data.(*testT).Uint64, value)
					}
				}
				if value, err := args.rdb.HGet(args.ctx, args.key, "Float32").Float32(); err != nil {
					return err
				} else {
					if value != args.data.(*testT).Float32 {
						return fmt.Errorf("HGet Float32 want %v got %v", args.data.(*testT).Float32, value)
					}
				}
				if value, err := args.rdb.HGet(args.ctx, args.key, "Float64").Float64(); err != nil {
					return err
				} else {
					if value != args.data.(*testT).Float64 {
						return fmt.Errorf("HGet Float64 want %v got %v", args.data.(*testT).Float64, value)
					}
				}
				if value, err := args.rdb.HGet(args.ctx, args.key, "Bytes").Bytes(); err != nil {
					return err
				} else {
					if !bytes.Equal(value, args.data.(*testT).Bytes) {
						return fmt.Errorf("HGet bytes want %v got %v", args.data.(*testT).Bytes, value)
					}
				}
				if value, err := args.rdb.HGet(args.ctx, args.key, "String").Result(); err != nil {
					return err
				} else {
					if value != args.data.(*testT).String {
						return fmt.Errorf("HGet String want %v got %v", args.data.(*testT).String, value)
					}
				}
				if value, err := args.rdb.HGet(args.ctx, args.key, "Time").Time(); err != nil {
					return err
				} else {
					if value.UTC() != args.data.(*testT).Time.UTC() {
						return fmt.Errorf("HGet time want %v got %v", args.data.(*testT).Time.UTC(), value.UTC())
					}
				}
				time.Sleep(time.Second * 3)
				if value := args.rdb.Exists(args.ctx, args.key).Val(); value != 0 {
					return fmt.Errorf("Should not exist key %v", args.key)
				}
				return nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := helper.HashUpdate(tt.args.ctx, tt.args.rdb, tt.args.key, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("HashUpdate() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.after != nil {
				err := tt.after(&tt.args)
				if err != nil {
					t.Errorf("HashUpdate() error = %v", err)
				}
			}
		})
	}
}

func TestHashGetOne(t *testing.T) {

	rdb := newRedisClient()

	type args struct {
		ctx  context.Context
		rdb  *redis.Client
		key  string
		want helper.MapMarshaler
		got  helper.MapUnmarshaler
	}
	tests := []struct {
		name    string
		before  func(*args) error
		args    args
		after   func(*args) error
		wantErr bool
	}{
		{
			name:   "Nil",
			before: nil,
			args: args{
				ctx: context.TODO(),
				rdb: rdb,
				key: "TestHashGetOne-NotExist",
				got: nil,
			},
			wantErr: true,
		},
		{
			name: "Normal",
			before: func(args *args) error {
				args.want = &testT{
					Int:     -10,
					Uint64:  10,
					Float32: -1.23456,
					Float64: 5.12345678,
					Bytes:   []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
					String:  "TestHashUpdate-Normal",
					Time:    time.Now(),
				}
				return helper.HashUpdate(args.ctx, args.rdb, args.key, args.want)
			},
			args: args{
				ctx: context.TODO(),
				rdb: rdb,
				key: "TestHashGetOne-Normal",
				got: &testT{},
			},
			after: func(args *args) error {
				want, err := json.Marshal(args.want)
				if err != nil {
					return err
				}
				got, err := json.Marshal(args.got)
				if err != nil {
					return err
				}
				if !bytes.Equal(want, got) {
					return fmt.Errorf("HashGetOne()\n\twant %+v\n\tgot %+v", args.want, args.got)
				}
				return nil
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.before != nil {
				err := tt.before(&tt.args)
				if err != nil {
					t.Errorf("HashGetOne() before error = %v", err)
					return
				}
			}
			if err := helper.HashGetOne(tt.args.ctx, tt.args.rdb, tt.args.key, tt.args.got); (err != nil) != tt.wantErr {
				t.Errorf("HashGetOne() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.after != nil {
				err := tt.after(&tt.args)
				if err != nil {
					t.Errorf("HashGetOne() after error = %v", err)
					return
				}
			}
		})
	}
}

func TestHashGetList(t *testing.T) {

	rdb := newRedisClient()

	type args struct {
		ctx           context.Context
		rdb           *redis.Client
		n             int
		prefix        string
		want          testTList
		got           interface{}
		gotPagination *db.Pagination
		gotErr        error
		conditions    *db.Conditions
	}
	tests := []struct {
		name    string
		before  func(*args) error
		args    args
		after   func(*args) error
		loop    func(*args) bool
		wantErr bool
	}{
		{
			name:   "list nil",
			before: nil,
			args: args{
				ctx:        context.TODO(),
				rdb:        rdb,
				want:       nil,
				got:        nil,
				conditions: nil,
			},
			after: nil,
			loop: func(args *args) bool {
				// first
				if args.gotPagination == nil {
					return true
				}
				return false
			},
			wantErr: true,
		},
		{
			name:   "conditions nil",
			before: nil,
			args: args{
				ctx:        context.TODO(),
				rdb:        rdb,
				want:       nil,
				got:        make(testTList, 0),
				conditions: nil,
			},
			after: nil,
			loop: func(args *args) bool {
				// first
				if args.gotPagination == nil {
					return true
				}
				return false
			},
			wantErr: true,
		},
		{
			name:   "conditions Match is empty",
			before: nil,
			args: args{
				ctx:        context.TODO(),
				rdb:        rdb,
				want:       nil,
				got:        make(testTList, 0),
				conditions: db.NewConditions(),
			},
			after: nil,
			loop: func(args *args) bool {
				// first
				if args.gotPagination == nil {
					return true
				}
				return false
			},
			wantErr: true,
		},
		{
			name: "Normal",
			args: args{
				ctx:        context.TODO(),
				rdb:        rdb,
				n:          100,
				prefix:     "TestHashGetList-Normal",
				want:       make(testTList, 0),
				got:        &testTList{},
				conditions: db.NewConditions(db.Match("TestHashGetList-Normal-*")),
			},
			before: func(args *args) error {
				for i := 0; i < args.n; i++ {
					data := &testT{
						Int:     int64(i),
						Uint64:  uint64(i * 2),
						Float32: -1.23456 * float32(i),
						Float64: 5.12345678 * float64(i),
						Bytes:   []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
						String:  fmt.Sprintf("%s-%d", args.prefix, i),
						Time:    time.Now(),
					}
					err := helper.HashUpdate(args.ctx, args.rdb, fmt.Sprintf("%s-%d", args.prefix, i), data)
					if err != nil {
						return err
					}
					args.want = append(args.want, data)
				}
				return nil
			},
			after: func(args *args) (err error) {
				defer func() {
					for i := 0; i < args.n; i++ {
						if e := args.rdb.Del(args.ctx, fmt.Sprintf("%s-%d", args.prefix, i)).Err(); e != nil {
							err = e
						}
					}
				}()
				if args.gotPagination.Count < uint(args.conditions.Limit) {
					return fmt.Errorf("got pagination count = %d want >= %d", args.gotPagination.Count, args.conditions.Limit)
				}
				if len(*args.got.(*testTList)) != int(args.gotPagination.Count) {
					return fmt.Errorf("got pagination count = %d want = %d", args.gotPagination.Count, len(*args.got.(*testTList)))
				}
				return nil
			},
			loop: func(args *args) bool {
				// first
				if args.gotPagination == nil {
					return true
				}
				return false
			},
			wantErr: false,
		},
		{
			name: "Less-Limit",
			args: args{
				ctx:        context.TODO(),
				rdb:        rdb,
				n:          5,
				prefix:     "TestHashGetList-Less-Limit",
				want:       make(testTList, 0),
				got:        &testTList{},
				conditions: db.NewConditions(db.Match("TestHashGetList-Less-Limit-*")),
			},
			before: func(args *args) error {
				for i := 0; i < args.n; i++ {
					data := &testT{
						Int:     int64(i),
						Uint64:  uint64(i * 2),
						Float32: -1.23456 * float32(i),
						Float64: 5.12345678 * float64(i),
						Bytes:   []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
						String:  fmt.Sprintf("%s-%d", args.prefix, i),
						Time:    time.Now(),
					}
					err := helper.HashUpdate(args.ctx, args.rdb, fmt.Sprintf("%s-%d", args.prefix, i), data)
					if err != nil {
						return err
					}
					args.want = append(args.want, data)
				}
				return nil
			},
			after: func(args *args) (err error) {
				defer func() {
					for i := 0; i < args.n; i++ {
						if e := args.rdb.Del(args.ctx, fmt.Sprintf("%s-%d", args.prefix, i)).Err(); e != nil {
							err = e
						}
					}
				}()
				if args.gotPagination.Count != uint(args.n) {
					return fmt.Errorf("got pagination count = %d want = %d", args.gotPagination.Count, args.n)
				}
				if len(*args.got.(*testTList)) != int(args.gotPagination.Count) {
					return fmt.Errorf("got count = %d want = %d", args.gotPagination.Count, len(*args.got.(*testTList)))
				}
				return nil
			},
			loop: func(args *args) bool {
				// first
				if args.gotPagination == nil {
					return true
				}
				return false
			},
			wantErr: false,
		},
		{
			name: "Not exist",
			args: args{
				ctx:        context.TODO(),
				rdb:        rdb,
				n:          5,
				prefix:     "TestHashGetList-",
				want:       make(testTList, 0),
				got:        &testTList{},
				conditions: db.NewConditions(db.Match("TestHashGetList-Exist-*")),
			},
			before: func(args *args) error {
				for i := 0; i < args.n; i++ {
					data := &testT{
						Int:     int64(i),
						Uint64:  uint64(i * 2),
						Float32: -1.23456 * float32(i),
						Float64: 5.12345678 * float64(i),
						Bytes:   []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
						String:  fmt.Sprintf("%s-%d", args.prefix, i),
						Time:    time.Now(),
					}
					err := helper.HashUpdate(args.ctx, args.rdb, fmt.Sprintf("%s-%d", args.prefix, i), data)
					if err != nil {
						return err
					}
					args.want = append(args.want, data)
				}
				return nil
			},
			after: func(args *args) (err error) {
				defer func() {
					for i := 0; i < args.n; i++ {
						if e := args.rdb.Del(args.ctx, fmt.Sprintf("%s-%d", args.prefix, i)).Err(); e != nil {
							err = e
						}
					}
				}()
				if args.gotPagination.Count != 0 {
					return fmt.Errorf("got pagination count = %d want = %d", args.gotPagination.Count, 0)
				}
				if len(*args.got.(*testTList)) != int(args.gotPagination.Count) {
					return fmt.Errorf("got count = %d want = %d", args.gotPagination.Count, len(*args.got.(*testTList)))
				}
				return nil
			},
			loop: func(args *args) bool {
				// first
				if args.gotPagination == nil {
					return true
				}
				return false
			},
			wantErr: false,
		},
		{
			name: "All",
			args: args{
				ctx:        context.TODO(),
				rdb:        rdb,
				n:          5000,
				prefix:     "TestHashGetList-All",
				want:       make(testTList, 0),
				got:        &testTList{},
				conditions: db.NewConditions(db.Match("TestHashGetList-All-*"), db.Limit(5000)),
			},
			before: func(args *args) error {
				for i := 0; i < args.n; i++ {
					data := &testT{
						Int:     int64(i),
						Uint64:  uint64(i * 2),
						Float32: -1.23456 * float32(i),
						Float64: 5.12345678 * float64(i),
						Bytes:   []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
						String:  fmt.Sprintf("%s-%d", args.prefix, i),
						Time:    time.Now(),
					}
					err := helper.HashUpdate(args.ctx, args.rdb, fmt.Sprintf("%s-%d", args.prefix, i), data)
					if err != nil {
						return err
					}
					args.want = append(args.want, data)
				}
				return nil
			},
			after: func(args *args) (err error) {
				defer func() {
					for i := 0; i < args.n; i++ {
						if e := args.rdb.Del(args.ctx, fmt.Sprintf("%s-%d", args.prefix, i)).Err(); e != nil {
							err = e
						}
					}
				}()
				if args.gotPagination.Count < uint(args.conditions.Limit) {
					return fmt.Errorf("got pagination count = %d want = %d", args.gotPagination.Count, args.conditions.Limit)
				}
				if len(*args.got.(*testTList)) != int(args.gotPagination.Count) {
					return fmt.Errorf("got count = %d want = %d", args.gotPagination.Count, len(*args.got.(*testTList)))
				}
				return nil
			},
			loop: func(args *args) bool {
				// first
				if args.gotPagination == nil {
					return true
				}
				return false
			},
			wantErr: false,
		},
		{
			name: "Pagination",
			args: args{
				ctx:        context.TODO(),
				rdb:        rdb,
				n:          5000,
				prefix:     "TestHashGetList-Pagination",
				want:       make(testTList, 0),
				got:        &testTList{},
				conditions: db.NewConditions(db.Match("TestHashGetList-Pagination-*"), db.Limit(100)),
			},
			before: func(args *args) error {
				for i := 0; i < args.n; i++ {
					data := &testT{
						Int:     int64(i),
						Uint64:  uint64(i * 2),
						Float32: -1.23456 * float32(i),
						Float64: 5.12345678 * float64(i),
						Bytes:   []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
						String:  fmt.Sprintf("%s-%d", args.prefix, i),
						Time:    time.Now(),
					}
					err := helper.HashUpdate(args.ctx, args.rdb, fmt.Sprintf("%s-%d", args.prefix, i), data)
					if err != nil {
						return err
					}
					args.want = append(args.want, data)
				}
				return nil
			},
			after: func(args *args) (err error) {
				defer func() {
					for i := 0; i < args.n; i++ {
						if e := args.rdb.Del(args.ctx, fmt.Sprintf("%s-%d", args.prefix, i)).Err(); e != nil {
							err = e
						}
					}
				}()
				if len(*args.got.(*testTList)) != int(args.n) {
					return fmt.Errorf("got count = %d want = %d", len(*args.got.(*testTList)), args.n)
				}
				return nil
			},
			loop: func(args *args) bool {
				// first
				if args.gotPagination == nil {
					return true
				}
				if db.GetUint64Cursor(args.gotPagination.Cursor) > 0 {
					args.conditions.Cursor = args.gotPagination.Cursor
					return true
				}
				return false
			},
			wantErr: false,
		},
		{
			name: "Multiple Type",
			args: args{
				ctx:        context.TODO(),
				rdb:        rdb,
				n:          20,
				prefix:     "TestHashGetList-MultipleType",
				want:       make(testTList, 0),
				got:        &testTList{},
				conditions: db.NewConditions(db.Match("TestHashGetList-MultipleType-*")),
			},
			before: func(args *args) error {
				for i := 0; i < args.n; i++ {
					data := &testT{
						Int:     int64(i),
						Uint64:  uint64(i * 2),
						Float32: -1.23456 * float32(i),
						Float64: 5.12345678 * float64(i),
						Bytes:   []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
						String:  fmt.Sprintf("%s-%d", args.prefix, i),
						Time:    time.Now(),
					}
					err := helper.HashUpdate(args.ctx, args.rdb, fmt.Sprintf("%s-%d", args.prefix, i), data)
					if err != nil {
						return err
					}
					args.want = append(args.want, data)
				}
				args.rdb.Set(args.ctx, fmt.Sprintf("%s-%s", args.prefix, "string"), "StringKey", time.Second*10)
				return nil
			},
			after: func(args *args) (err error) {
				defer func() {
					for i := 0; i < args.n; i++ {
						if e := args.rdb.Del(args.ctx, fmt.Sprintf("%s-%d", args.prefix, i)).Err(); e != nil {
							err = e
						}
					}
				}()
				if len(args.gotErr.(helper.RedisErrors)) != 1 {
					return fmt.Errorf("got error count = %v want = %v", len(args.gotErr.(helper.RedisErrors)), 1)
				}
				if el, ok := args.gotErr.(helper.RedisErrors); ok {
					for _, item := range el {
						if e, ok := item.(*helper.RedisError); !ok {
							return fmt.Errorf("got type error = %v want = %v", reflect.TypeOf(e), reflect.TypeOf(&helper.RedisError{}))
						}
					}
				}
				if args.gotPagination.Count < uint(args.conditions.Limit) {
					return fmt.Errorf("got pagination count = %d want >= %d", args.gotPagination.Count, args.conditions.Limit)
				}
				if len(*args.got.(*testTList)) >= int(args.gotPagination.Count) {
					return fmt.Errorf("got pagination count = %d want >= %d", args.gotPagination.Count, len(*args.got.(*testTList)))
				}
				return nil
			},
			loop: func(args *args) bool {
				// first
				if args.gotPagination == nil {
					return true
				}
				return false
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.before != nil {
				err := tt.before(&tt.args)
				if err != nil {
					t.Errorf("HashGetList() before error = %v", err)
					return
				}
			}
			for tt.loop(&tt.args) {
				tt.args.gotPagination, tt.args.gotErr = helper.HashGetList(tt.args.ctx, tt.args.rdb, tt.args.got, tt.args.conditions)
				if (tt.args.gotErr != nil) != tt.wantErr {
					t.Errorf("HashGetList() error = %v, wantErr %v", tt.args.gotErr, tt.wantErr)
					return
				}
				if tt.args.gotErr != nil {
					break
				}
			}
			if tt.after != nil {
				err := tt.after(&tt.args)
				if err != nil {
					t.Errorf("HashGetList() after error = %v", err)
					return
				}
			}
		})
	}
}
