package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLRUCache(t *testing.T) {
	lru := Constructor(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	assert.Equal(t, 1, (lru.Get(1)))
	lru.Put(3, 3)
	assert.Equal(t, -1, (lru.Get(2)))
}
