package dinnerplatestacks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDinnerPlates(t *testing.T) {
	obj := Constructor(2)

	obj.Push(1)
	obj.Push(2)
	obj.Push(3)
	obj.Push(4)
	obj.Push(5)
	obj.PrintStack()
	obj.PrintLast()
	assert.Equal(t, obj.PopAtStack(0), 2, "")
	obj.Push(20)
	obj.Push(21)
	obj.PrintStack()
	obj.PrintLast()
	assert.Equal(t, obj.PopAtStack(0), 20, "")
	assert.Equal(t, obj.PopAtStack(2), 21, "")
	obj.PrintStack()
	assert.Equal(t, obj.Pop(), 5, "")
	assert.Equal(t, obj.Pop(), 4, "")
	assert.Equal(t, obj.Pop(), 3, "")
	assert.Equal(t, obj.Pop(), 1, "")
	assert.Equal(t, obj.Pop(), -1, "")
}

func TestDinnerPlates2(t *testing.T) {
	obj := Constructor(2)

	obj.Push(1)
	obj.Push(2)
	obj.Push(3)
	obj.Push(4)
	obj.Push(5)
	obj.Push(6)
	obj.Push(7)
	obj.PrintStack()
	assert.Equal(t, obj.PopAtStack(2), 6, "")
	assert.Equal(t, obj.PopAtStack(2), 5, "")
	assert.Equal(t, obj.PopAtStack(1), 4, "")
	assert.Equal(t, obj.PopAtStack(1), 3, "")
	assert.Equal(t, obj.PopAtStack(0), 2, "")
	obj.PrintStack()
	obj.Push(8)
	obj.Push(9)
	obj.PrintStack()
	assert.Equal(t, obj.Pop(), 7, "")
	assert.Equal(t, obj.Pop(), 9, "")
	assert.Equal(t, obj.Pop(), 8, "")
	assert.Equal(t, obj.Pop(), 1, "")
	assert.Equal(t, obj.Pop(), -1, "")
}
