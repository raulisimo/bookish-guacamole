package sorting

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

// NameSorter sorts by the "name" field.
type NameSorter struct{}

// Sort sorts data by name in ascending or descending order.
func (s *NameSorter) Sort(data interface{}, order string) error {
	// Ensure data is a slice of structs that contains a "Name" field
	v := reflect.ValueOf(data)
	if v.Kind() != reflect.Slice {
		return fmt.Errorf("data must be a slice")
	}

	// Check if the elements of the slice have a "Name" field
	if v.Len() == 0 || !v.Index(0).FieldByName("Name").IsValid() {
		return fmt.Errorf("data elements must have a 'Name' field")
	}

	sort.SliceStable(data, func(i, j int) bool {
		name1 := reflect.ValueOf(data).Index(i).FieldByName("Name").String()
		name2 := reflect.ValueOf(data).Index(j).FieldByName("Name").String()

		switch order {
		case "asc":
			return strings.ToLower(name1) < strings.ToLower(name2)
		case "desc":
			return strings.ToLower(name1) > strings.ToLower(name2)
		default:
			return false
		}
	})

	return nil
}
