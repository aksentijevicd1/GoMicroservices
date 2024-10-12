package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "CoffeTea",
		Price: 45,
		SKU:   "abs-abs-abs",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
