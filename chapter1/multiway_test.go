package chapter1

import "testing"

func TestOR8(t *testing.T) {
	var in [8]bool

	if OR8(in) {
		t.Fatalf("8 false bits should output false")
	} else if PureOR8(in) {
		t.Fatalf("8 false bits should output false")
	}

	for i := 1; i < 8; i++ {
		in[i] = true
		if !OR8(in) {
			t.Fatalf("8 with any true bits should output true")
		} else if !PureOR8(in) {
			t.Fatalf("8 with any true bits should output true")
		}
	}
}

func cmp8(a, b [8]bool) bool {
	for i := 0; i < 8; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestMUX4W8(t *testing.T) {
	a := [8]bool{true, false, false, false, true, false, false, false}
	b := [8]bool{false, true, false, false, false, true, false, false}
	c := [8]bool{false, false, true, false, false, false, true, false}
	d := [8]bool{false, false, false, true, false, false, false, true}

	out := MUX4W8(a, b, c, d, [2]bool{false, false})
	if !cmp8(a, out) {
		t.Fatalf("sel{0,0} should output bus A (%v)", out)
	}
	out = MUX4W8(a, b, c, d, [2]bool{false, true})
	if !cmp8(b, out) {
		t.Fatalf("sel{0,1} should output bus B (%v)", out)
	}
	out = MUX4W8(a, b, c, d, [2]bool{true, false})
	if !cmp8(c, out) {
		t.Fatalf("sel{1,0} should output bus C (%v)", out)
	}
	out = MUX4W8(a, b, c, d, [2]bool{true, true})
	if !cmp8(d, out) {
		t.Fatalf("sel{1,1} should output bus D (%v)", out)
	}

	out = PureMUX4W8(a, b, c, d, [2]bool{false, false})
	if !cmp8(a, out) {
		t.Fatalf("sel{0,0} should output bus A (%v)", out)
	}
	out = PureMUX4W8(a, b, c, d, [2]bool{false, true})
	if !cmp8(b, out) {
		t.Fatalf("sel{0,1} should output bus B (%v)", out)
	}
	out = PureMUX4W8(a, b, c, d, [2]bool{true, false})
	if !cmp8(c, out) {
		t.Fatalf("sel{1,0} should output bus C (%v)", out)
	}
	out = PureMUX4W8(a, b, c, d, [2]bool{true, true})
	if !cmp8(d, out) {
		t.Fatalf("sel{1,1} should output bus D (%v)", out)
	}
}
