package sorting

// Sorter is the interface for sorting strategies
type Sorter interface {
	Sort(data interface{}, order string) error
}



