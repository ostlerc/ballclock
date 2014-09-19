package baller

//First implementation, slow
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

//Second implementation of ballclock using circular queue with no leaked memory
func EvalBallClockV2(input_size int8) int {
	q, e := newQueue(input_size)
	var min [4]int8
	var fmin [11]int8
	var hour [11]int8
	fminv := 0
	hourv := 0
	hours := 0

	//fmt.Println(q)
	for {
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

		//fmt.Println(q)

		if fminv == 11 {
			for i := 10; i >= 0; i-- {
				e = e.next
				e.v = fmin[i]
			}
			fminv = 0

			if hourv == 11 {
				for i := 10; i >= 0; i-- {
					e = e.next
					e.v = hour[i]
				}

				//fmt.Printf("hours %d %s\n", hours, q)
				if queueDone(q, input_size) {
					return (hours + 1) / 24
				}

				hourv = 0
				e = e.next
				e.v = v
			} else {
				hour[hourv] = v
				hourv++
			}

			hours++
		} else {
			fmin[fminv] = v
			fminv++
		}
	}

	return hours / 24
}

func EvalBallClockV6(input_size int8) int {
	var q [256]uint8
	var min1 uint8
	var min2 uint8
	var min3 uint8
	var min4 uint8
	var fmin [11]uint8
	var hour [11]uint8
	at := uint8(0)
	end := uint8(input_size - 1)
	fminv := 0
	hourv := 0
	hours := 0
	for i := at + 1; i <= end; i++ {
		q[i] = i
	}

	for {

		min1 = q[at]
		at++
		min2 = q[at]
		at++
		min3 = q[at]
		at++
		min4 = q[at]
		at++

		v := q[at]
		at++

		end++
		q[end] = min4
		end++
		q[end] = min3
		end++
		q[end] = min2
		end++
		q[end] = min1

		if fminv == 11 {
			end++
			q[end] = fmin[10]
			end++
			q[end] = fmin[9]
			end++
			q[end] = fmin[8]
			end++
			q[end] = fmin[7]
			end++
			q[end] = fmin[6]
			end++
			q[end] = fmin[5]
			end++
			q[end] = fmin[4]
			end++
			q[end] = fmin[3]
			end++
			q[end] = fmin[2]
			end++
			q[end] = fmin[1]
			end++
			q[end] = fmin[0]
			fminv = 0

			if hourv == 11 {
				for i := 10; i >= 0; i-- {
					end++
					q[end] = hour[i]
				}

				x := uint8(0)
				i := at
				for ; i != end; i++ {
					if q[i] != x {
						break
					}
					x++
				}
				if i == end {
					return (hours + 1) / 24
				}

				hourv = 0
				end++
				q[end] = v
			} else {
				hour[hourv] = v
				hourv++
			}

			hours++
		} else {
			fmin[fminv] = v
			fminv++
		}
	}

	panic("Couldn't finish!")
}
