package sorting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Person struct {
	Name string
}

func TestNameSorter(t *testing.T) {
	sorter := &NameSorter{}

	tests := []struct {
		name           string
		data           interface{}
		order          string
		expectedResult interface{}
		expectedError  string
	}{
		{
			name:   "Success - Ascending order",
			data:   []Person{{Name: "Luke"}, {Name: "Leia"}, {Name: "Han"}},
			order:  "asc",
			expectedResult: []Person{
				{Name: "Han"},
				{Name: "Leia"},
				{Name: "Luke"},
			},
			expectedError: "",
		},
		{
			name:   "Success - Descending order",
			data:   []Person{{Name: "Luke"}, {Name: "Leia"}, {Name: "Han"}},
			order:  "desc",
			expectedResult: []Person{
				{Name: "Luke"},
				{Name: "Leia"},
				{Name: "Han"},
			},
			expectedError: "",
		},
		{
			name:          "Error - Invalid data type",
			data:          "invalid string", // Not a slice
			order:         "asc",
			expectedError: "data must be a slice",
		},
		{
			name:          "Error - No Name field",
			data:          []struct{ Age int }{{Age: 30}, {Age: 25}}, // Missing 'Name' field
			order:         "asc",
			expectedError: "data elements must have a 'Name' field",
		},
		{
			name:   "Error - Invalid order",
			data:   []Person{{Name: "Luke"}, {Name: "Leia"}, {Name: "Han"}},
			order:  "invalid", // Invalid order
			expectedError:    "invalid order: must be 'asc' or 'desc'",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Making a copy of the data for comparison after sorting
			dataCopy := tt.data
			err := sorter.Sort(dataCopy, tt.order)

			if tt.expectedError == "" {
				// No error expected, check if data matches the expected result
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResult, dataCopy)
			} else {
				// Error expected, check for error and ensure no sorting occurred
				assert.EqualError(t, err, tt.expectedError)
			}
		})
	}
}
