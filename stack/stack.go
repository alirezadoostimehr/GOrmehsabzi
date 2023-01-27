package stack

type Stack []int

func (st *Stack) IsEmpty() bool {
	return len(*st) == 0
}

func (st *Stack) Top() (int, bool) {
	if (*st).IsEmpty() {
		return -1, false
	}
	return (*st)[0], true
}

func (st *Stack) Pop() (int, bool) {
	if (*st).IsEmpty() {
		return -1, false
	}
	result := (*st)[0]
	(*st) = (*st)[1:]
	return result, true
}

func (st *Stack) Push(val int) {
	(*st) = append([]int{val}, (*st)...)
}
