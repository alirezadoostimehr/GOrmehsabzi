package queue

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEmpty(t *testing.T) {
	var testQueue Queue
	assert.Equal(t, testQueue.IsEmpty(), true)
}

func TestFront(t *testing.T) {
	var testQueue Queue
	val, check := testQueue.Front()
	assert.Equal(t, struct {
		int
		bool
	}{val, check}, struct {
		int
		bool
	}{-1, false})

	testQueue.Push(1)
	val, check = testQueue.Front()
	assert.Equal(t, struct {
		int
		bool
	}{val, check}, struct {
		int
		bool
	}{1, true})

}

func TestPop(t *testing.T) {
	var teskQueue Queue

	val, check := teskQueue.Pop()
	assert.Equal(t, struct {
		int
		bool
	}{val, check}, struct {
		int
		bool
	}{-1, false})
}

func TestPushPop(t *testing.T) {
	var testQueue Queue

	ln := (int)(1e6)
	ls := make([]int, 0)
	for i := 0; i < ln; i++ {
		tmpInt := rand.Int()
		ls = append(ls, tmpInt)
		testQueue.Push(tmpInt)
	}

	for val, check := testQueue.Pop(); check; val, check = testQueue.Pop() {
		assert.Equal(t, val, ls[0])
		ls = ls[1:]
	}
}
