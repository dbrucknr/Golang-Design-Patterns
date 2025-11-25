package examples

import (
	"time"
)

// Factory Pattern: Create an instance of an object with sensible default values.
// factory := products.Product{}
// product := factory.New()
// The presence of a New function can be a signal to the developer that creating the object is a complex process.
// There may be default values that need to be set, or dependencies that need to be injected.

type Product struct {
	ProductName string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (p *Product) New() *Product {
	product := Product{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return &product
}
