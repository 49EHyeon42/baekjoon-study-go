package main

type node struct {
	value, priority int
}

func newNode(value, priority int) *node {
	return &node{value: value, priority: priority}
}

type heap []int

func NewPriorityQueue() *heap {
	return &heap{}
}

func (h heap) Len() int {
	return len(h)
}
func (h heap) IsEmpty() bool {
	return h.Len() == 0
}

func (h *heap) peek() {
}

func main() {

}

/*
peek : 최대 우선순위 값 반환
enqueue : 삽입
dequeue : 최대 우선순위 요소를 삭제하고 그 값을 반환
heapify : 힙 속성을 유지하는 작업
*/
