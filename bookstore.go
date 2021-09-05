package bookstore

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const (
	CategoryAutobiography = iota
	CategoryLargePrintRomance
	CategoryParticlePhysics
)

var categories = map[int]bool{
	CategoryAutobiography:     true,
	CategoryLargePrintRomance: true,
	CategoryParticlePhysics:   true,
}

// Book represents a book in our bookstore
type Book struct {
	ID              string
	Title           string
	category        int
	Author          string
	Description     string
	discountPercent int
	PriceCents      int
	Quantity        int
}

func (b Book) SalePriceCents() int {
	saving := b.PriceCents * b.discountPercent / 100
	return b.PriceCents - saving
}

func (b *Book) SetPriceCents(price int) {
	b.PriceCents = price
}

// SetDiscountPercent sets the hidden field discountPercent
func (b *Book) SetDiscountPercent(percent int) error {
	if !(percent >= 0 && percent <= 100) {
		return fmt.Errorf("invalid percent: %d", percent)
	}
	b.discountPercent = percent
	return nil
}

// GetDiscountPercent returns the hidden field discountPercent
func (b *Book) GetDiscountPercent() int {
	return b.discountPercent
}

// SetCategory sets the unexported field for Book type
// returns an error if the category is wrong
func (b *Book) SetCategory(category int) error {
	if !validCategory(category) {
		return fmt.Errorf("unacceptable category: %d", category)
	}
	b.category = category
	return nil
}

// validCategory takes a string representing a category, and returns true if the
// category is valid.
func validCategory(category int) bool {
	return categories[category]
}

// GetCategory returns the category of the book
func (b *Book) GetCategory() int {
	return b.category
}

// Library is a type that represents the storage of the available books
type Library struct {
	Books map[string]Book
}

// NewID generates a new book ID that has two capital letters and two digits
func NewID() string {
	r1 := rand.Intn(26) + 'A'
	r2 := rand.Intn(26) + 'A'
	num := rand.Intn(100)
	return fmt.Sprintf("%c%c%02d", r1, r2, num)
}

// Customer represents a paying customer
type Customer struct {
	Title   string
	Name    string
	Address string
}

// PrintMailingLabel prints a string with the customer's credentials
func (c *Customer) PrintMailingLabel() string {
	return fmt.Sprintf("%s %s, %s", c.Title, c.Name, c.Address)
}

// Catalog represents a catalog of books.
type Catalog []Book

func (c *Catalog) AddBook(b Book) {
	*c = append(*c, b)
}

func (c Catalog) GetAllBooks() []Book {
	return []Book(c)
}

func (c *Catalog) GetLen() int {
	return len(*c)
}

func (c *Catalog) GetAllTitles() []string {
	titles := []string{}
	for _, b := range *c {
		titles = append(titles, b.Title)
	}
	return titles
}

func (c *Catalog) GetUniqueAuthors() []string {
	authors := []string{}
	seen := map[string]bool{}
	for _, b := range *c {
		if !seen[b.Author] {
			authors = append(authors, b.Author)
			seen[b.Author] = true
		}
	}
	return authors
}
