package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestPlanet_JSONSerialization(t *testing.T) {
	planet := Planet{
		Name:           "Tatooine",
		RotationPeriod: "23",
		OrbitalPeriod:  "304",
		Diameter:       "10465",
		Climate:        "arid",
		Gravity:        "1 standard",
		Terrain:        "desert",
		SurfaceWater:   "1",
		Population:     "200000",
		Residents:      []string{"https://swapi.dev/api/people/1/", "https://swapi.dev/api/people/2/"},
		Films:          []string{"https://swapi.dev/api/films/1/", "https://swapi.dev/api/films/3/"},
		Created:        "2014-12-09T13:50:49.641000Z",
		Edited:         "2014-12-20T20:58:18.411000Z",
		URL:            "https://swapi.dev/api/planets/1/",
	}

	// Serialize to JSON
	data, err := json.Marshal(planet)
	if err != nil {
		t.Fatalf("Failed to serialize Planet to JSON: %v", err)
	}

	// Deserialize back to a Planet object
	var deserialized Planet
	err = json.Unmarshal(data, &deserialized)
	if err != nil {
		t.Fatalf("Failed to deserialize JSON to Planet: %v", err)
	}

	// Compare original and deserialized struct
	if !reflect.DeepEqual(planet, deserialized) {
		t.Errorf("Deserialized Planet does not match original.\nExpected: %+v\nGot: %+v", planet, deserialized)
	}
}
