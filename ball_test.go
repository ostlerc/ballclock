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
