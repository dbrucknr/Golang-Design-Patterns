package pets

import "errors"

type PetInterface interface {
	SetSpecies(string) *Pet
	SetBreed(string) *Pet
	SetMinWeight(int) *Pet
	SetMaxWeight(int) *Pet
	SetWeight(int) *Pet
	SetDescription(string) *Pet
	SetLifespan(int) *Pet
	SetGeographicOrigin(string) *Pet
	SetColor(string) *Pet
	SetAge(int) *Pet
	SetAgeEstimated(bool) *Pet
	Build() (*Pet, error)
}

// Factory function to create a new PetBuilder instance
func NewPetBuilder() PetInterface {
	return &Pet{}
}

func (p *Pet) Build() (*Pet, error) {
	// You could add many more validation checks here
	if p.MinWeight > p.MaxWeight {
		return nil, errors.New("min weight cannot be greater than max weight")
	}
	// Computed Field
	p.AverageWeight = (p.MinWeight + p.MaxWeight) / 2
	return p, nil
}

func (p *Pet) SetSpecies(species string) *Pet {
	p.Species = species
	return p
}

func (p *Pet) SetBreed(breed string) *Pet {
	p.Breed = breed
	return p
}

func (p *Pet) SetMinWeight(minWeight int) *Pet {
	p.MinWeight = minWeight
	return p
}

func (p *Pet) SetMaxWeight(maxWeight int) *Pet {
	p.MaxWeight = maxWeight
	return p
}

func (p *Pet) SetWeight(weight int) *Pet {
	p.Weight = weight
	return p
}

func (p *Pet) SetDescription(description string) *Pet {
	p.Description = description
	return p
}

func (p *Pet) SetLifespan(lifespan int) *Pet {
	p.Lifespan = lifespan
	return p
}

func (p *Pet) SetGeographicOrigin(geographicOrigin string) *Pet {
	p.GeographicOrigin = geographicOrigin
	return p
}

func (p *Pet) SetColor(color string) *Pet {
	p.Color = color
	return p
}

func (p *Pet) SetAge(age int) *Pet {
	p.Age = age
	return p
}

func (p *Pet) SetAgeEstimated(ageEstimated bool) *Pet {
	p.AgeEstimated = ageEstimated
	return p
}
