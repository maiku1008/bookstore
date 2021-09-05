package bookstore_test

import (
	"bookstore"
	"testing"

	"github.com/google/go-cmp/cmp/cmpopts"

	"github.com/google/go-cmp/cmp"
)

func TestBook(t *testing.T) {
	t.Parallel()
	_ = bookstore.Book{
		ID:          "87AS",
		Title:       "Spark Joy",
		Author:      "Marie Kondo",
		Description: "A tiny, cheerful Japanese woman explains tidying.",
		PriceCents:  1199,
		Quantity:    1,
	}
}

func TestNewID(t *testing.T) {
	t.Parallel()
	unique := make(map[string]bool)
	for i := 0; i < 100; i++ {
		unique[bookstore.NewID()] = true
	}

	want := 100
	got := len(unique)
	if want != got {
		t.Errorf("want %d unique IDs, got %d", want, got)
	}
}

func TestSalePriceCents(t *testing.T) {
	t.Parallel()
	// TODO: fix this up
	b := bookstore.Book{
		Title:      "Ponyo on the cliff by the sea",
		PriceCents: 500,
	}
	b.SetDiscountPercent(50)
	want := 250
	got := b.SalePriceCents()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestSetDiscountPercent(t *testing.T) {
	t.Parallel()
	var testCases = []struct {
		name        string
		percent     int
		want        int
		errExpected bool
	}{
		{name: "zero percent", percent: 0, want: 0, errExpected: false},
		{name: "50 percent", percent: 50, want: 50, errExpected: false},
		{name: "zero percent", percent: 100, want: 100, errExpected: false},
		{name: "invalid negative percent", percent: -100, want: 0, errExpected: true},
		{name: "invalid positive percent", percent: 101, want: 100, errExpected: true},
	}
	for _, tc := range testCases {
		b := bookstore.Book{}
		err := b.SetDiscountPercent(tc.percent)
		got := b.GetDiscountPercent()
		errReceived := err != nil
		if errReceived != tc.errExpected {
			t.Fatalf("%s: Received unexpected error status: %s", tc.name, err.Error())
		}
		if !tc.errExpected && (tc.want != got) {
			t.Errorf("%s: want %d, got %d", tc.name, tc.want, got)
		}
	}
}

func TestCustomer(t *testing.T) {
	t.Parallel()
	_ = bookstore.Customer{
		Title:   "Dr.",
		Name:    "Alan Beatty",
		Address: "Imaginary rd 108, Kansas City, Missouri 19874",
	}
}

func TestPrintMailingLabel(t *testing.T) {
	t.Parallel()
	c := bookstore.Customer{
		Title:   "Dr.",
		Name:    "Alan Beatty",
		Address: "Imaginary rd 108, Kansas City, Missouri 19874",
	}
	want := "Dr. Alan Beatty, Imaginary rd 108, Kansas City, Missouri 19874"
	got := c.PrintMailingLabel()
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

var books = []bookstore.Book{
	{Title: "Foundation", Author: "Isaac Asimov"},
	{Title: "Foundation and Empire", Author: "Isaac Asimov"},
	{Title: "Ponyo on the cliff by the sea", Author: "Hayao Miyazaki"},
}

func TestCatalogAddBook(t *testing.T) {
	t.Parallel()
	want := books[2]
	b := bookstore.Catalog{}
	b.AddBook(books[2])
	got := b[0]
	if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(bookstore.Book{})) {
		t.Errorf(cmp.Diff(want, got))
	}
}

func TestCatalogGetAllBooks(t *testing.T) {
	t.Parallel()
	want := books
	c := bookstore.Catalog(want)
	got := c.GetAllBooks()
	if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(bookstore.Book{})) {
		t.Errorf(cmp.Diff(want, got))
	}
}

func TestCatalogLen(t *testing.T) {
	t.Parallel()
	c := bookstore.Catalog(books)
	want := 3
	got := c.GetLen()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestCatalogGetAllTitles(t *testing.T) {
	t.Parallel()
	c := bookstore.Catalog(books)
	want := []string{"Foundation", "Foundation and Empire", "Ponyo on the cliff by the sea"}
	got := c.GetAllTitles()
	if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(bookstore.Book{})) {
		t.Errorf(cmp.Diff(want, got))
	}
}

func TestCatalogGetUniqueAuthors(t *testing.T) {
	t.Parallel()
	c := bookstore.Catalog(books)
	want := []string{"Isaac Asimov", "Hayao Miyazaki"}
	got := c.GetUniqueAuthors()
	if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(bookstore.Book{})) {
		t.Errorf(cmp.Diff(want, got))
	}
}

func TestSetPriceCents(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{PriceCents: 1000}
	want := 1200
	b.SetPriceCents(1200)
	got := b.PriceCents
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestSetCategory(t *testing.T) {
	t.Parallel()
	var testCases = []struct {
		name        string
		category    int
		errExpected bool
	}{
		{name: "An autobiography", category: bookstore.CategoryAutobiography, errExpected: false},
		{name: "An mystery novel", category: bookstore.CategoryLargePrintRomance, errExpected: false},
		{name: "An erotic novel", category: bookstore.CategoryParticlePhysics, errExpected: false},
		{name: "A fantasy novel", category: 1001, errExpected: true},
	}
	for _, tc := range testCases {
		b := bookstore.Book{}
		err := b.SetCategory(tc.category)
		want := tc.category
		got := b.GetCategory()
		errReceived := err != nil
		if errReceived != tc.errExpected {
			t.Fatalf("%s: Received unexpected error status: %s", tc.name, err.Error())
		}
		if !tc.errExpected && (want != got) {
			t.Errorf("%s: want %d, got %d", tc.name, want, got)
		}
	}
}
