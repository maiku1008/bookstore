// Package creditcard provides facilities for working with credit cards.
package creditcard

import "fmt"

// card represents a credit card
type card struct {
	number string
}

// New takes a credit card number and returns a 'card' value
// representing that card, or an error if the number is invalid.
func New(number string) (*card, error) {
	c := &card{}
	if err := c.SetNumber(number); err != nil {
		return nil, err
	}
	return c, nil
}

// Number returns the credit card's number
func (c *card) Number() string {
	return c.number
}

// SetNumber sets a digit string as the credit card number
// returns an error if the string is empty
func (c *card) SetNumber(number string) error {
	if number == "" {
		return fmt.Errorf("bad number %q", number)
	}
	c.number = number
	return nil
}
