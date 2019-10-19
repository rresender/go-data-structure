package lang

import "fmt"

type Heap struct {
	capacity int
	size     int
	items    []int
}

func (h *Heap) getLeftChildIndex(parentIndex int) int {
	return 2*parentIndex + 1
}

func (h *Heap) getRightChildIndex(parentIndex int) int {
	return 2*parentIndex + 2
}

func (h *Heap) getParentIndex(childIndex int) int {
	return (childIndex - 1) / 2
}

func (h *Heap) hasLeftChild(index int) bool {
	return h.getLeftChildIndex(index) < h.size
}

func (h *Heap) hasRightChild(index int) bool {
	return h.getRightChildIndex(index) < h.size
}

func (h *Heap) hasParent(index int) bool {
	return h.getParentIndex(index) >= 0
}

func (h *Heap) leftChild(index int) int {
	return h.items[h.getLeftChildIndex(index)]
}

func (h *Heap) rightChild(index int) int {
	return h.items[h.getRightChildIndex(index)]
}

func (h *Heap) parent(index int) int {
	return h.items[h.getParentIndex(index)]
}

func (h *Heap) swap(indexOne, indexTwo int) {
	temp := h.items[indexOne]
	h.items[indexOne] = h.items[indexTwo]
	h.items[indexTwo] = temp
}

func (h *Heap) ensureExtraCapacity() {
	if h.size == h.capacity {
		temp := make([]int, h.capacity*2)
		copy(temp, h.items)
		h.items = temp
		h.capacity *= 2
	}
}

func (h *Heap) heapfyUp() {
	index := h.size - 1
	for h.hasParent(index) && (h.parent(index) > h.items[index]) {
		h.swap(h.getParentIndex(index), index)
		index = h.getParentIndex(index)
	}
}

func (h *Heap) heapfyDown() {
	index := 0
	for h.hasLeftChild(index) {
		smallerChildIndex := h.getLeftChildIndex(index)
		if h.hasRightChild(index) && (h.rightChild(index) < h.leftChild(index)) {
			smallerChildIndex = h.getRightChildIndex(index)
		}

		if h.items[index] < h.items[smallerChildIndex] {
			break
		} else {
			h.swap(index, smallerChildIndex)
		}
		index = smallerChildIndex
	}
}

func (h *Heap) Peek() (int, error) {
	if h.size == 0 {
		return -1, fmt.Errorf("queue is empty")
	}
	return h.items[0], nil
}

func (h *Heap) Pool() (int, error) {
	if h.size == 0 {
		return -1, fmt.Errorf("queue is empty")
	}
	item := h.items[0]
	h.items[0] = h.items[h.size-1]
	h.items[h.size-1] = 0
	h.size--
	h.heapfyDown()
	return item, nil
}

func (h *Heap) Add(item int) {
	h.ensureExtraCapacity()
	h.items[h.size] = item
	h.size++
	h.heapfyUp()
}

func (h *Heap) Print() {
	fmt.Printf("%v\n", h.items)
}

func NewHeap(capacity int) *Heap {
	var h Heap
	h.capacity = capacity
	h.items = make([]int, capacity)
	return &h
}
