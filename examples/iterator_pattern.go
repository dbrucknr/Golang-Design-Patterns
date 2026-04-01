package examples

import "fmt"

func iteratorSample() {
	// Step 1: Build Objects
	obj1 := ComplexObject{
		key:   "Some-Key-1",
		value: "Some-Value-1",
	}

	obj2 := ComplexObject{
		key:   "Some-Key-2",
		value: "Some-Value-2",
	}

	// Step 2: Create Collection
	collection := ComplexObjectCollection{
		objects: []*ComplexObject{&obj1, &obj2},
	}
	// Step 3: Generate Iterator
	iterator := collection.CreateIterator()

	// Step 4: Perform Iteration + Do Something
	for iterator.HasNext() {
		obj := iterator.GetNext()
		fmt.Printf("Object is %+v\n", obj)
	}
}

// --- Interfaces
type Collection interface {
	CreateIterator() Iterator
}

// NOTE - I wonder if a generic type T would be appropriate?
type Iterator interface {
	HasNext() bool
	GetNext() *ComplexObject
}

// --- Concrete Iterator Concept
type ComplexObjectIterator struct {
	Index   int
	Objects []*ComplexObject
}

func (i *ComplexObjectIterator) HasNext() bool {
	if i.Index < len(i.Objects) {
		return true
	}
	return false
}

func (i *ComplexObjectIterator) GetNext() *ComplexObject {
	if i.HasNext() {
		next := i.Objects[i.Index]
		i.Index++
		return next
	}
	return nil
}

// --- Concrete Objects
type ComplexObject struct {
	key   string
	value string
}

type ComplexObjectCollection struct {
	objects []*ComplexObject
}

func (c *ComplexObjectCollection) CreateIterator() Iterator {
	return &ComplexObjectIterator{
		Objects: c.objects,
	}
}
