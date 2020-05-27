package shazam

// Cities allows named access to cities known to Shazam.
func Cities() City {
	return City{}
}

// City is a city known to Shazam.
type City struct{}
