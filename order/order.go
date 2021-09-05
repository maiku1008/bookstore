package order

import "fmt"

type order struct {
	id string
}

func New(id string) (*order, error) {
	if id == "" {
		return nil, fmt.Errorf("bad id: %q", id)
	}
	return &order{id: id}, nil
}
