package pets

import (
	"errors"
	"fmt"

	"github.com/dbrucknr/go-design-patterns/models"
)

// General interface - enforces the constraint
type AnimalInterface interface {
	Show() string
}

// This will be the concrete dog implementation that satisfies the AnimalInterface
type DogFromFactory struct {
	Pet *models.Dog
}

// Satisfies the AnimalInterface
func (dff *DogFromFactory) Show() string {
	return fmt.Sprintf("This animal is a %s", dff.Pet.Breed.Breed)
}

// This will be the concrete cat implementation that satisfies the AnimalInterface
type CatFromFactory struct {
	Pet *models.Cat
}

// Satisfies the AnimalInterface
func (cff *CatFromFactory) Show() string {
	return fmt.Sprintf("This animal is a %s", cff.Pet.Breed.Breed)
}

// General interface to generate either a Pet, which satisfies the AnimalInterface constraints
type PetFactoryInterface interface {
	newPet() AnimalInterface // Non-Exported (not capitalized)
}

type DogAbstractFactory struct{}

func (df *DogAbstractFactory) newPet() AnimalInterface {
	return &DogFromFactory{
		Pet: &models.Dog{},
	}
}

type CatAbstractFactory struct{}

func (df *CatAbstractFactory) newPet() AnimalInterface {
	return &CatFromFactory{
		Pet: &models.Cat{},
	}
}

func NewPetFromAbstractFactory(species string) (AnimalInterface, error) {
	switch species {
	case "dog":
		var dogFactory DogAbstractFactory
		dog := dogFactory.newPet()
		return dog, nil
	case "cat":
		var catFactory CatAbstractFactory
		cat := catFactory.newPet()
		return cat, nil
	default:
		return nil, errors.New("Invalid species supplied.")
	}
}
