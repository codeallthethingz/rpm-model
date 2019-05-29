package model

// PackageJSON defines the structure of the package.json file.
type PackageJSON struct {
	Name         string            `json:"name,omitempty"`
	Version      string            `json:"version,omitempty"`
	Description  string            `json:"description"`
	License      string            `json:"license,omitempty"`
	Units        Unit              `json:"units,omitempty"`
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

// Unit of measure
type Unit string

// PredefinedUnits that must be used in a package.json
var PredefinedUnits = []Unit{
	Millimeter, Centimeter, Meter, Kilometer,
	Inch, Foot, Yard, Mile,
}

// Units
const (
	Millimeter Unit = "mm"
	Centimeter Unit = "cm"
	Meter      Unit = "m"
	Kilometer  Unit = "km"
	Inch       Unit = "inch"
	Foot       Unit = "foot"
	Yard       Unit = "yard"
	Mile       Unit = "mile"
)

// PredefinedBoundingTypes that must be used for the mandatory total-area bounding type
var PredefinedBoundingTypes = []BoundingType{
	TriangularPrism, Cuboid, PentagonalPrism,
	HexagonalPrism, HeptagonalPrism, OctagonalPrism,
	NonagonPrism, DecagonalPrism, RoundCylinder,
	OvalCylinder, Sphere, Ellipsoid,
}

// Ellipsoid 2 radius ellipsoid
var Ellipsoid = BoundingType{
	Name: "ellipsoid",
	Measurements: map[string]string{
		"radius1": "",
		"radius2": "",
	},
}

// Sphere just is a sphere.  The simplest shape to describe, the hardest to create.
var Sphere = BoundingType{
	Name: "sphere",
	Measurements: map[string]string{
		"radius": "",
	},
}

// OvalCylinder oval prism
var OvalCylinder = BoundingType{
	Name: "oval-cylinder",
	Measurements: map[string]string{
		"radius1": "",
		"radius2": "",
		"height":  "",
	},
}

// RoundCylinder round prism
var RoundCylinder = BoundingType{
	Name: "round-cylinder",
	Measurements: map[string]string{
		"radius": "",
		"height": "",
	},
}

// DecagonalPrism 10 equal sided prism
var DecagonalPrism = BoundingType{
	Name: "pecagonal-prism",
	Measurements: map[string]string{
		"radius": "",
		"height": "",
	},
}

// NonagonPrism 9 equal sided prism
var NonagonPrism = BoundingType{
	Name: "nonagon-prism",
	Measurements: map[string]string{
		"radius": "",
		"height": "",
	},
}

// OctagonalPrism 8 equal sided prism
var OctagonalPrism = BoundingType{
	Name: "octagonal-prism",
	Measurements: map[string]string{
		"radius": "",
		"height": "",
	},
}

// HeptagonalPrism 7 equal sided prism
var HeptagonalPrism = BoundingType{
	Name: "heptagonal-prism",
	Measurements: map[string]string{
		"radius": "",
		"height": "",
	},
}

// HexagonalPrism 6 equal sided prism
var HexagonalPrism = BoundingType{
	Name: "hexagonal-prism",
	Measurements: map[string]string{
		"radius": "",
		"height": "",
	},
}

// PentagonalPrism 5 equal sided prism
var PentagonalPrism = BoundingType{
	Name: "pentagonal-prism",
	Measurements: map[string]string{
		"radius": "",
		"height": "",
	},
}

// Cuboid right-angle rectangular prism
var Cuboid = BoundingType{
	Name: "cuboid",
	Measurements: map[string]string{
		"width":  "",
		"length": "",
		"height": "",
	},
}

// TriangularPrism 3d Triangle.
var TriangularPrism = BoundingType{
	Name: "triangular-prism",
	Measurements: map[string]string{
		"radius": "",
		"height": "",
	},
}
