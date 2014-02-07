package chapter1

func OR8(in [8]bool) bool {
	return OR(OR(OR(in[0], in[1]), OR(in[2], in[3])),
		OR(OR(in[4], in[5]), OR(in[6], in[7])))
}

func PureOR8(in [8]bool) bool {
	var s1a, s1b, s1c, s1d, s2a, s2b bool

	s1a = NAND(NAND(in[0], in[0]), NAND(in[1], in[1]))
	s1b = NAND(NAND(in[2], in[2]), NAND(in[3], in[3]))
	s1c = NAND(NAND(in[4], in[4]), NAND(in[5], in[5]))
	s1d = NAND(NAND(in[6], in[6]), NAND(in[7], in[7]))
	s2a = NAND(NAND(s1a, s1a), NAND(s1b, s1b))
	s2b = NAND(NAND(s1c, s1c), NAND(s1d, s1d))
	return NAND(NAND(s2a, s2a), NAND(s2b, s2b))
}

func MUX4W8(a, b, c, d [8]bool, sel [2]bool) (out [8]bool) {
	var mux0, mux1 [8]bool
	for i := 0; i < 8; i++ {
		mux0[i] = MUX(a[i], c[i], sel[0])
		mux1[i] = MUX(b[i], d[i], sel[0])
	}
	for i := 0; i < 8; i++ {
		out[i] = MUX(mux0[i], mux1[i], sel[1])
	}
	return
}

func PureMUX4W8(a, b, c, d [8]bool, sel [2]bool) (out [8]bool) {
	var mux0, mux1 [8]bool
	for i := 0; i < 8; i++ {
		mux0[i] = NAND(NAND(a[i], NAND(sel[0], sel[0])), NAND(c[i], sel[0]))
		mux1[i] = NAND(NAND(b[i], NAND(sel[0], sel[0])), NAND(d[i], sel[0]))
	}
	for i := 0; i < 8; i++ {
		out[i] = NAND(NAND(mux0[i], NAND(sel[1], sel[1])), NAND(mux1[i], sel[1]))
	}
	return
}
