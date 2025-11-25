package models

import "time"

// This will be stored in a database
type DogBreed struct {
	Id               int    `json:"id"`
	Breed            string `json:"breed"`
	WeightLowLbs     int    `json:"weight_low_lbs"`
	WeightHighLbs    int    `json:"weight_high_lbs"`
	AverageWeightLbs int    `json:"average_weight_lbs"`
	Lifespan         int    `json:"lifespan"`
	Details          string `json:"details"`
	AlternateNames   string `json:"alternate_names"`
	GeographicOrigin string `json:"geographic_origin"`
}

// This will be a remote datasource (which is why its duplicated, even though it is identical to DogBreed)
type CatBreed struct {
	Id               int    `json:"id"`
	Breed            string `json:"breed"`
	WeightLowLbs     int    `json:"weight_low_lbs"`
	WeightHighLbs    int    `json:"weight_high_lbs"`
	AverageWeightLbs int    `json:"average_weight_lbs"`
	Lifespan         int    `json:"lifespan"`
	Details          string `json:"details"`
	AlternateNames   string `json:"alternate_names"`
	GeographicOrigin string `json:"geographic_origin"`
}

type Dog struct {
	Id               int       `json:"id"`
	DogName          string    `json:"dog_name"`
	BreedId          int       `json:"breed_id"`
	BreederId        int       `json:"breeder_id"`
	Color            string    `json:"color"`
	DateOfBirth      time.Time `json:"date_of_birth"`
	SpayedOrNeutered int       `json:"spayed_or_neutered"`
	Description      string    `json:"description"`
	WeightLbs        int       `json:"weight_lbs"`
	Breed            DogBreed  `json:"breed"`
	Breeder          Breeder   `json:"breeder"`
}
type Cat struct {
	Id               int       `json:"id"`
	DogName          string    `json:"dog_name"`
	BreedId          int       `json:"breed_id"`
	BreederId        int       `json:"breeder_id"`
	Color            string    `json:"color"`
	DateOfBirth      time.Time `json:"date_of_birth"`
	SpayedOrNeutered int       `json:"spayed_or_neutered"`
	Description      string    `json:"description"`
	WeightLbs        int       `json:"weight_lbs"`
	Breed            CatBreed  `json:"breed"`
	Breeder          Breeder   `json:"breeder"`
}

type Breeder struct {
	Id        int        `json:"id"`
	Name      string     `json:"name"`
	Address   string     `json:"address"`
	City      string     `json:"city"`
	ProvState string     `json:"prov_state"`
	ZipCode   string     `json:"zip_code"`
	Phone     string     `json:"phone"`
	Email     string     `json:"email"`
	Active    int        `json:"active"`
	DogBreeds []DogBreed `json:"dog_breeds"`
	CatBreeds []CatBreed `json:"cat_breeds"`
}

type Pet struct {
	Species     string `json:"species"`
	Breed       string `json:"breed"`
	MinWeight   int    `json:"min_weight"`
	MaxWeight   int    `json:"max_weight"`
	Description string `json:"description"`
	Lifespan    int    `json:"lifespan"`
}
