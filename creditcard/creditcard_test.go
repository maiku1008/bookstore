package creditcard_test

import (
	"bookstore/creditcard"
	"testing"
)

func TestNew(t *testing.T) {
	t.Parallel()
	var testCases = []struct {
		name        string
		number      string
		errExpected bool
	}{
		{name: "A valid number", number: "1234567890", errExpected: false},
		{name: "A valid number with letters", number: "1234CCC567890", errExpected: false},
		{name: "An empty string as the number", number: "", errExpected: true},
	}
	for _, tc := range testCases {
		_, err := creditcard.New(tc.number)
		errReceived := err != nil
		if errReceived != tc.errExpected {
			t.Fatalf("%s: Received unexpected error status: %s", tc.name, err.Error())
		}
	}
}

func TestNumber(t *testing.T) {
	t.Parallel()
	cc, _ := creditcard.New("1234567890")
	want := "1234567890"
	got := cc.Number()
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestSetNumber(t *testing.T) {
	t.Parallel()
	var testCases = []struct {
		name        string
		number      string
		errExpected bool
	}{
		{name: "", number: "5222", errExpected: false},
		{name: "", number: "", errExpected: true},
	}
	for _, tc := range testCases {

		cc, _ := creditcard.New("1")
		err := cc.SetNumber(tc.number)
		want := tc.number
		got := cc.Number()
		errReceived := err != nil
		if errReceived != tc.errExpected {
			t.Fatalf("%s: received unexpected error status: %s", tc.name, err.Error())
		}
		if !tc.errExpected && (want != got) {
			t.Errorf("%s: want %s, got %s", tc.name, want, got)
		}
	}
}
