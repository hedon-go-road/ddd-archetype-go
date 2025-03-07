package typeconv

import (
	"fmt"

	"google.golang.org/protobuf/proto"
)

func FromProto[T any](bs []byte) (*T, error) {
	var t T
	msg, ok := any(&t).(proto.Message)
	if !ok {
		return nil, fmt.Errorf("type %T does not implement proto.Message", t)
	}
	err := proto.Unmarshal(bs, msg)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func MustFromProto[T any](bs []byte) *T {
	var t T
	msg, ok := any(&t).(proto.Message)
	if !ok {
		panic(fmt.Sprintf("type *%T does not implement proto.Message", t))
	}
	err := proto.Unmarshal(bs, msg)
	if err != nil {
		panic(fmt.Sprintf("unmarshal protobuf data to *T(type=%T) error: %v", t, err))
	}
	return &t
}
