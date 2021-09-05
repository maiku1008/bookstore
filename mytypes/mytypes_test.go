package mytypes_test

import (
	"bookstore/mytypes"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMyStringLen(t *testing.T) {
	s := mytypes.MyString("Hello, Gophers!")
	want := 15
	got := s.Len()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestMultiStringJoin(t *testing.T) {
	m := mytypes.MultiString{"a", "b", "c"}
	want := "a plus b plus c"
	got := m.Join(" plus ")
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestMultiStringAdd(t *testing.T) {
	m := mytypes.MultiString{}
	want := mytypes.MultiString{"a"}
	m.Add("a")
	got := m
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestStringUpperCaser(t *testing.T) {
	m := mytypes.StringUpperCaser{}
	want := "HELLO WORLD!"
	m.WriteString("hello world!")
	got := m.ToUpper()
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestSquare(t *testing.T) {
	x := 4
	want := 16
	mytypes.Square(&x)
	got := x
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestSwapInts(t *testing.T) {
	x, y := 4, 8
	wantX, wantY := 8, 4
	mytypes.SwapInts(&x, &y)
	gotX, gotY := x, y
	if wantX != gotX || wantY != gotY {
		t.Error("no")
	}
}

func TestDouble(t *testing.T) {
	var x mytypes.MyInt = 3
	var want mytypes.MyInt = 6
	x.Double()
	got := x
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestMultiplyBy(t *testing.T) {
	var x, y, want mytypes.MyInt = 12, 2, 24
	x.MultiplyBy(y)
	got := x
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestMyMap(t *testing.T) {
	m := mytypes.MyMap{}
	m.Add("name", "Michael")
	want := "Michael"
	got := m["name"]
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestMySlice(t *testing.T) {
	m := mytypes.MySlice{}
	m.Add("Michael")
	want := "Michael"
	got := m[0]
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestDigitString(t *testing.T) {
	t.Parallel()
	var testCases = []struct {
		name        string
		want        string
		errExpected bool
	}{
		{name: "valid input with 1 digit", want: "5", errExpected: false},
		{name: "valid input with 2 digits", want: "53", errExpected: false},
		{name: "valid input with 3 digits", want: "542", errExpected: false},
		{name: "valid input with 4 digits", want: "5222", errExpected: false},
		{name: "valid input with 5 digits", want: "543423", errExpected: false},
		{name: "invalid input with one letter", want: "a", errExpected: true},
		{name: "invalid input with one letter and one number", want: "a3", errExpected: true},
		{name: "invalid input with many numbers and one character", want: "3485623895623895s", errExpected: true},
		{name: "invalid input a word", want: "hello", errExpected: true},
		{name: "invalid input 3letter and a number", want: "aaa5", errExpected: true},
		{name: "a negative number", want: "-5", errExpected: false},
	}
	for _, tc := range testCases {
		d := mytypes.NewDigitString()
		err := d.SetDigitString(tc.want)
		got := d.GetDigitString()
		errReceived := err != nil
		if errReceived != tc.errExpected {
			t.Fatalf("%s: received unexpected error status: %s", tc.name, err.Error())
		}
		if !tc.errExpected && (tc.want != got) {
			t.Errorf("%s: want %s, got %s", tc.name, tc.want, got)
		}
	}
}
