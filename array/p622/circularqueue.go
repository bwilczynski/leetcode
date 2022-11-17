package p622

type MyCircularQueue struct {
	items           []int
	cap, len, front int
}

func NewMyCircularQueue(k int) MyCircularQueue {
	return MyCircularQueue{
		cap:   k,
		items: make([]int, k),
	}
}

func (q *MyCircularQueue) EnQueue(value int) bool {
	if q.len == q.cap {
		return false
	}
	q.items[(q.front+q.len)%q.cap] = value
	q.len++
	return true
}

func (q *MyCircularQueue) DeQueue() bool {
	if q.len == 0 {
		return false
	}
	q.front = (q.front + 1) % q.cap
	q.len--
	return true
}

func (q *MyCircularQueue) Front() int {
	if q.len == 0 {
		return -1
	}
	return q.items[q.front]
}

func (q *MyCircularQueue) Rear() int {
	if q.len == 0 {
		return -1
	}
	return q.items[(q.front+q.len-1)%q.cap]
}

func (q *MyCircularQueue) IsEmpty() bool {
	return q.len == 0
}

func (q *MyCircularQueue) IsFull() bool {
	return q.len == q.cap
}
