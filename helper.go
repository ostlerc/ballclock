package ballclock

import "fmt"

type int2array struct {
	v0 int8
	v1 int8
	v2 int8
	v3 int8
}

func (a *int2array) at(i int8) int8 {
	switch i {
	case 0:
		return a.v0
	case 1:
		return a.v1
	case 2:
		return a.v2
	case 3:
		return a.v3
	default:
		panic("Not a valid number")
	}
}

func (a *int2array) set(i int8, v int8) {
	switch i {
	case 0:
		a.v0 = v
	case 1:
		a.v1 = v
	case 2:
		a.v2 = v
	case 3:
		a.v3 = v
	default:
		panic("out of range!")
	}
}

type int4array struct {
	v0  int8
	v1  int8
	v2  int8
	v3  int8
	v4  int8
	v5  int8
	v6  int8
	v7  int8
	v8  int8
	v9  int8
	v10 int8
	v11 int8
	v12 int8
	v13 int8
	v14 int8
	v15 int8
}

func (a *int4array) at(i int8) int8 {
	switch i {
	case 0:
		return a.v0
	case 1:
		return a.v1
	case 2:
		return a.v2
	case 3:
		return a.v3
	case 4:
		return a.v4
	case 5:
		return a.v5
	case 6:
		return a.v6
	case 7:
		return a.v7
	case 8:
		return a.v8
	case 9:
		return a.v9
	case 10:
		return a.v10
	case 11:
		return a.v11
	case 12:
		return a.v12
	case 13:
		return a.v13
	case 14:
		return a.v14
	case 15:
		return a.v15
	default:
		panic("Not a valid number")
	}
}

func (a *int4array) set(i int8, v int8) {
	switch i {
	case 0:
		a.v0 = v
	case 1:
		a.v1 = v
	case 2:
		a.v2 = v
	case 3:
		a.v3 = v
	case 4:
		a.v4 = v
	case 5:
		a.v5 = v
	case 6:
		a.v6 = v
	case 7:
		a.v7 = v
	case 8:
		a.v8 = v
	case 9:
		a.v9 = v
	case 10:
		a.v10 = v
	case 11:
		a.v11 = v
	case 12:
		a.v12 = v
	case 13:
		a.v13 = v
	case 14:
		a.v14 = v
	case 15:
		a.v15 = v
	default:
		panic("Not a valid number")
	}
}

type int7array struct {
	v0   int8
	v1   int8
	v2   int8
	v3   int8
	v4   int8
	v5   int8
	v6   int8
	v7   int8
	v8   int8
	v9   int8
	v10  int8
	v11  int8
	v12  int8
	v13  int8
	v14  int8
	v15  int8
	v16  int8
	v17  int8
	v18  int8
	v19  int8
	v20  int8
	v21  int8
	v22  int8
	v23  int8
	v24  int8
	v25  int8
	v26  int8
	v27  int8
	v28  int8
	v29  int8
	v30  int8
	v31  int8
	v32  int8
	v33  int8
	v34  int8
	v35  int8
	v36  int8
	v37  int8
	v38  int8
	v39  int8
	v40  int8
	v41  int8
	v42  int8
	v43  int8
	v44  int8
	v45  int8
	v46  int8
	v47  int8
	v48  int8
	v49  int8
	v50  int8
	v51  int8
	v52  int8
	v53  int8
	v54  int8
	v55  int8
	v56  int8
	v57  int8
	v58  int8
	v59  int8
	v60  int8
	v61  int8
	v62  int8
	v63  int8
	v64  int8
	v65  int8
	v66  int8
	v67  int8
	v68  int8
	v69  int8
	v70  int8
	v71  int8
	v72  int8
	v73  int8
	v74  int8
	v75  int8
	v76  int8
	v77  int8
	v78  int8
	v79  int8
	v80  int8
	v81  int8
	v82  int8
	v83  int8
	v84  int8
	v85  int8
	v86  int8
	v87  int8
	v88  int8
	v89  int8
	v90  int8
	v91  int8
	v92  int8
	v93  int8
	v94  int8
	v95  int8
	v96  int8
	v97  int8
	v98  int8
	v99  int8
	v100 int8
	v101 int8
	v102 int8
	v103 int8
	v104 int8
	v105 int8
	v106 int8
	v107 int8
	v108 int8
	v109 int8
	v110 int8
	v111 int8
	v112 int8
	v113 int8
	v114 int8
	v115 int8
	v116 int8
	v117 int8
	v118 int8
	v119 int8
	v120 int8
	v121 int8
	v122 int8
	v123 int8
	v124 int8
	v125 int8
	v126 int8
	v127 int8
}

func (a *int7array) at(i int8) int8 {
	switch i {
	case 0:
		return a.v0
	case 1:
		return a.v1
	case 2:
		return a.v2
	case 3:
		return a.v3
	case 4:
		return a.v4
	case 5:
		return a.v5
	case 6:
		return a.v6
	case 7:
		return a.v7
	case 8:
		return a.v8
	case 9:
		return a.v9
	case 10:
		return a.v10
	case 11:
		return a.v11
	case 12:
		return a.v12
	case 13:
		return a.v13
	case 14:
		return a.v14
	case 15:
		return a.v15
	case 16:
		return a.v16
	case 17:
		return a.v17
	case 18:
		return a.v18
	case 19:
		return a.v19
	case 20:
		return a.v20
	case 21:
		return a.v21
	case 22:
		return a.v22
	case 23:
		return a.v23
	case 24:
		return a.v24
	case 25:
		return a.v25
	case 26:
		return a.v26
	case 27:
		return a.v27
	case 28:
		return a.v28
	case 29:
		return a.v29
	case 30:
		return a.v30
	case 31:
		return a.v31
	case 32:
		return a.v32
	case 33:
		return a.v33
	case 34:
		return a.v34
	case 35:
		return a.v35
	case 36:
		return a.v36
	case 37:
		return a.v37
	case 38:
		return a.v38
	case 39:
		return a.v39
	case 40:
		return a.v40
	case 41:
		return a.v41
	case 42:
		return a.v42
	case 43:
		return a.v43
	case 44:
		return a.v44
	case 45:
		return a.v45
	case 46:
		return a.v46
	case 47:
		return a.v47
	case 48:
		return a.v48
	case 49:
		return a.v49
	case 50:
		return a.v50
	case 51:
		return a.v51
	case 52:
		return a.v52
	case 53:
		return a.v53
	case 54:
		return a.v54
	case 55:
		return a.v55
	case 56:
		return a.v56
	case 57:
		return a.v57
	case 58:
		return a.v58
	case 59:
		return a.v59
	case 60:
		return a.v60
	case 61:
		return a.v61
	case 62:
		return a.v62
	case 63:
		return a.v63
	case 64:
		return a.v64
	case 65:
		return a.v65
	case 66:
		return a.v66
	case 67:
		return a.v67
	case 68:
		return a.v68
	case 69:
		return a.v69
	case 70:
		return a.v70
	case 71:
		return a.v71
	case 72:
		return a.v72
	case 73:
		return a.v73
	case 74:
		return a.v74
	case 75:
		return a.v75
	case 76:
		return a.v76
	case 77:
		return a.v77
	case 78:
		return a.v78
	case 79:
		return a.v79
	case 80:
		return a.v80
	case 81:
		return a.v81
	case 82:
		return a.v82
	case 83:
		return a.v83
	case 84:
		return a.v84
	case 85:
		return a.v85
	case 86:
		return a.v86
	case 87:
		return a.v87
	case 88:
		return a.v88
	case 89:
		return a.v89
	case 90:
		return a.v90
	case 91:
		return a.v91
	case 92:
		return a.v92
	case 93:
		return a.v93
	case 94:
		return a.v94
	case 95:
		return a.v95
	case 96:
		return a.v96
	case 97:
		return a.v97
	case 98:
		return a.v98
	case 99:
		return a.v99
	case 100:
		return a.v100
	case 101:
		return a.v101
	case 102:
		return a.v102
	case 103:
		return a.v103
	case 104:
		return a.v104
	case 105:
		return a.v105
	case 106:
		return a.v106
	case 107:
		return a.v107
	case 108:
		return a.v108
	case 109:
		return a.v109
	case 110:
		return a.v110
	case 111:
		return a.v111
	case 112:
		return a.v112
	case 113:
		return a.v113
	case 114:
		return a.v114
	case 115:
		return a.v115
	case 116:
		return a.v116
	case 117:
		return a.v117
	case 118:
		return a.v118
	case 119:
		return a.v119
	case 120:
		return a.v120
	case 121:
		return a.v121
	case 122:
		return a.v122
	case 123:
		return a.v123
	case 124:
		return a.v124
	case 125:
		return a.v125
	case 126:
		return a.v126
	case 127:
		return a.v127
	default:
		panic("Not a valid number")
	}
}

func (a *int7array) set(i int8, v int8) {
	switch i {
	case 0:
		a.v0 = v
	case 1:
		a.v1 = v
	case 2:
		a.v2 = v
	case 3:
		a.v3 = v
	case 4:
		a.v4 = v
	case 5:
		a.v5 = v
	case 6:
		a.v6 = v
	case 7:
		a.v7 = v
	case 8:
		a.v8 = v
	case 9:
		a.v9 = v
	case 10:
		a.v10 = v
	case 11:
		a.v11 = v
	case 12:
		a.v12 = v
	case 13:
		a.v13 = v
	case 14:
		a.v14 = v
	case 15:
		a.v15 = v
	case 16:
		a.v16 = v
	case 17:
		a.v17 = v
	case 18:
		a.v18 = v
	case 19:
		a.v19 = v
	case 20:
		a.v20 = v
	case 21:
		a.v21 = v
	case 22:
		a.v22 = v
	case 23:
		a.v23 = v
	case 24:
		a.v24 = v
	case 25:
		a.v25 = v
	case 26:
		a.v26 = v
	case 27:
		a.v27 = v
	case 28:
		a.v28 = v
	case 29:
		a.v29 = v
	case 30:
		a.v30 = v
	case 31:
		a.v31 = v
	case 32:
		a.v32 = v
	case 33:
		a.v33 = v
	case 34:
		a.v34 = v
	case 35:
		a.v35 = v
	case 36:
		a.v36 = v
	case 37:
		a.v37 = v
	case 38:
		a.v38 = v
	case 39:
		a.v39 = v
	case 40:
		a.v40 = v
	case 41:
		a.v41 = v
	case 42:
		a.v42 = v
	case 43:
		a.v43 = v
	case 44:
		a.v44 = v
	case 45:
		a.v45 = v
	case 46:
		a.v46 = v
	case 47:
		a.v47 = v
	case 48:
		a.v48 = v
	case 49:
		a.v49 = v
	case 50:
		a.v50 = v
	case 51:
		a.v51 = v
	case 52:
		a.v52 = v
	case 53:
		a.v53 = v
	case 54:
		a.v54 = v
	case 55:
		a.v55 = v
	case 56:
		a.v56 = v
	case 57:
		a.v57 = v
	case 58:
		a.v58 = v
	case 59:
		a.v59 = v
	case 60:
		a.v60 = v
	case 61:
		a.v61 = v
	case 62:
		a.v62 = v
	case 63:
		a.v63 = v
	case 64:
		a.v64 = v
	case 65:
		a.v65 = v
	case 66:
		a.v66 = v
	case 67:
		a.v67 = v
	case 68:
		a.v68 = v
	case 69:
		a.v69 = v
	case 70:
		a.v70 = v
	case 71:
		a.v71 = v
	case 72:
		a.v72 = v
	case 73:
		a.v73 = v
	case 74:
		a.v74 = v
	case 75:
		a.v75 = v
	case 76:
		a.v76 = v
	case 77:
		a.v77 = v
	case 78:
		a.v78 = v
	case 79:
		a.v79 = v
	case 80:
		a.v80 = v
	case 81:
		a.v81 = v
	case 82:
		a.v82 = v
	case 83:
		a.v83 = v
	case 84:
		a.v84 = v
	case 85:
		a.v85 = v
	case 86:
		a.v86 = v
	case 87:
		a.v87 = v
	case 88:
		a.v88 = v
	case 89:
		a.v89 = v
	case 90:
		a.v90 = v
	case 91:
		a.v91 = v
	case 92:
		a.v92 = v
	case 93:
		a.v93 = v
	case 94:
		a.v94 = v
	case 95:
		a.v95 = v
	case 96:
		a.v96 = v
	case 97:
		a.v97 = v
	case 98:
		a.v98 = v
	case 99:
		a.v99 = v
	case 100:
		a.v100 = v
	case 101:
		a.v101 = v
	case 102:
		a.v102 = v
	case 103:
		a.v103 = v
	case 104:
		a.v104 = v
	case 105:
		a.v105 = v
	case 106:
		a.v106 = v
	case 107:
		a.v107 = v
	case 108:
		a.v108 = v
	case 109:
		a.v109 = v
	case 110:
		a.v110 = v
	case 111:
		a.v111 = v
	case 112:
		a.v112 = v
	case 113:
		a.v113 = v
	case 114:
		a.v114 = v
	case 115:
		a.v115 = v
	case 116:
		a.v116 = v
	case 117:
		a.v117 = v
	case 118:
		a.v118 = v
	case 119:
		a.v119 = v
	case 120:
		a.v120 = v
	case 121:
		a.v121 = v
	case 122:
		a.v122 = v
	case 123:
		a.v123 = v
	case 124:
		a.v124 = v
	case 125:
		a.v125 = v
	case 126:
		a.v126 = v
	case 127:
		a.v127 = v
	default:
		panic("Not a valid number")
	}
}

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

type staticqueue struct {
	front int8
	end   int8
	data  int7array
}

func newStaticQueue(size int8) *staticqueue {
	q := staticqueue{end: size - 1}
	for i := int8(1); i < size; i++ {
		q.data.set(i, i)
	}
	return &q
}

func (q *staticqueue) at() int8 {
	return q.data.at(q.front)
}

func (q *staticqueue) pop() int8 {
	v := q.at()
	switch q.front {
	case 127:
		q.front = 0
	default:
		q.front++
	}
	return v
}

func (q *staticqueue) push(v int8) {
	switch q.end {
	case 127:
		q.end = 0
	default:
		q.end++
	}
	q.data.set(q.end, v)
}

func (q *staticqueue) push11(v int8) {
	//TODO:
}

func (q *staticqueue) done() bool {
	tf := q.front
	i := int8(0)
	for ; tf != q.end; i++ {
		if q.data.at(tf) != i {
			return false
		}
		switch tf {
		case 127:
			tf = 0
		default:
			tf++
		}
	}
	if q.data.at(q.end) != i {
		return false
	}
	return true
}

func (q *queue) String() string {
	ret := "["
	t := q
	ret += fmt.Sprintf("%d ", t.v)
	t = t.next
	for t != q {
		ret += fmt.Sprintf("%d ", t.v)
		t = t.next
	}
	return ret + "]"
}

func (q *staticqueue) String() string {
	ret := fmt.Sprintf("%d %d [", q.front, q.end)
	for i := q.front; i != q.end; {
		ret += fmt.Sprintf("%d ", q.data.at(i))
		switch i {
		case 127:
			i = 0
		default:
			i++
		}
	}
	ret += fmt.Sprintf("%d ", q.data.at(q.end))
	return ret + "]"
}
