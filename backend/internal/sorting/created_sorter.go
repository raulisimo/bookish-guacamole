package sorting

import (
	"fmt"
	"reflect"
	"sort"
	"time"
)

// CreatedSorter sorts by the "created" field
type CreatedSorter struct{}

// Sort sorts data by the created field in ascending or descending order
func (s *CreatedSorter) Sort(data interface{}, order string) error {
	// Ensure data is a slice of structs that contains a "Created" field
	v := reflect.ValueOf(data)
	if v.Kind() != reflect.Slice {
		return fmt.Errorf("data must be a slice")
	}

	// Check if the elements of the slice have a "Created" field
	if v.Len() == 0 || !v.Index(0).FieldByName("Created").IsValid() {
		return fmt.Errorf("data elements must have a 'Created' field")
	}

	sort.SliceStable(data, func(i, j int) bool {
		created1 := reflect.ValueOf(data).Index(i).FieldByName("Created").String()
		created2 := reflect.ValueOf(data).Index(j).FieldByName("Created").String()

		t1, err1 := time.Parse(time.RFC3339, created1)
		t2, err2 := time.Parse(time.RFC3339, created2)
		if err1 != nil || err2 != nil {
			return false
		}

		switch order {
		case "asc":
			return t1.Before(t2)
		case "desc":
			return t1.After(t2)
		default:
			return false
		}
	})

	return nil
}
