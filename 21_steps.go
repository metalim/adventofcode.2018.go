package main

import "time"

func step2() {
	var r0, r1, r2, r4, r5 int

	// L0:
	r4 = 123
L1:
	r4 = r4 & 456
	// L2:
	r4 = b2i(r4 == 72)
	// L3:
	if r4 != 0 {
		goto L5
	}
	// L4:
	goto L1
L5:
	r4 = 0
L6:
	r1 = r4 | 65536
	// L7:
	r4 = 678134
L8:
	r5 = r1 & 255
	// L9:
	r4 = r4 + r5
	// L10:
	r4 = r4 & 16777215
	// L11:
	r4 = r4 * 65899
	// L12:
	r4 = r4 & 16777215
	// L13:
	r5 = b2i(256 > r1)
	// L14:
	if r5 != 0 {
		goto L16
	}
	// L15:
	goto L17
L16:
	goto L28
L17:
	r5 = 0
L18:
	r2 = r5 + 1
	// L19:
	r2 = r2 * 256
	// L20:
	r2 = b2i(r2 > r1)
	// L21:
	if r2 != 0 {
		goto L23
	}
	// L22:
	goto L24
L23:
	goto L26
L24:
	r5 = r5 + 1
	// L25:
	goto L18
L26:
	r1 = r5
	// L27:
	goto L8
L28:
	r5 = b2i(r4 == r0)
	// L29:
	if r5 != 0 {
		return
	}
	// L30:
	goto L6
}

//////////////////////////////////////////

func step3() {
	var r0, r1, r2, r4, r5 int

	r4 = 123
L1:
	r4 = r4 & 456
	r4 = b2i(r4 == 72)
	if r4 != 0 {
		goto L5
	}
	goto L1
L5:
	r4 = 0
L6:
	r1 = r4 | 65536
	r4 = 678134
L8:
	r5 = r1 & 255
	r4 = r4 + r5
	r4 = r4 & 16777215
	r4 = r4 * 65899
	r4 = r4 & 16777215
	r5 = b2i(256 > r1)
	if r5 != 0 {
		goto L16
	}
	goto L17
L16:
	goto L28
L17:
	r5 = 0
L18:
	r2 = r5 + 1
	r2 = r2 * 256
	r2 = b2i(r2 > r1)
	if r2 != 0 {
		goto L23
	}
	goto L24
L23:
	goto L26
L24:
	r5 = r5 + 1
	goto L18
L26:
	r1 = r5
	goto L8
L28:
	r5 = b2i(r4 == r0)
	if r5 != 0 {
		return
	}
	goto L6
}

//////////////////////////////////////////

func step4() {
	var r0, r1, r2, r4, r5 int

	r4 = 123
	// L1:
	for {
		r4 = r4 & 456
		r4 = b2i(r4 == 72)
		if r4 != 0 {
			goto L5
		}
		// goto L1
	}

L5:
	r4 = 0

	// L6:
	for {
		r1 = r4 | 65536
		r4 = 678134
		// L8:
		for {
			r5 = r1 & 255
			r4 = r4 + r5
			r4 = r4 & 16777215
			r4 = r4 * 65899
			r4 = r4 & 16777215
			r5 = b2i(256 > r1)
			if r5 != 0 {
				goto L16
			}
			goto L17
		L16:
			goto L28
		L17:
			r5 = 0
			// L18:
			for {
				r2 = r5 + 1
				r2 = r2 * 256
				r2 = b2i(r2 > r1)
				if r2 != 0 {
					goto L23
				}
				goto L24
			L23:
				goto L26
			L24:
				r5 = r5 + 1
				// goto L18
			}
		L26:
			r1 = r5
			// goto L8
		}
	L28:
		r5 = b2i(r4 == r0)
		if r5 != 0 {
			return
		}
		// goto L6
	}
}

//////////////////////////////////////////

func step5() {
	var r0, r1, r2, r4, r5 int

	r4 = 123
	for {
		r4 = r4 & 456
		r4 = b2i(r4 == 72)
		if r4 != 0 {
			// goto L5
			break
		}
	}

	// L5:
	r4 = 0

	for {
		r1 = r4 | 65536
		r4 = 678134
		for {
			r5 = r1 & 255
			r4 = r4 + r5
			r4 = r4 & 16777215
			r4 = r4 * 65899
			r4 = r4 & 16777215
			r5 = b2i(256 > r1)
			if r5 != 0 {
				goto L16
			}
			goto L17
		L16:
			// goto L28
			break
		L17:
			r5 = 0
			for {
				r2 = r5 + 1
				r2 = r2 * 256
				r2 = b2i(r2 > r1)
				if r2 != 0 {
					goto L23
				}
				goto L24
			L23:
				// goto L26
				break
			L24:
				r5 = r5 + 1
			}
			// L26:
			r1 = r5
		}
		// L28:
		r5 = b2i(r4 == r0)
		if r5 != 0 {
			return
		}
	}
}

//////////////////////////////////////////

func step6() {
	var r0, r1, r2, r4, r5 int

	r4 = 123
	for {
		r4 = r4 & 456
		r4 = b2i(r4 == 72)
		if r4 != 0 {
			break
		}
	}

	r4 = 0

	for {
		r1 = r4 | 65536
		r4 = 678134
		for {
			r5 = r1 & 255
			r4 = r4 + r5
			r4 = r4 & 16777215
			r4 = r4 * 65899
			r4 = r4 & 16777215
			r5 = b2i(256 > r1)
			// if r5 != 0 {
			// 	goto L16
			// }
			if r5 == 0 {
				goto L17
			}
			// L16:
			break
		L17:
			r5 = 0
			for {
				r2 = r5 + 1
				r2 = r2 * 256
				r2 = b2i(r2 > r1)
				// if r2 != 0 {
				// 	goto L23
				// }
				if r2 == 0 {
					goto L24
				}
				// L23:
				break
			L24:
				r5 = r5 + 1
			}
			r1 = r5
		}
		r5 = b2i(r4 == r0)
		if r5 != 0 {
			return
		}
	}
}

//////////////////////////////////////////

func step7() {
	var r0, r1, r2, r4, r5 int

	r4 = 123
	for {
		r4 = r4 & 456
		r4 = b2i(r4 == 72)
		if r4 != 0 {
			break
		}
	}

	r4 = 0

	for {
		r1 = r4 | 65536
		r4 = 678134
		for {
			r5 = r1 & 255
			r4 = r4 + r5
			r4 = r4 & 16777215
			r4 = r4 * 65899
			r4 = r4 & 16777215
			r5 = b2i(256 > r1)
			// if r5 == 0 {
			// 	goto L17
			// }
			if r5 != 0 {
				break
			}
			// L17:
			r5 = 0
			for {
				r2 = r5 + 1
				r2 = r2 * 256
				r2 = b2i(r2 > r1)
				// if r2 == 0 {
				// 	goto L24
				// }
				if r2 != 0 {
					break
				}
				// L24:
				r5 = r5 + 1
			}
			r1 = r5
		}
		r5 = b2i(r4 == r0)
		if r5 != 0 {
			return
		}
	}
}

//////////////////////////////////////////

func step8() (r0 int) {
	var r1, r2, r4, r5 int

	r4 = 123
	for {
		r4 = r4 & 456
		r4 = b2i(r4 == 72)
		if r4 != 0 {
			break
		}
	}

	r4 = 0

	for {
		r1 = r4 | 65536
		r4 = 678134
		for {
			r5 = r1 & 255
			r4 = r4 + r5
			r4 = r4 & 16777215
			r4 = r4 * 65899
			r4 = r4 & 16777215
			r5 = b2i(256 > r1)
			if r5 != 0 {
				break
			}
			r5 = 0
			for {
				r2 = r5 + 1
				r2 = r2 * 256
				r2 = b2i(r2 > r1)
				if r2 != 0 {
					break
				}
				r5 = r5 + 1
			}
			r1 = r5
		}
		return r4
		r5 = b2i(r4 == r0)
		if r5 != 0 {
			return
		}
	}
}

//////////////////////////////////////////

func step9() (r0 int) {
	var r1, r2, r4, r5 int

	r4 = 123
	for {
		r4 = r4 & 456
		r4 = b2i(r4 == 72)
		if r4 != 0 {
			break
		}
	}

	r4 = 0

	for {
		r1 = r4 | 65536
		r4 = 678134
		for {
			r5 = r1 & 255
			r4 = r4 + r5
			r4 = r4 & 16777215
			r4 = r4 * 65899
			r4 = r4 & 16777215
			r5 = b2i(256 > r1)
			if r5 != 0 {
				break
			}

			for r5 = 0; ; r5++ {
				r2 = r5 + 1
				r2 = r2 * 256
				r2 = b2i(r2 > r1)
				if r2 != 0 {
					break
				}
			}
			r1 = r5
		}
		return r4
		_log(r4)
		r5 = b2i(r4 == r0)
		if r5 != 0 {
			return
		}
	}
}

//////////////////////////////////////////

const check = 10000

func step9() (r0 int) {
	var r1, r2, r4, r5, count int

	for {
		r1 = r4 | 65536
		r4 = 678134
		for {
			r5 = r1 & 255
			r4 = r4 + r5
			r4 = r4 & 16777215
			r4 = r4 * 65899
			r4 = r4 & 16777215
			r5 = b2i(256 > r1)
			if r5 != 0 {
				break
			}
			r5 = 0
			for {
				r2 = r5 + 1
				r2 = r2 * 256
				r2 = b2i(r2 > r1)
				if r2 != 0 {
					break
				}
				r5 = r5 + 1
			}
			r1 = r5
		}
		count++
		if count > check {
			_log("check", check, r1, r2, r4, r5)
			return r4
		}
		r5 = b2i(r4 == r0)
		if r5 != 0 {
			return
		}
	}
}

//////////////////////////////////////////

func step10() (r0 int) {
	var r1, r2, r4, r5, count int

	for {
		r1 = r4 | 65536
		r4 = 678134
		for {
			r4 = r4 + r1&255
			r4 = (r4 & 16777215 * 65899) & 16777215
			if 256 > r1 {
				break
			}
			// r5 = 0
			// for {
			// 	if (r5 + 1) * 256 > r1 {
			// 		break
			// 	}
			// 	r5 = r5 + 1
			// }
			r5 = r1 / 256
			r1 = r5
		}
		count++
		if count > check {
			_log("check", check, r1, r2, r4, r5)
			return r4
		}
		if r4 == r0 {
			return
		}
	}
}

//////////////////////////////////////////

func step11() (r0 int) {
	var r1, r2, r4, r5, count int

	for {
		r1 = r4 | 65536
		r4 = 678134
		for {
			r4 = r4 + r1&255
			r4 = (r4 & 16777215 * 65899) & 16777215
			if 256 > r1 {
				break
			}
			r1 = r1 / 256
		}
		count++
		if count > check {
			_log("check", check, r1, r2, r4, r5)
			return r4
		}
		if r4 == r0 {
			return
		}
	}
}

//////////////////////////////////////////

func step12() (r0 int) {
	var r1, r2, r4, r5, last, count int
	can := map[int]bool{}

	// ticker
	go func() {
		for {
			<-time.After(time.Second)
			_log("check", r1, r2, r4, r5, len(can), last, count)
		}
	}()

	for {
		r1 = r4 | 65536
		r4 = 678134
		for {
			r4 = r4 + r1&255
			r4 = (r4 & 16777215 * 65899) & 16777215
			if 256 > r1 {
				break
			}
			r1 = r1 / 256
		}
		count++
		if can[r4] {
			return last
		}
		can[r4] = true
		last = r4

		if r4 == r0 {
			return
		}
	}
}

//////////////////////////////////////////

func step13(init int) (r0 int) {
	var r1, r4, last int
	can := map[int]bool{}
	for {
		r1 = r4 | 65536
		for r4 = init; ; r1 /= 256 {
			r4 = ((r4 + r1&255) & 16777215 * 65899) & 16777215
			if r1 < 256 {
				break
			}
		}
		if can[r4] {
			_log(len(can))
			return last
		}
		can[r4] = true
		last = r4
	}
}
