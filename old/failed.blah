package ballclock

//First implementation of ballclock using slices and append
func EvalBallClock(input_size int8) int {

	q := make([]int8, input_size, input_size)
	var min [4]int8
	minv := 0
	var fmin [11]int8
	fminv := 0
	var hour [11]int8
	hourv := 0

	for i := int8(0); i < input_size; i++ {
		q[i] = i
	}

	hours := 0
	for {
		v := q[minv]
		if minv == 4 {
			q = q[5:] //pop mins and 'v'
			q = append(q, min[3], min[2], min[1], min[0])
			minv = 0

			if fminv == 11 {
				q = append(q, fmin[10], fmin[9], fmin[8], fmin[7], fmin[6], fmin[5], fmin[4], fmin[3], fmin[2], fmin[1], fmin[0])
				fminv = 0

				if hourv == 11 {
					q = append(q, hour[10], hour[9], hour[8], hour[7], hour[6], hour[5], hour[4], hour[3], hour[2], hour[1], hour[0], v)
					hourv = 0
					if clockDone(q) {
						hours++
						break
					}
				} else {
					hour[hourv] = v
					hourv++
				}

				hours++
				//fmt.Printf("Hours=%d\n", hours)
			} else {
				fmin[fminv] = v
				fminv++
			}
		} else {
			min[minv] = v
			minv++
		}
	}

	return hours / 24
}

//This version is exactly like V1 but it uses int instead of int8
//The performance decline in this version is significant
func EvalBallClockV3(input_size int) int {

	queue := make([]int, input_size, input_size)
	var min [4]int
	minv := 0
	var fmin [11]int
	fminv := 0
	var hour [11]int
	hourv := 0

	for i := int(0); i < input_size; i++ {
		queue[i] = i
	}

	hours := 0
	for {
		v := queue[minv]
		//fmt.Printf("q %d minv %d fmin %d hour %d %d %d\n", v, minv, fminv, hourv, hours, hours/24)
		//fmt.Printf("queue %d minv %d fmin %d hour %d %d %d\n", len(queue), minv, len(fmin), len(hour), hours, hours/24)
		if minv == 4 {
			queue = queue[5:] //pop mins and 'v'
			queue = append(queue, min[3], min[2], min[1], min[0])
			minv = 0

			if fminv == 11 {
				queue = append(queue, fmin[10], fmin[9], fmin[8], fmin[7], fmin[6], fmin[5], fmin[4], fmin[3], fmin[2], fmin[1], fmin[0])
				fminv = 0

				if hourv == 11 {
					queue = append(queue, hour[10], hour[9], hour[8], hour[7], hour[6], hour[5], hour[4], hour[3], hour[2], hour[1], hour[0], v)
					hourv = 0
					if clockDoneV3(queue) {
						hours++
						break
					}
				} else {
					hour[hourv] = v
					hourv++
				}

				hours++
				//fmt.Printf("Hours=%d\n", hours)
			} else {
				fmin[fminv] = v
				fminv++
			}
		} else {
			min[minv] = v
			minv++
		}
	}

	return hours / 24
}

func EvalBallClockV4(input_size int8) int {
	q, e := newQueue(input_size)
	var min [4]int8
	var fmin [11]int8
	var hour [11]int8
	hourv := 0
	hours := 0

	for {
		for fminv := 0; fminv < 11; fminv++ {
			for minv := 0; minv < 4; minv++ {
				min[minv] = q.v
				q = q.next
			}
			fmin[fminv] = q.v
			q = q.next
			for i := 3; i >= 0; i-- {
				e = e.next
				e.v = min[i]
			}
		}

		hours++

		for minv := 0; minv < 4; minv++ {
			min[minv] = q.v
			q = q.next
		}
		v := q.v
		q = q.next
		for i := 3; i >= 0; i-- {
			e = e.next
			e.v = min[i]
		}

		for i := 10; i >= 0; i-- {
			e = e.next
			e.v = fmin[i]
		}

		if hourv == 11 {
			for i := 10; i >= 0; i-- {
				e = e.next
				e.v = hour[i]
			}

			if queueDone(q, input_size) {
				return hours / 24
			}

			hourv = 0
			e = e.next
			e.v = v
		} else {
			hour[hourv] = v
			hourv++
		}
	}

	return hours / 24
}

func EvalBallClockV5(input_size int8) int {
	q := newStaticQueue(input_size)
	var min int2array
	var fmin int4array
	var hour int4array
	fminv := int8(0)
	hourv := int8(0)
	hours := 0

	//fmt.Println(q)
	for {
		min.set(0, q.pop())
		min.set(1, q.pop())
		min.set(2, q.pop())
		min.set(3, q.pop())

		v := q.pop()

		q.push(min.at(3))
		q.push(min.at(2))
		q.push(min.at(1))
		q.push(min.at(0))

		//fmt.Println(q)

		if fminv == 11 {
			for i := int8(10); i >= 0; i-- {
				//could push lots and add once (unroll)
				q.push(fmin.at(i))
			}
			fminv = 0

			if hourv == 11 {
				for i := int8(10); i >= 0; i-- {
					q.push(hour.at(i))
				}

				//fmt.Printf("hours %d %s\n", hours, q)
				if q.done() {
					return (hours + 1) / 24
				}

				hourv = 0
				q.push(v)
			} else {
				hour.set(hourv, v)
				hourv++
			}

			hours++
		} else {
			fmin.set(fminv, v)
			fminv++
		}
	}

	return hours / 24
}
