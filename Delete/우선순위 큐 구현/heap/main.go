package main

type node struct {
	priority, value int
}

func NewPriorityQueue() heap {
	return heap{}
}

type heap []node

func (h heap) IsEmpty() bool {
	return h.Len() == 0
}

func (h heap) Len() int {
	return len(h)
}

func (h *heap) Enqueue(priority, value int) {
	newNode := node{priority: priority, value: value}
	*h = append(*h, newNode)
	h.heapSort()
}

func (h *heap) Dequeue() node {
	if h.IsEmpty() {
		return node{-1, -1}
	}
	answer := (*h)[0]
	*h = (*h)[1:]
	return answer
}

func (h heap) Peek() node {
	if h.IsEmpty() {
		return node{-1, -1}
	}
	return h[0]
}

func (h *heap) heapSort() {
	for i := h.Len()/2 - 1; i >= 0; i-- {
		h.heapify(i, h.Len())
	}
	for i := h.Len() - 1; i >= 0; i-- {
		(*h)[0], (*h)[i] = (*h)[i], (*h)[0]
		h.heapify(0, i)
	}
}

func (h *heap) heapify(idx, heapSize int) {
	largest, leftIdx, rightIdx := idx, idx*2+1, idx*2+2
	if leftIdx < heapSize && (*h)[leftIdx].priority > (*h)[largest].priority {
		largest = leftIdx
	}
	if rightIdx < heapSize && (*h)[rightIdx].priority > (*h)[largest].priority {
		largest = rightIdx
	}
	if largest != idx {
		(*h)[idx], (*h)[largest] = (*h)[largest], (*h)[idx]
		h.heapify(largest, heapSize)
	}
}

func main() {

}
