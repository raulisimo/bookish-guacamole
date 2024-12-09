package sorting

import (
	"fmt"
)

// SorterFactory creates a sorter based on the "sort" field
// If the field is unsupported, it returns nil to indicate no sorting should be applied
func SorterFactory(sortField string) (Sorter, error) {
	switch sortField {
	case "name":
		return &NameSorter{}, nil
	case "created":
		return &CreatedSorter{}, nil
	case "":
		// If no "sort" parameter is provided, return nil, indicating no sorting
		return nil, nil
	default:
		// If an invalid "sort" field is provided, return nil with an error
		return nil, fmt.Errorf("unsupported field: %s", sortField)
	}
}
