package baller

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
