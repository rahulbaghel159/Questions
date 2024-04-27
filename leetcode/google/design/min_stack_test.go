package design

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinStack(t *testing.T) {
	obj := Constructor1()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)

	assert.Equal(t, -3, obj.GetMin())

	obj.Pop()
	assert.Equal(t, 0, obj.Top())

	assert.Equal(t, -2, obj.GetMin())
}
