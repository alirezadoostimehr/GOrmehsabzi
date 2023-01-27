package queue

type Queue []int

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) Front() (int, bool) {
	if (*q).IsEmpty() {
		return -1, false
	}
	return (*q)[0], true
}

func (q *Queue) Push(val int) {
	(*q) = append((*q), val)
}

func (q *Queue) Pop() (int, bool) {
	if (*q).IsEmpty() {
		return -1, false
	}
	result := (*q)[0]
	(*q) = (*q)[1:]
	return result, true
}
