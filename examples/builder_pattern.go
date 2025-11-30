package examples

import "fmt"

func SampleBuilderPatternUsage() {
	address := CreateAddress().
		WithStreet("123 Main St").
		WithCity("Anytown").
		WithState("CA").
		WithPostalCode("12345")

	fmt.Println(address)
}

type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}

// Fluent Interface / Builder Pattern
func CreateAddress() *Address {
	return &Address{}
}

func (a *Address) WithStreet(street string) *Address {
	a.Street = street
	return a
}

func (a *Address) WithCity(city string) *Address {
	a.City = city
	return a
}

func (a *Address) WithState(state string) *Address {
	a.State = state
	return a
}

func (a *Address) WithPostalCode(postalCode string) *Address {
	a.PostalCode = postalCode
	return a
}
