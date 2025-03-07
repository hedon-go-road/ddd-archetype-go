package collection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestManager(t *testing.T) {
	type Item struct {
		Name string
	}

	manager := New[int, *Item]()
	t.Run("the type of T should be ineterface or pointer", func(t *testing.T) {
		assert.Equal(t, 0, len(manager.items))
	})

	t.Run("no item, get should return nil", func(t *testing.T) {
		get := manager.Get(1)
		assert.Nil(t, get)
	})

	t.Run("no item, delete should return nil", func(t *testing.T) {
		d := manager.Delete(1)
		assert.Nil(t, d)
	})

	t.Run("get after add, should return the item", func(t *testing.T) {
		manager.Add(1, &Item{Name: "1"})
		item := manager.Get(1)
		assert.NotNil(t, item)
		assert.Equal(t, "1", item.Name)
		assert.Equal(t, 1, len(manager.items))
		assert.Equal(t, true, manager.Exists(1))
	})

	t.Run("delete after add, should return the item", func(t *testing.T) {
		d := manager.Delete(1)
		assert.NotNil(t, d)
		assert.Equal(t, "1", d.Name)
		assert.Equal(t, 0, len(manager.items))
		assert.Equal(t, false, manager.Exists(1))
	})

	t.Run("range return true should work", func(t *testing.T) {
		manager.Add(1, &Item{Name: "1"})
		manager.Add(2, &Item{Name: "2"})
		manager.Add(3, &Item{Name: "3"})
		var names []string
		manager.Range(func(_ int, v *Item) bool {
			names = append(names, v.Name)
			return true
		})
		assert.Equal(t, 3, len(names))
	})

	t.Run("range return false should work", func(t *testing.T) {
		manager.Add(1, &Item{Name: "1"})
		manager.Add(2, &Item{Name: "2"})
		manager.Add(3, &Item{Name: "3"})
		var names []string
		j := 0
		manager.Range(func(i int, v *Item) bool {
			names = append(names, v.Name)
			j++
			return j < 2
		})
		assert.Equal(t, 2, len(names))
	})
}
