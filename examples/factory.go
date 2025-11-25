package examples

import (
	"fmt"
	"time"
)

// Factory Pattern: Create an instance of an object with sensible default values.
// The presence of a New function can be a signal to the developer that creating the object is a complex process.
// There may be default values that need to be set, validation logic for required fields, or dependencies that need to be injected.
// Example usage:
func SampleFactoryUsage() {
	factory := Product{}
	product := factory.New()

	fmt.Println(product)
}

type Product struct {
	ProductName string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// We could add validation logic in this method to ensure that a product name is provided.
func (p *Product) New() *Product {
	product := Product{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return &product
}
