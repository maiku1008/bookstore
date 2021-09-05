package order_test

import (
	"bookstore/order"
	"testing"
)

func TestNew(t *testing.T) {
	t.Parallel()
	var testCases = []struct {
		name        string
		id          string
		errExpected bool
	}{
		{name: "A valid id", id: "1234567890", errExpected: false},
		{name: "An empty string as the id", id: "", errExpected: true},
	}
	for _, tc := range testCases {
		o, err := order.New(tc.id)
		errReceived := err != nil
		if errReceived != tc.errExpected {
			t.Fatalf("%s: Received unexpected error status: %s", tc.name, err.Error())
		}
		if !tc.errExpected && (o == nil) {
			t.Errorf("%s: got unexpected nil order", tc.name)
		}
	}
}
