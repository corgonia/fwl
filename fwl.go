// Package fwl provides Valheim .fwl world metadata file import/export
package fwl

import (
	"bytes"
	"encoding/binary"
	"math/rand"
	"os"
	"time"
)

var (
	// Based on the latest version of Valheim on 20 March 2021

	// DefaultWorldVersion is the default value for world file version
	DefaultWorldVersion int32 = 26

	// DefaultWorldGenVersion is the default value for world generation type
	DefaultWorldGenVersion int32 = 1
)

// World represents a Valheim game world's metadata as stored in the .fwl file
type World struct {
	Name            string // Name of the world. Must match the name of the .fwl and .db files
	Seed            string // Human readable ASCII seed provided as input for world generation
	UID             int64  // Unique identifier for this world used to sync player state in clients
	SeedValue       int32  // Actual int32 value of the seed that Valheim uses
	WorldVersion    int32  // Version of the world file
	WorldGenVersion int32  // Version of the world generation
}

// MarshalBinary converts a World into binary form suitable for being written to a .fwl file
func (w *World) MarshalBinary() (data []byte, err error) {

	var out = bytes.Buffer{}

	if w.WorldVersion == 0 {
		w.WorldVersion = DefaultWorldVersion
	}
	err = binary.Write(&out, binary.LittleEndian, w.WorldVersion)
	if err != nil {
		return nil, err
	}

	length := int8(len(w.Name))
	err = binary.Write(&out, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}
	err = binary.Write(&out, binary.LittleEndian, []byte(w.Name))
	if err != nil {
		return nil, err
	}

	length = int8(len(w.Seed))
	err = binary.Write(&out, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}
	err = binary.Write(&out, binary.LittleEndian, []byte(w.Seed))
	if err != nil {
		return nil, err
	}

	seedValue := GetStableHashCode(w.Seed)

	err = binary.Write(&out, binary.LittleEndian, seedValue)
	if err != nil {
		return nil, err
	}
	if w.UID == 0 {
		r := rand.NewSource(time.Now().UnixNano())
		w.UID = r.Int63()
	}

	//  UID = w.UID
	err = binary.Write(&out, binary.LittleEndian, w.UID)
	if err != nil {
		return nil, err
	}

	if w.WorldGenVersion == 0 {
		w.WorldGenVersion = DefaultWorldGenVersion
	}
	err = binary.Write(&out, binary.LittleEndian, w.WorldGenVersion)
	if err != nil {
		return nil, err
	}

	// todo: fix this crappy append stuff to get the content length at the front of the slice
	var temp bytes.Buffer
	err = binary.Write(&temp, binary.LittleEndian, int32(out.Len()))
	if err != nil {
		return nil, err
	}
	result := append(temp.Bytes(), out.Bytes()...)
	//	fmt.Printf("% x\n", result)
	return result, nil

}

// UnmarshalBinary reads raw binary data from a .fwl file and populates a World
func (w *World) UnmarshalBinary(data []byte) error {
	buf := bytes.NewBuffer(data)
	var dataLen int32
	var nameLen int8
	var name []byte
	var seedNameLen int8
	var seedName []byte
	//var seed int32
	//	var UID int64
	//var WorldGenVersion int32

	var headerData = []interface{}{&dataLen, &w.WorldVersion, &nameLen}
	for _, d := range headerData {
		err := binary.Read(buf, binary.LittleEndian, d)
		if err != nil {
			return err
		}
	}

	//	fmt.Println("Data length:", dataLen)
	name = make([]byte, nameLen)
	err := binary.Read(buf, binary.LittleEndian, &name)
	if err != nil {
		return err
	}

	w.Name = string(name)
	err = binary.Read(buf, binary.LittleEndian, &seedNameLen)
	if err != nil {
		return err
	}

	seedName = make([]byte, seedNameLen)
	err = binary.Read(buf, binary.LittleEndian, &seedName)
	if err != nil {
		return err
	}
	w.Seed = string(seedName)
	err = binary.Read(buf, binary.LittleEndian, &w.SeedValue)
	if err != nil {
		return err
	}

	err = binary.Read(buf, binary.LittleEndian, &w.UID)
	if err != nil {
		return err
	}

	err = binary.Read(buf, binary.LittleEndian, &w.WorldGenVersion)
	if err != nil {
		return err
	}

	return nil
}

// NewWorld creates a new World with the specified Name and Seed
func NewWorld(name, seed string) *World {
	w := World{
		WorldVersion:    DefaultWorldVersion,
		Name:            name,
		Seed:            seed,
		SeedValue:       GetStableHashCode(seed),
		WorldGenVersion: DefaultWorldGenVersion,
	}

	w.UID = w.GenerateUID()

	return &w
}

// GetStableHashCode is used to convert strings to int32's. Most notably for hashing supplied Seed values.
func GetStableHashCode(str string) int32 {
	var hash1 int32 = 5381
	hash2 := hash1

	for i := 0; i < len(str) && str[i] != '\x00'; i += 2 {
		hash1 = ((hash1 << 5) + hash1) ^ int32(str[i])
		if i == len(str)-1 || str[i+1] == '\x00' {
			break
		}
		hash2 = ((hash2 << 5) + hash2) ^ int32(str[i+1])
	}

	return hash1 + (hash2 * 1566083941)
}

// GenerateUID creates a random UID based on the World information
func (w *World) GenerateUID() int64 {
	// Create a new seeded random source
	randSource := rand.NewSource(time.Now().UnixNano())
	r := rand.New(randSource)

	// Try to incorporate the system's hostname in the UID like Valheim does
	hostName, _ := os.Hostname()
	if hostName == "" {
		hostName = "unknown"
	}

	// Hash the host name, world name, and world seed to a unique value
	hash := GetStableHashCode(hostName + ":" + w.Name + ":" + w.Seed)

	var result int64

	// Shift it into the left part of a 64 bit int
	result = int64(hash) << 32

	// Set the right part to some random numbers
	result = result + int64(r.Int31())

	return result
}
