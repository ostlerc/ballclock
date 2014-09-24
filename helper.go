package ballclock

type queue struct {
	v    int8
	next *queue
}

type uintqueue struct {
	v    uint8
	next *uintqueue
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
