package examples

import (
	"fmt"
)

// Abstract Factory: Provides an interface for creating families of related or dependent objects without specifying their concrete classes.
// Decouple client code from specific classes of products.
// Example usage:
func SampleAbstractFactoryUsage() {
	DogFactory := &DogFactory{}
	CatFactory := &CatFactory{}

	dog := DogFactory.New()
	cat := CatFactory.New()

	fmt.Println(dog.Speak(), dog.LikesWater())
	fmt.Println(cat.Speak(), cat.LikesWater())
}

// Type for AbstractFactory
// This enforces a constraint on the implementation of the Animal interface.
// Client code doesn't care how the information is put together.
// This makes it easy to switch where an information source comes from (Database vs. API vs. ...)
type Animal interface {
	// Methods
	Speak() string
	LikesWater() bool
}

// Create a ConcreteFactory for Dogs
type Dog struct{}

// Implement the AbstractFactory interface for Dogs
func (d *Dog) Speak() string {
	return "Woof!"
}

func (d *Dog) LikesWater() bool {
	return true
}

// Create a ConcreteFactory for Cats
type Cat struct{}

// Implement the AbstractFactory interface for Cats
func (c *Cat) Speak() string {
	return "Meow!"
}

func (c *Cat) LikesWater() bool {
	return false
}

type AnimalFactory interface {
	New() Animal
}

type DogFactory struct{}

func (df *DogFactory) New() Animal {
	return &Dog{}
}

type CatFactory struct{}

func (cf *CatFactory) New() Animal {
	return &Cat{}
}
