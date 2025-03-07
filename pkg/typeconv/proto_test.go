package typeconv

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"

	"github.com/hedon-go-road/ddd-archetype-go/fixtures"
)

func TestFromProto(t *testing.T) {
	t.Run("nil slice should success", func(t *testing.T) {
		data, err := FromProto[fixtures.Request](nil)
		assert.Nil(t, err)
		assert.NotNil(t, data)
	})

	t.Run("empty bytes should success", func(t *testing.T) {
		data, err := FromProto[fixtures.Request]([]byte{})
		assert.Nil(t, err)
		assert.NotNil(t, data)
	})

	t.Run("proto struct should work", func(t *testing.T) {
		req := fixtures.Request{
			Name: "hedon",
		}
		bs, err := proto.Marshal(&req)
		assert.Nil(t, err)
		data, err := FromProto[fixtures.Request](bs)
		assert.Nil(t, err)
		assert.Equal(t, "hedon", data.Name)
	})

	t.Run("invalid proto should fail", func(t *testing.T) {
		data, err := FromProto[fixtures.Request]([]byte("hello"))
		assert.NotNil(t, err)
		assert.Nil(t, data)
	})

	t.Run("invalid proto should fail", func(t *testing.T) {
		// An even number of Xs is a success, and an odd number of Xs is a failure
		// Hhhh, interesting 😄
		// https://github.com/golang/protobuf/issues/1641
		//
		// x -> field 15 | type varint
		// x -> 120
		data, err := FromProto[fixtures.Request]([]byte("xxxxx"))
		assert.NotNil(t, err)
		assert.Nil(t, data)
	})

	t.Run("xxx1 should success, x1 would override xx", func(t *testing.T) {
		data, err := FromProto[fixtures.TestMsg]([]byte("xxx1"))
		assert.Nil(t, err)
		assert.Equal(t, int32(49), data.Id)
	})

	t.Run("basic type should failed", func(t *testing.T) {
		data, err := FromProto[int64]([]byte("1"))
		assert.NotNil(t, err)
		assert.Nil(t, data)
	})
}

//nolint:dupl
func TestMustFromProto(t *testing.T) {
	req := fixtures.Request{
		Name: "hedon",
	}
	bs, err := proto.Marshal(&req)
	assert.Nil(t, err)

	t.Run("proto struct should work", func(t *testing.T) {
		data := MustFromProto[fixtures.Request](bs)
		assert.Equal(t, "hedon", data.Name)
	})

	t.Run("invalid proto should panic", func(t *testing.T) {
		assert.Panics(t, func() {
			_ = MustFromProto[fixtures.Request]([]byte("hello"))
		})
	})

	t.Run("invalid type should panic", func(t *testing.T) {
		assert.Panics(t, func() {
			_ = MustFromProto[int64](bs)
		})
	})
}
