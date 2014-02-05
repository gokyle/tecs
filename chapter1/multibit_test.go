package chapter1

import "testing"

func TestNOT16(t *testing.T) {
	var in [16]bool
	out := NOT16(in)
	for i, v := range out {
		if !v {
			t.Fatalf("out[%d]=%v should be true", i, in[i])
		}
	}

	in[7] = true
	out = NOT16(in)
	for i, v := range out {
		if !v && i != 7 {
			t.Fatalf("a[%d]=%v should be true", i, in[i])
		} else if v && i == 7 {
			t.Fatalf("a[%d]=%v should be false", i, in[i])
		}
	}
}

func TestAND16(t *testing.T) {
	var a, b, out [16]bool

	out = AND16(a, b)
	for i, v := range out {
		if v {
			t.Fatalf("AND(a[%d]=%v, b[%d]=%v) should be false",
				i, a[i], i, b[i])
		}
	}

	out = PureAND16(a, b)
	for i, v := range out {
		if v {
			t.Fatalf("AND(a[%d]=%v, b[%d]=%v) should be false",
				i, a[i], i, b[i])
		}
	}

	a[0] = true
	b[0] = true
	out = AND16(a, b)
	for i, v := range out {
		if v && i != 0 {
			t.Fatalf("AND(a[%d]=%v, b[%d]=%v) should be false",
				i, a[i], i, b[i])
		} else if !v && i == 0 {
			t.Fatalf("AND(a[%d]=%v, b[%d]=%v) should be false",
				i, a[i], i, b[i])
		}
	}

	out = PureAND16(a, b)
	for i, v := range out {
		if v && i != 0 {
			t.Fatalf("AND(a[%d]=%v, b[%d]=%v) should be false",
				i, a[i], i, b[i])
		} else if !v && i == 0 {
			t.Fatalf("AND(a[%d]=%v, b[%d]=%v) should be false",
				i, a[i], i, b[i])
		}
	}
}

func TestOR16(t *testing.T) {
	var a, b, out [16]bool

	out = OR16(a, b)
	for i, v := range out {
		if v {
			t.Fatalf("AND(a[%d]=%v, b[%d]=%v) should be false",
				i, a[i], i, b[i])
		}
	}

	out = PureOR16(a, b)
	for i, v := range out {
		if v {
			t.Fatalf("AND(a[%d]=%v, b[%d]=%v) should be false",
				i, a[i], i, b[i])
		}
	}

	a[0] = true
	out = OR16(a, b)
	for i, v := range out {
		if v && i != 0 {
			t.Fatalf("AND(a[%d]=%v, b[%d]=%v) should be false",
				i, a[i], i, b[i])
		} else if !v && i == 0 {
			t.Fatalf("AND(a[%d]=%v, b[%d]=%v) should be false",
				i, a[i], i, b[i])
		}
	}

	out = PureOR16(a, b)
	for i, v := range out {
		if v && i != 0 {
			t.Fatalf("AND(a[%d]=%v, b[%d]=%v) should be false",
				i, a[i], i, b[i])
		} else if !v && i == 0 {
			t.Fatalf("AND(a[%d]=%v, b[%d]=%v) should be false",
				i, a[i], i, b[i])
		}
	}

	a[0] = false
	b[0] = true
	out = OR16(a, b)
	for i, v := range out {
		if v && i != 0 {
			t.Fatalf("AND(a[%d]=%v, b[%d]=%v) should be false",
				i, a[i], i, b[i])
		} else if !v && i == 0 {
			t.Fatalf("AND(a[%d]=%v, b[%d]=%v) should be false",
				i, a[i], i, b[i])
		}
	}

	out = PureOR16(a, b)
	for i, v := range out {
		if v && i != 0 {
			t.Fatalf("AND(a[%d]=%v, b[%d]=%v) should be false",
				i, a[i], i, b[i])
		} else if !v && i == 0 {
			t.Fatalf("AND(a[%d]=%v, b[%d]=%v) should be false",
				i, a[i], i, b[i])
		}
	}
}

func TestMUX16(t *testing.T) {
	var a, b [16]bool
	var out [16]bool

	for i := 0; i < 16; i++ {
		a[i] = true
		i++
		b[i] = true
	}

	out = MUX16(a, b, false)
	for i, v := range out {
		if (i%2 == 0) && !v {
			t.Fatalf("bit %d should be true", i)
		} else if (i%2 != 0) && v {
			t.Fatalf("bit %d should be false", i)
		}
	}

	out = PureMUX16(a, b, false)
	for i, v := range out {
		if (i%2 == 0) && !v {
			t.Fatalf("bit %d should be true", i)
		} else if (i%2 != 0) && v {
			t.Fatalf("bit %d should be false", i)
		}
	}

	out = MUX16(a, b, true)
	for i, v := range out {
		if (i%2 != 0) && !v {
			t.Fatalf("bit %d should be true", i)
		} else if (i%2 == 0) && v {
			t.Fatalf("bit %d should be false", i)
		}
	}

	out = PureMUX16(a, b, true)
	for i, v := range out {
		if (i%2 != 0) && !v {
			t.Fatalf("bit %d should be true", i)
		} else if (i%2 == 0) && v {
			t.Fatalf("bit %d should be false", i)
		}
	}
}
