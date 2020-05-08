package logins

// Location represents a geographical location
type Location struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
	Radius    int     `json:"radius"`
}
