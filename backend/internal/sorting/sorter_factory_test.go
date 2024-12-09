package sorting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSorterFactory(t *testing.T) {
	tests := []struct {
		name           string
		sortField      string
		expectedSorter Sorter
		expectedError  string
	}{
		{
			name:           "Valid SortField - 'name'",
			sortField:      "name",
			expectedSorter: &NameSorter{},
			expectedError:  "",
		},
		{
			name:           "Valid SortField - 'created'",
			sortField:      "created",
			expectedSorter: &CreatedSorter{},
			expectedError:  "",
		},
		{
			name:           "Empty SortField",
			sortField:      "",
			expectedSorter: nil,
			expectedError:  "",
		},
		{
			name:           "Invalid SortField",
			sortField:      "invalid",
			expectedSorter: nil,
			expectedError:  "unsupported field: invalid",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sorter, err := SorterFactory(tt.sortField)

			if tt.expectedError == "" {
				assert.NoError(t, err)
				assert.IsType(t, tt.expectedSorter, sorter)
			} else {
				assert.EqualError(t, err, tt.expectedError)
				assert.Nil(t, sorter)
			}
		})
	}
}
