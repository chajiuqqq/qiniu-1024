package utils

// Mapable is an interface for types that can provide a unique key
type Mapable interface {
	Key() any
}

// ToMap is a generic function that converts a slice of Keyer elements to a map
func ToMap[T Mapable](slice []T) map[any]T {
	result := make(map[any]T)
	for _, item := range slice {
		result[item.Key()] = item
	}
	return result
}
