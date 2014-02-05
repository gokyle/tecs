package chapter1

import "testing"

type basicGateTest struct {
	A, B bool // inputs
	Out  bool // expected output
}

func newBasicSuite(out1, out2, out3, out4 bool) []basicGateTest {
	return []basicGateTest{
		basicGateTest{false, false, out1},
		basicGateTest{false, true, out2},
		basicGateTest{true, false, out3},
		basicGateTest{true, true, out4},
	}

}

func testBasicGate(t *testing.T, name string, gate BasicGate, tests []basicGateTest) {
	for _, tc := range tests {
		if gate(tc.A, tc.B) != tc.Out {
			t.Fatalf("expected %s(%v, %v) -> %v", name, tc.A, tc.B, tc.Out)
		}
	}
}

func TestNAND(t *testing.T) {
	tests := newBasicSuite(true, true, true, false)
	testBasicGate(t, "NAND", NAND, tests)
}

func TestNOT(t *testing.T) {
	if NOT(true) {
		t.Fatal("expected NOT(true) -> false")
	} else if !(NOT(false)) {
		t.Fatal("expected NOT(false) -> true")
	}
}

func TestAND(t *testing.T) {
	tests := newBasicSuite(false, false, false, true)
	testBasicGate(t, "ANDP", PureAND, tests)
}

func TestOR(t *testing.T) {
	tests := newBasicSuite(false, true, true, true)
	testBasicGate(t, "OR", OR, tests)
	testBasicGate(t, "ORP", PureOR, tests)
}

func TestXOR(t *testing.T) {
	tests := newBasicSuite(false, true, true, false)
	testBasicGate(t, "XOR", XOR, tests)
	testBasicGate(t, "XORP", PureXOR, tests)
}

func TestMUX(t *testing.T) {
	type MuxTest struct {
		A, B, Sel bool
		Out       bool
	}
	tests := []MuxTest{
		MuxTest{false, false, false, false},
		MuxTest{false, true, false, false},
		MuxTest{true, false, false, true},
		MuxTest{true, false, false, true},
		MuxTest{false, false, true, false},
		MuxTest{false, true, true, true},
		MuxTest{true, false, true, false},
		MuxTest{true, true, true, true},
	}
	for _, tc := range tests {
		if MUX(tc.A, tc.B, tc.Sel) != tc.Out {
			t.Fatalf("expected MUX(%v, %v, %v) -> %v", tc.A, tc.B,
				tc.Sel, tc.Out)
		}
		if PureMUX(tc.A, tc.B, tc.Sel) != tc.Out {
			t.Fatalf("expected MUXP(%v, %v, %v) -> %v", tc.A, tc.B,
				tc.Sel, tc.Out)
		}
	}
}
