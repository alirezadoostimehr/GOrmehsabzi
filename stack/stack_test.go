package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEmpty(t *testing.T) {
	var testStack Stack
	assert.Equal(t, testStack.IsEmpty(), true)
}

func TestPushIsEmpty(t *testing.T) {
	var testStack Stack
	testStack.Push(10)
	assert.Equal(t, testStack.IsEmpty(), false)
}

func TestPop(t *testing.T) {
	var testStack Stack
	val, check := testStack.Pop()
	assert.Equal(t, struct {
		int
		bool
	}{val, check}, struct {
		int
		bool
	}{-1, false})
}

func TestPopPushTop(t *testing.T) {
	var testStack Stack

	testStack.Push(1)
	val, check := testStack.Pop()
	assert.Equal(t, struct {
		int
		bool
	}{val, check}, struct {
		int
		bool
	}{1, true})

	testStack.Push(2)
	testStack.Push(3)
	val, check = testStack.Top()
	assert.Equal(t, struct {
		int
		bool
	}{val, check}, struct {
		int
		bool
	}{3, true})

	val, check = testStack.Pop()
	assert.Equal(t, struct {
		int
		bool
	}{val, check}, struct {
		int
		bool
	}{3, true})

	val, check = testStack.Pop()
	assert.Equal(t, struct {
		int
		bool
	}{val, check}, struct {
		int
		bool
	}{2, true})

	val, check = testStack.Pop()
	assert.Equal(t, struct {
		int
		bool
	}{val, check}, struct {
		int
		bool
	}{-1, false})

}
