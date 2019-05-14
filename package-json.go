package model

// PackageJSON defines the structure of the package.json file.
type PackageJSON struct {
	Name        string `json:"name,omitempty"`
	Version     string `json:"version,omitempty"`
	Description string `json:"description"`
	License     string `json:"license,omitempty"`
}
