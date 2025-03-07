package typeconv

import (
	"encoding/json"
	"fmt"
)

func FromJson[T any](bs []byte) (*T, error) {
	var t T
	if err := json.Unmarshal(bs, &t); err != nil {
		return nil, err
	}
	return &t, nil
}

func MustFromJson[T any](bs []byte) *T {
	var t T
	err := json.Unmarshal(bs, &t)
	if err != nil {
		panic(fmt.Sprintf("unmarshal json data to *T(type=%T) error: %v", t, err))
	}
	return &t
}
