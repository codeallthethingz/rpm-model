package model

// DependencyResponse the response from the dependency edpoint
type DependencyResponse struct {
	Dependencies []*Dependency `json:"dependencies"`
}

// Dependency A dependency from one object to another
type Dependency struct {
	Owner   string `json:"owner"`
	OwnerID string `json:"ownerId"`
	Name    string `json:"name"`
	Version string `json:"version"`
}
