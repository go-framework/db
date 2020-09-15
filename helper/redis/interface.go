package redis

type MapMarshaler interface {
	MarshalMap() (map[string]interface{}, error)
}

type MapUnmarshaler interface {
	UnmarshalMap(map[string]string) error
}
