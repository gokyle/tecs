package chapter1

// This file contains gate specifications.

// A BasicGate is a two-input, single output gate.
type BasicGate func(bool, bool) bool

// NAND is the foundational gate upon which all other gates are build.
func NAND(a, b bool) bool {
	return !(a && b)
}

// NOT is a logic inverter.
func NOT(a bool) bool {
	return NAND(a, a)
}

// PureAND is a NAND-only implementation of an AND gate, which returns
// true iff both of its inputs are true.
func PureAND(a, b bool) bool {
	tmp := NAND(a, b)
	return NAND(tmp, tmp)
}

// AND is a gate which returns true iff both of its inputs are true.
func AND(a, b bool) bool {
	tmp := NAND(a, b)
	return NOT(tmp)
}

// OR is a gate which returns true if either of its inputs are true.
func OR(a, b bool) bool {
	return NAND(NOT(a), NOT(b))
}

// PureOR is a NAND-only implementation of an OR gate, which returns
// true if either of its inputs are true.
func PureOR(a, b bool) bool {
	return NAND(NAND(a, a), NAND(b, b))
}

// XOR is a gate which returns true if its inputs are different, and
// false if they are the same.
func XOR(a, b bool) bool {
	and1 := NAND(NOT(a), b)
	and2 := NAND(a, NOT(b))
	return NAND(and1, and2)
}

// PureXOR is a NAND-only implementation of a XOR gate, which returns
// true if its inputs are different, and false if they are the same.
func PureXOR(a, b bool) bool {
	and1 := NAND(NAND(a, a), b)
	and2 := NAND(a, NAND(b, b))
	return NAND(and1, and2)
}

// MUX is a multiplexer, which returns the value of input A if the
// selector input is false, and the value of input B if the selector pin
// is true.
func MUX(a, b, sel bool) bool {
	outa := AND(a, NOT(sel))
	outb := AND(b, sel)
	return OR(outa, outb)
}

// PureMUX is a NAND-only implementation of a multiplexer, which returns
// the value of input A if the selector input is false, and the value of
// input B if the selector pin is true.
func PureMUX(a, b, sel bool) bool {
	return NAND(NAND(a, NAND(sel, sel)), NAND(b, sel))
}

// DMUX is a gate which returns the value of the input on A if sel is
// false, or the value of the input on B if the sel is true.
func DMUX(in, sel bool) (a, b bool) {
	a = AND(in, NOT(sel))
	b = AND(in, sel)
	return
}

// PureDMUX is a NAND-only implementation of a DMUX gate, which returns
// the value of the input on A if sel is false, or the value of the input
// on B if the sel is true.
func PureDMUX(in, sel bool) (a, b bool) {
	a = NAND(NAND(in, NAND(sel, sel)), NAND(in, NAND(sel, sel)))
	b = NAND(NAND(in, sel), NAND(in, sel))
	return
}
