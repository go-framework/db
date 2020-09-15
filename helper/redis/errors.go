package redis

import (
	"fmt"
	"strings"
)

type NotExistError struct {
	Key string `json:"key"`
}

func (e NotExistError) Error() string {
	return fmt.Sprintf("key: %s is not exist", e.Key)
}

type RedisError struct {
	Key     string `json:"key"`
	Operate string `json:"operate"`
	Err     string `json:"error"`
}

func (e RedisError) Error() string {
	return fmt.Sprintf("%s key: %s error: %v", e.Operate, e.Key, e.Err)
}

type RedisErrors []error

func (l RedisErrors) Error() string {
	buf := strings.Builder{}
	buf.WriteByte('[')
	for idx, e := range l {
		buf.WriteString(e.Error())
		if idx < len(l)-1 {
			buf.WriteByte(',')
		}
	}
	buf.WriteByte(']')
	return buf.String()
}

func (l RedisErrors) Nil() error {
	if len(l) == 0 {
		return nil
	}
	return l
}
