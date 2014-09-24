package ballclock

type queue struct {
	v    int8
	next *queue
}

type uintqueue struct {
	v    uint8
	next *uintqueue
}

func queueDone(q *queue, size int8) bool {
	for i := int8(0); i < size; i++ {
		if q.v != i {
			return false
		}
		q = q.next
	}
	return true
}

func queueUintDone(q *uintqueue, size uint8) bool {
	for i := uint8(0); i < size; i++ {
		if q.v != i {
			return false
		}
		q = q.next
	}
	return true
}

func clockDone(a []int8) bool {
	for i := int8(len(a)) - 1; i > 0; i-- {
		if a[i] != i {
			return false
		}
	}
	return true
}

func clockDoneV3(a []int) bool {
	for i := len(a) - 1; i > 0; i-- {
		if a[i] != i {
			return false
		}
	}
	return true
}

func newQueue(size int8) (*queue, *queue) {
	front := &queue{}
	at := front
	for i := int8(1); i < size; i++ {
		at.next = &queue{v: i}
		at = at.next
	}

	at.next = front

	return front, at
}

func newUintQueue(size uint8) (*uintqueue, *uintqueue) {
	front := &uintqueue{}
	at := front
	for i := uint8(1); i < size; i++ {
		at.next = &uintqueue{v: i}
		at = at.next
	}

	at.next = front

	return front, at
}
