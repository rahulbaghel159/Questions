package design

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLRUCache(t *testing.T) {
	obj := Constructor(2)

	obj.Put(1, 1)
	obj.Put(2, 2)
	assert.Equal(t, 1, obj.Get(1))

	obj.Put(3, 3)
	assert.Equal(t, -1, obj.Get(2))

	obj.Put(4, 4)
	assert.Equal(t, -1, obj.Get(1))

	assert.Equal(t, 3, obj.Get(3))

	assert.Equal(t, 4, obj.Get(4))
}
