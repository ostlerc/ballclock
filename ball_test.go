package ballclock

import (
	"testing"
)

func TestQueue(t *testing.T) {
	n := int8(123)
	q, e := newQueue(n)
	if q == e {
		t.Fatal("empty queue!")
	}
	var i int8
	for ; q != e; i++ {
		if q == nil {
			t.Fatal("q is nil")
		}
		if q.v != i {
			t.Fatal(q)
		}
		q = q.next
	}

	if i != int8(n-1) {
		t.Fatal("size was not correct", i)
	}
}

func TestStaticQueue(t *testing.T) {
	n := int8(123)
	q := newStaticQueue(n)
	if q.end != 122 {
		t.Fatal("End not correct", q.end)
	}
	if q.front != 0 {
		t.Fatal("Front not correct", q.end)
	}
	for i := int8(0); i < int8(123); i++ {
		if q.data.at(i) != i {
			t.Fatal("incorrect data!")
		}
	}

	if !q.done() {
		t.Fatal("Done not correct")
	}

	q.push(4)
	if q.front != 0 || q.end != 123 {
		t.Fatal("Push did not work")
	}
	q.push(4)
	q.push(4)
	q.push(4)
	q.push(4)
	q.push(4)
	if q.end != 0 {
		t.Fatal("Push rollover did not work", q.end)
	}

	q = newStaticQueue(n)
	for i := int8(0); i < int8(123); i++ {
		if q.pop() != i || q.front != i+1 {
			t.Fatal("Pop failed")
		}
	}

	q.data.set(int8(4), int8(30))
	if q.data.at(int8(4)) != int8(30) {
		t.Fatal("set failed!")
	}
}

/*
func TestBall(t *testing.T) {
	n := EvalBallClock(123)
	if n != 108855 {
		t.Fatal(n)
	}
}
*/
func TestBallV2(t *testing.T) {
	n := EvalBallClockV2(123)
	if n != 108855 {
		t.Fatal(n)
	}
}

func TestBallV6(t *testing.T) {
	n := EvalBallClockV6(123)
	if n != 108855 {
		t.Fatal("result was: ", n)
	}
}

func BenchmarkIndex(b *testing.B) {
	at := uint8(0)
	var data [256]uint8
	set := uint8(14)
	for i := 0; i < b.N; i++ {
		at++
		data[at] = set
	}
}

func BenchmarkIndexRoll(b *testing.B) {
	at := uint8(0)
	var data [256]uint8
	set := uint8(14)
	for i := 0; i < b.N; i++ {
		at++
		data[at] = set
	}
}

func BenchmarkVarIndex(b *testing.B) {
	var data [4]uint8
	for i := 0; i < b.N; i++ {
		data[0] = 1
		data[1] = 2
		data[2] = 3
		data[3] = 4
	}
}

func BenchmarkVarNamed(b *testing.B) {
	var data1 uint8
	var data2 uint8
	var data3 uint8
	var data4 uint8
	for i := 0; i < b.N; i++ {
		data1 = 1
		data2 = 2
		data3 = 3
		data4 = 4
	}
	if data1+data2+data3+data4 == 70 {
	}
}

/*
func TestBallV3(t *testing.T) {
	n := EvalBallClockV3(123)
	if n != 108855 {
		t.Fatal(n)
	}
}

func TestBallV4(t *testing.T) {
	n := EvalBallClockV4(123)
	if n != 108855 {
		t.Fatal("result was: ", n)
	}
}

func TestBallV5(t *testing.T) {
	n := EvalBallClockV5(27)
	if n != 23 {
		t.Fatal("result was: ", n)
	}
}
*/
func BenchmarkClock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EvalBallClock(123)
	}
}

func BenchmarkClockV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EvalBallClockV2(123)
	}
}

func BenchmarkClockV6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EvalBallClockV6(123)
	}
}

func BenchmarkRollUint8(b *testing.B) {
	var v uint8
	for i := 0; i < b.N; i++ {
		v++
	}
}

func BenchmarkRollInt8(b *testing.B) {
	var v uint8
	for i := 0; i < b.N; i++ {
		v++
	}
}

func BenchmarkRollSwitch(b *testing.B) {
	var v int8
	for i := 0; i < b.N; i++ {
		switch v {
		case 127:
			v = 0
		default:
			v++
		}
	}
}

/*
func BenchmarkClockV3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EvalBallClockV3(123)
	}
}

func BenchmarkClockV4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EvalBallClockV4(123)
	}
}

func BenchmarkClockV5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EvalBallClockV5(123)
	}
}

func BenchmarkArrayGet(b *testing.B) {
	var a [128]int8
	for i := int8(0); i < int8(127); i++ {
		a[i] = i
	}
	for i := 0; i < b.N; i++ {
		for j := int8(0); j < int8(127); j++ {
			t := a[j]
			t++
		}
	}
}

func BenchmarkInt7Get(b *testing.B) {
	var a int7array
	for i := int8(0); i < int8(127); i++ {
		a.set(i, i)
	}
	for i := 0; i < b.N; i++ {
		for j := int8(0); j < int8(127); j++ {
			t := a.at(j)
			t++
		}
	}
}

func BenchmarkArraySet(b *testing.B) {
	var a [128]int8
	for i := 0; i < b.N; i++ {
		for j := int8(0); j < int8(127); j++ {
			a[j] = j
		}
	}
}

func BenchmarkInt7Set(b *testing.B) {
	var a int7array
	for i := 0; i < b.N; i++ {
		for j := int8(0); j < int8(127); j++ {
			a.set(j, j)
		}
	}
}*/
