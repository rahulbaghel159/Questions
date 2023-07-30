package lrucache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lruCache(t *testing.T) {
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

func Test_lruCache1(t *testing.T) {
	obj := Constructor(1)
	obj.Put(2, 1)

	assert.Equal(t, 1, obj.Get(2))

	obj.Put(3, 2)

	// assert.Equal(t, -1, obj.Get(2))

	// assert.Equal(t, 2, obj.Get(3))
}

func Test_lruCache2(t *testing.T) {
	obj := Constructor(2)

	assert.Equal(t, -1, obj.Get(2))

	obj.Put(2, 6)

	assert.Equal(t, -1, obj.Get(1))

	obj.Put(1, 5)
	obj.Put(1, 2)

	assert.Equal(t, 2, obj.Get(1))

	assert.Equal(t, 6, obj.Get(2))
}
