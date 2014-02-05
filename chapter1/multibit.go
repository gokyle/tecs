package chapter1

// This file contains multibit gate specifications.

// NOT16 is a 16-bit variant of the NOT gate.
func NOT16(a [16]bool) (out [16]bool) {
	for i := range a {
		out[i] = NAND(a[i], a[i])
	}
	return
}

// AND16 is a 16-bit variant of the AND gate.
func AND16(a, b [16]bool) (out [16]bool) {
	for i := range a {
		out[i] = AND(a[i], b[i])
	}
	return
}

// PureAND16 is a NAND-only implementation of a 16-bit variant of the
// AND gate.
func PureAND16(a, b [16]bool) (out [16]bool) {
	for i := range a {
		out[i] = NAND(NAND(a[i], b[i]), NAND(a[i], b[i]))
	}
	return
}

// OR16 is a 16-bit variant of the OR gate.
func OR16(a, b [16]bool) (out [16]bool) {
	for i := range a {
		out[i] = NAND(NOT(a[i]), NOT(b[i]))
	}
	return
}

// PureOR16 is a NAND-only implementation of a 16-bit variant of the
// OR gate.
func PureOR16(a, b [16]bool) (out [16]bool) {
	for i := range out {
		out[i] = NAND(NAND(a[i], a[i]), NAND(b[i], b[i]))
	}
	return
}

// MUX16 is a 16-bit variant of the MUX gate.
func MUX16(a, b [16]bool, sel bool) (out [16]bool) {
	for i := range out {
		out[i] = OR(AND(a[i], NOT(sel)), AND(b[i], sel))
	}
	return
}

// PureMUX16 is a NAND-only implementation of a 16-bit variant of the
// MUX gate.
func PureMUX16(a, b [16]bool, sel bool) (out [16]bool) {
	for i := range out {
		out[i] = NAND(
			NAND(
				NAND(NAND(a[i], NAND(sel, sel)), NAND(a[i], NAND(sel, sel))),
				NAND(NAND(a[i], NAND(sel, sel)), NAND(a[i], NAND(sel, sel)))),
			NAND(
				NAND(NAND(b[i], sel), NAND(b[i], sel)),
				NAND(NAND(b[i], sel), NAND(b[i], sel))))

	}
	return
}
