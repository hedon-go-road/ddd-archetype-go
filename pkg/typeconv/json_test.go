package typeconv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMustFromJson(t *testing.T) {
	t.Run("integer should work", func(t *testing.T) {
		res := MustFromJson[int64]([]byte("1"))
		assert.Equal(t, int64(1), *res)
	})

	t.Run("string should work", func(t *testing.T) {
		res := MustFromJson[string]([]byte(`"hello"`))
		assert.Equal(t, "hello", *res)
	})

	t.Run("bool should work", func(t *testing.T) {
		res := MustFromJson[bool]([]byte("true"))
		assert.Equal(t, true, *res)
	})

	t.Run("float should work", func(t *testing.T) {
		res := MustFromJson[float64]([]byte("3.14"))
		assert.Equal(t, 3.14, *res)
	})

	t.Run("struct should work", func(t *testing.T) {
		type Info struct {
			Name string `json:"name"`
		}
		res := MustFromJson[Info]([]byte(`{"name": "hello"}`))
		assert.Equal(t, "hello", res.Name)
	})

	t.Run("map should work", func(t *testing.T) {
		res := MustFromJson[map[string]string]([]byte(`{"name": "hello"}`))
		assert.Equal(t, "hello", (*res)["name"])
	})

	t.Run("slice should work", func(t *testing.T) {
		res := MustFromJson[[]int64]([]byte(`[1,2,3,4]`))
		assert.Equal(t, int64(1), (*res)[0])
	})

	t.Run("invalid json should panic", func(t *testing.T) {
		assert.Panics(t, func() {
			_ = MustFromJson[int64]([]byte("hello"))
		})
	})
}

func TestFromJson(t *testing.T) {
	t.Run("struct should work", func(t *testing.T) {
		type Info struct {
			Name string `json:"name"`
		}
		res, err := FromJson[Info]([]byte(`{"name": "hello"}`))
		assert.Nil(t, err)
		assert.Equal(t, "hello", res.Name)
	})

	t.Run("invalid string should fail", func(t *testing.T) {
		res, err := FromJson[map[string]string]([]byte(`{"name": hello"}`))
		assert.NotNil(t, err)
		assert.Nil(t, res)
	})
}
