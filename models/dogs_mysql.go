package models

import (
	"context"
	"time"
)

// Satisfies the Repository interface
func (m *MySQLRepository) AllDogBreeds() ([]*DogBreed, error) {
	// If the Db takes longer than 3 seconds, cancel the request
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// We use coalesce to handle possible null values. Go doesn't like those, so we set defaults
	query := `
		SELECT
			id,
			breed,
			weight_low_lbs,
			weight_high_lbs,
			cast(((weight_low_lbs + weight_high_lbs) /2) as unsigned) as avg_weight,
		 	lifespan,
			coalesce(details, ''),
			coalesce(alternate_names, ''),
			coalesce(geographic_origin, '')
		FROM
			dog_breeds
		ORDER BY
			breed
	`

	var breeds []*DogBreed
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	// When the function exits, close the rows
	defer rows.Close()

	for rows.Next() {
		var b DogBreed
		err := rows.Scan(
			&b.Id,
			&b.Breed,
			&b.WeightLowLbs,
			&b.WeightHighLbs,
			&b.AverageWeightLbs,
			&b.Lifespan,
			&b.Details,
			&b.AlternateNames,
			&b.GeographicOrigin,
		)
		if err != nil {
			return nil, err
		}
		breeds = append(breeds, &b)
	}

	return breeds, nil
}
