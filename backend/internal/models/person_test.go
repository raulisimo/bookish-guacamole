package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestPerson_JSONSerialization(t *testing.T) {
	person := Person{
		Name:        "Luke Skywalker",
		Height:      "172",
		Mass:        "77",
		HairColor:   "blond",
		SkinColor:   "fair",
		EyeColor:    "blue",
		BirthYear:   "19BBY",
		Gender:      "male",
		Homeworld:   "https://swapi.dev/api/planets/1/",
		Films:       []string{"https://swapi.dev/api/films/1/"},
		Species:     []string{},
		Vehicles:    []string{"https://swapi.dev/api/vehicles/14/"},
		Starships:   []string{"https://swapi.dev/api/starships/12/"},
		Created:     "2014-12-09T13:50:51.644000Z",
		Edited:      "2014-12-20T21:17:56.891000Z",
		URL:         "https://swapi.dev/api/people/1/",
	}

	// Serialize to JSON
	data, err := json.Marshal(person)
	if err != nil {
		t.Fatalf("Failed to serialize Person to JSON: %v", err)
	}

	// Deserialize back to a Person object
	var deserialized Person
	err = json.Unmarshal(data, &deserialized)
	if err != nil {
		t.Fatalf("Failed to deserialize JSON to Person: %v", err)
	}

	// Check if the original and deserialized Person are the same
	if !reflect.DeepEqual(person, deserialized) {
		t.Errorf("Deserialized Person does not match original.\nExpected: %+v\nGot: %+v", person, deserialized)
	}
}
