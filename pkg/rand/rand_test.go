package rand

import (
	"reflect"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermFrom1(t *testing.T) {
	// n == 0, return empty slice
	assert.Equal(t, []int{}, PermFrom1(0))

	// n < 0, should panic
	assert.Panics(t, func() {
		PermFrom1(-1)
	})

	// n > 0, should rand n different numbers
	res := PermFrom1(5)
	slices.Sort(res)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, res)

	// n > 0, res should be random in most cases
	diff := false
	for i := 0; i < 10; i++ {
		if !reflect.DeepEqual(PermFrom1(10), PermFrom1(10)) {
			diff = true
			break
		}
	}
	assert.True(t, diff)
}

func TestUUIDV7(t *testing.T) {
	var pre string
	var cur string
	for i := 0; i < 100; i++ {
		cur = UUIDV7()
		assert.NotEqual(t, cur, pre)
		pre = cur
	}
}
