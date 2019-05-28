package model

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	contents, err := ioutil.ReadFile("examples/bolt-round-20mm/package.json")
	require.Nil(t, err)

	var p PackageJSON
	require.Nil(t, json.Unmarshal(contents, &p))

	data, err := json.MarshalIndent(&p, "", "  ")
	require.Nil(t, err)

	require.Equal(t, "bolt-round-20mm", p.Name)
	require.Equal(t, "0.0.1", p.Version)
	require.Equal(t, "simple round-bolt 1.5mm pitch", p.Description)
	require.Equal(t, "MIT", p.License)
	require.Equal(t, "mm", p.Units)
	require.Equal(t, []string{"flat-head"}, p.Metedata)

	require.Equal(t, 3, len(p.Bounds))
	bounds0 := p.Bounds[0]
	require.Equal(t, "total-area", bounds0.Name)
	require.Equal(t, "total-area", bounds0.Name)
	require.Equal(t, "round-cylinder", bounds0.BoundingType.Name)
	require.Equal(t, "0", bounds0.Coordinates.X)
	require.Equal(t, "0", bounds0.Coordinates.Y)
	require.Equal(t, "0", bounds0.Coordinates.Z)
	require.Equal(t, false, bounds0.BoundingType.ScalesX)
	require.Equal(t, false, bounds0.BoundingType.ScalesY)
	require.Equal(t, false, bounds0.BoundingType.ScalesZ)
	require.Equal(t, "7.5", bounds0.BoundingType.Measurements["radius"])
	require.Equal(t, "20", bounds0.BoundingType.Measurements["height"])

	require.Equal(t, "shaft", p.Joins[0])
	require.Equal(t, "0.0.3", p.Dependencies["flat-head-screwdriver-hole"])
	screwdriverHole := p.Manifest["flat-head-screwdriver-hole"]
	require.Equal(t, "0", screwdriverHole[0].Coordinates.X)
	require.Equal(t, "0", screwdriverHole[0].Coordinates.Y)
	require.Equal(t, "18", screwdriverHole[0].Coordinates.Z)
	require.Equal(t, "0", screwdriverHole[0].Rotation.X)
	require.Equal(t, "0", screwdriverHole[0].Rotation.Y)
	require.Equal(t, "0", screwdriverHole[0].Rotation.Z)
	require.Equal(t, "1", screwdriverHole[0].Scale.X)
	require.Equal(t, "1", screwdriverHole[0].Scale.Y)
	require.Equal(t, "1", screwdriverHole[0].Scale.Z)

	require.Equal(t, string(data), string(contents))
}
