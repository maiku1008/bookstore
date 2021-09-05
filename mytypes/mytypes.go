package mytypes

import (
	"fmt"
	"strconv"
	"strings"
)

// MyInt is a custom int type
type MyInt int

// Twice returns a value of the myInt multiplied by 2
func (i MyInt) Twice() MyInt {
	return i * 2
}

// Double doubles the value pointed at by i, returns nothing
func (i *MyInt) Double() {
	*i *= 2
}

// MultiplyBy multiplies the value pointed at by i, returns nothing
func (i *MyInt) MultiplyBy(operand MyInt) {
	*i *= operand
}

// MyString is a custom string
type MyString string

// Len returns a value with the length of the string s
func (s MyString) Len() int {
	return len(s)
}

// MultiString is a custom type []string
type MultiString []string

// Join returns a value string with all elements of MultiString, separated by separator
func (m MultiString) Join(separator string) string {
	return strings.Join(m, separator)
}

// Add appends a string c to the value pointed at by m
func (m *MultiString) Add(c string) {
	*m = append(*m, c)
}

// MyBuilder is a custom version of strings.Builder
type MyBuilder struct {
	strings.Builder
}

// Hello returns the string
func (mb MyBuilder) Hello() string {
	return "Hello, Gophers!"
}

// StringUpperCaser is a custom struct that embeds a strings.Builder
type StringUpperCaser struct {
	strings.Builder
}

// ToUpper returns an uppercase version of the string
func (suc *StringUpperCaser) ToUpper() string {
	return strings.ToUpper(suc.String())
}

// SetString adds an input string to the embedded strings.Builder string
func (suc *StringUpperCaser) SetString(input string) {
	suc.WriteString(input)
}

// GetString returns the strings belonging to the embedded strings.Builder string
func (suc *StringUpperCaser) GetString() string {
	return suc.String()
}

// Square squares the int value pointed at by x, and returns nothing
func Square(x *int) {
	*x *= *x
}

// SwapInts points the value pointed at by x to y and viceversa, and returns nothing
func SwapInts(x, y *int) {
	*x, *y = *y, *x
}

// MyMap is a custom map type
type MyMap map[string]string

// Add adds a key and a value to m
func (m *MyMap) Add(k, v string) {
	(*m)[k] = v
}

// MySlice is a custom []string type
type MySlice []string

// Add appends a string v to the slice pointed at by m
func (m *MySlice) Add(v string) {
	*m = append(*m, v)
}

// DigitString is a type that contains only strings that have valid numbers
type DigitString struct {
	digitString string
}

func NewDigitString() *DigitString {
	var x DigitString
	return &x
}

// GetDigitString returns the value of the protected field digitString
func (ds *DigitString) GetDigitString() string {
	return ds.digitString
}

// SetDigitString sets the value of the protected field digitString
// returns an error if the input doesn't contain valid digits
func (ds *DigitString) SetDigitString(number string) error {
	if _, err := strconv.Atoi(number); err != nil {
		return fmt.Errorf("invalid input is not a number: %s", number)
	}
	ds.digitString = number
	return nil
}
