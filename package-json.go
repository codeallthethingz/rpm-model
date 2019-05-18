package model

// PackageJSON defines the structure of the package.json file.
type PackageJSON struct {
	Name         string            `json:"name,omitempty"`
	Version      string            `json:"version,omitempty"`
	Description  string            `json:"description"`
	License      string            `json:"license,omitempty"`
	Units        string            `json:"units,omitempty"`
	Metedata     []string          `json:"metadata,omitempty"`
	Bounds       []Bound           `json:"bounds,omitempty"`
	Joins        []string          `json:"joins,omitempty"`
	Dependencies map[string]string `json:"dependencies,omitempty"`
	Manifest     map[string][]CRS  `json:"manifest,omitempty"`
}

// CRS Coordinates Rotation Scale
type CRS struct {
	Coordinates Coordinates `json:"coordinates,omitempty"`
	Rotation    Coordinates `json:"rotation,omitempty"`
	Scale       Coordinates `json:"scale,omitempty"`
}

// Bound defines an area of a 3D object
type Bound struct {
	Name         string       `json:"name,omitempty"`
	BoundingType BoundingType `json:"boundingType,omitempty"`
	Coordinates  Coordinates  `json:"coordinates,omitempty"`
}

// Coordinates where is something in 3 dimensions
type Coordinates struct {
	X string `json:"x,omitempty"`
	Y string `json:"y,omitempty"`
	Z string `json:"z,omitempty"`
}

// BoundingType holds the info about the maximum dimensions of the object
type BoundingType struct {
	Name         string            `json:"name,omitempty"`
	Measurements map[string]string `json:"measurements,omitempty"`
	ScalesX      bool              `json:"scalesX,omitempty"`
	ScalesY      bool              `json:"scalesY,omitempty"`
	ScalesZ      bool              `json:"scalesZ,omitempty"`
}
