package fwl

import "testing"

func TestWorld_GenerateUID(t *testing.T) {
	var seed int64
	var hostname = "hostname"
	w := World{
		Name:            "testing",
		Seed:            "testing",
		SeedValue:       GetStableHashCode("testing"),
		WorldGenVersion: DefaultWorldGenVersion,
		WorldVersion:    DefaultWorldVersion,
	}

	if w.generateUID(hostname, seed) != -3088824465891774470 {
		t.Error("Invalid UID generated")
	}
}
