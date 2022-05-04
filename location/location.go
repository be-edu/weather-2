package location

// GPSCoords holds a lat/long coordinate pair.
type GPSCoords [2]float64

// PlusCode encodes a location using a plus code.
type PlusCode string

// Address encapsulates the components of an address.
type Address struct {
	Street   string
	City     string
	PostCode string
	Country  string
}

func (g GPSCoords) Coords() (float64, float64, error) {
	return g[0], g[1], nil
}
func (p PlusCode) Coords() (float64, float64, error) {
	// Decode plus code to gps coordinates...

	return 0, 0, nil
}
func (a Address) Coords() (float64, float64, error) {
	// Use an external geocoding service to convert the address into a set
	// of GPS coordinates...

	return 0, 0, nil
}
