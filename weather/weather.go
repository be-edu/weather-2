package weather

import (
	"errors"
	"fmt"
)

// Prediction describes a weather prediction.
type Prediction uint8

// The supported weather prediction types (enum type).
const (
	Sunny Prediction = iota
	Rain
	Overcast
	Snow
	Unknown
)

// ToString returns the corresponding weather as a string
func (p Prediction) ToString() string {
	return [...]string{"Sunny", "Rain", "Overcast", "Snow", "Unknown"}[p]
}

// Locator is implemented by objects that can represent a location as a
// pair of GPS coordinates.
type Locator interface {
	Coords() (float64, float64, error)
}

// Predict the weather at the specified location.
func Predict(loc Locator) (Prediction, error) {
	coord1, coord2, err := loc.Coords()

	if err != nil {
		fmt.Println(coord1)
		fmt.Println(coord2)
		return Unknown, err
	}

	if coord1 < -90 || coord1 > 90 {
		return Unknown, errors.New("invalid latitude")
	} else if coord2 < -180 || coord2 > 180 {
		return Unknown, errors.New("invalid longitude")
	} else if coord1 >= -90 && coord1 < -60 {
		return Snow, nil
	} else if coord1 >= -60 && coord1 < -20 {
		return Sunny, nil
	} else if coord1 >= -20 && coord1 < 20 {
		return Overcast, nil
	} else if coord1 >= 20 && coord1 < 60 {
		return Rain, nil
	} else if coord1 >= 60 && coord1 <= 90 {
		return Snow, nil
	}

	return Unknown, nil
}

// GetVersion returns the package version
func GetVersion() string {
	return "v2.0.0"
}
