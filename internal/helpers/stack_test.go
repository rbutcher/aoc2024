package helpers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack_Push(t *testing.T) {
	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)

	assert.Equal(t, stack.Size(), 2)
}

func TestStack_Pop(t *testing.T) {
	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)

	i, ok := stack.Pop()
	assert.Equal(t, i, 2)
	assert.True(t, ok)

	i, ok = stack.Pop()
	assert.Equal(t, i, 1)
	assert.True(t, ok)

	i, ok = stack.Pop()
	assert.Equal(t, i, 0)
	assert.False(t, ok)
}

func TestStack_Peek(t *testing.T) {
	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)

	item, ok := stack.Peek()
	assert.Equal(t, item, 2)
	assert.True(t, ok)

	stack.Pop()
	item, ok = stack.Peek()
	assert.Equal(t, item, 1)
	assert.True(t, ok)

	stack.Pop()
	item, ok = stack.Peek()
	assert.Equal(t, item, 0)
	assert.False(t, ok)
}

func TestStack_IsEmpty(t *testing.T) {
	stack := NewStack[int]()

	assert.True(t, stack.IsEmpty())

	stack.Push(1)
	assert.False(t, stack.IsEmpty())

	stack.Pop()
	assert.True(t, stack.IsEmpty())
}

func TestStack_Size(t *testing.T) {
	stack := NewStack[int]()

	assert.Equal(t, stack.Size(), 0)

	stack.Push(1)
	stack.Push(2)
	assert.Equal(t, stack.Size(), 2)

	stack.Pop()
	assert.Equal(t, stack.Size(), 1)
}
