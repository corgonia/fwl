package fwl_test

import (
	"github.com/corgonia/fwl"
	"testing"
)

type WorldTest struct {
	RawData []byte
	World   fwl.World
}

var testData = []WorldTest{
	{
		RawData: []byte{
			0x2A, 0x00, 0x00, 0x00, 0x1A, 0x00, 0x00, 0x00,
			0x0A, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x74,
			0x65, 0x73, 0x74, 0x0A, 0x74, 0x4B, 0x71, 0x47,
			0x42, 0x7A, 0x4E, 0x67, 0x66, 0x41, 0xD3, 0xC4,
			0x82, 0xBC, 0x56, 0xA1, 0xDB, 0x02, 0x00, 0x00,
			0x00, 0x00, 0x01, 0x00, 0x00, 0x00},
		World: fwl.World{
			Name:            "servertest",
			Seed:            "tKqGBzNgfA",
			SeedValue:       -1132280621,
			UID:             47948118,
			WorldVersion:    26,
			WorldGenVersion: 1,
		},
	},
	{
		RawData: []byte{
			0x26, 0x00, 0x00, 0x00, 0x1A, 0x00, 0x00, 0x00,
			0x06, 0x41, 0x62, 0x61, 0x72, 0x61, 0x74, 0x0A,
			0x76, 0x6E, 0x7A, 0x57, 0x43, 0x73, 0x6E, 0x49,
			0x6D, 0x71, 0x7C, 0x1B, 0x72, 0x31, 0x62, 0x81,
			0x91, 0xBB, 0xFF, 0xFF, 0xFF, 0xFF, 0x01, 0x00,
			0x00, 0x00},
		World: fwl.World{
			Name:            "Abarat",
			Seed:            "vnzWCsnImq",
			SeedValue:       829561724,
			UID:             -1148092062,
			WorldVersion:    26,
			WorldGenVersion: 1,
		},
	},
	{
		RawData: []byte{
			0x29, 0x00, 0x00, 0x00, 0x1A, 0x00, 0x00, 0x00,
			0x09, 0x44, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74,
			0x65, 0x64, 0x0A, 0x6D, 0x45, 0x43, 0x69, 0x73,
			0x76, 0x76, 0x62, 0x43, 0x71, 0xA9, 0x1B, 0x37,
			0xD4, 0x99, 0x81, 0x10, 0x45, 0x00, 0x00, 0x00,
			0x00, 0x01, 0x00, 0x00, 0x00},
		World: fwl.World{
			Name:            "Dedicated",
			Seed:            "mECisvvbCq",
			SeedValue:       -734585943,
			UID:             1158709657,
			WorldVersion:    26,
			WorldGenVersion: 1,
		},
	},
	{
		RawData: []byte{
			0x27, 0x00, 0x00, 0x00, 0x1A, 0x00, 0x00, 0x00,
			0x07, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6E, 0x67,
			0x0A, 0x42, 0x4B, 0x34, 0x72, 0x45, 0x56, 0x53,
			0x48, 0x6B, 0x64, 0x4C, 0x25, 0x02, 0x48, 0xC1,
			0xC7, 0xC4, 0xB5, 0xFF, 0xFF, 0xFF, 0xFF, 0x01,
			0x00, 0x00, 0x00},
		World: fwl.World{
			Name:            "testing",
			Seed:            "BK4rEVSHkd",
			SeedValue:       1208100172,
			UID:             -1245395007,
			WorldVersion:    26,
			WorldGenVersion: 1,
		},
	},
}

func TestGetStableHashCode(t *testing.T) {

	for _, w := range testData {

		world := fwl.World{}

		err := world.UnmarshalBinary(w.RawData)
		if err != nil {
			t.Error("UnmashalBinary failed:", err)
		}

		if fwl.GetStableHashCode(world.Seed) != w.World.SeedValue {
			t.Error("Seed value does not match seed hash", world.Seed, fwl.GetStableHashCode(world.Seed), w.World.SeedValue)
		}

	}

}
func TestWorld_UnmarshalBinary(t *testing.T) {
	for _, w := range testData {

		world := fwl.World{}

		err := world.UnmarshalBinary(w.RawData)
		if err != nil {
			t.Error("UnmashalBinary failed:", err)
		}

		if world.Name != w.World.Name {
			t.Error("Invalid world name")
		}

		if world.Seed != w.World.Seed {
			t.Error("Invalid seed")
		}

		if world.SeedValue != w.World.SeedValue {
			t.Error("Invalid seed value", world.SeedValue, w.World.SeedValue)
		}

		if world.UID != w.World.UID {
			t.Error("Invalid uid")
		}

		if world.WorldVersion != w.World.WorldVersion {
			t.Error("Invalid world version")
		}

		if world.WorldGenVersion != w.World.WorldGenVersion {
			t.Error("Invalid world gen version")
		}
	}
}

func TestWorld_MarshalBinary(t *testing.T) {
	for _, w := range testData {

		world := fwl.World{}

		err := world.UnmarshalBinary(w.RawData)
		if err != nil {
			t.Error("UnmashalBinary failed:", err)
		}

		if world.Name != w.World.Name {
			t.Error("Invalid world name")
		}

		if world.Seed != w.World.Seed {
			t.Error("Invalid seed")
		}

		if world.SeedValue != w.World.SeedValue {
			t.Error("Invalid seed value", world.SeedValue, w.World.SeedValue)
		}

		if world.UID != w.World.UID {
			t.Error("Invalid uid")
		}

		if world.WorldVersion != w.World.WorldVersion {
			t.Error("Invalid world version")
		}

		if world.WorldGenVersion != w.World.WorldGenVersion {
			t.Error("Invalid world gen version")
		}

	}
}
