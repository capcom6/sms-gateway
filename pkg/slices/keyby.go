package slices

// KeyBy generates a map by associating keys from the given slice using the provided function.
//
// slice: The input slice of type T.
// f: The key function that maps elements of type T to keys of type K.
// Returns a map with keys of type K and values of type T.
func KeyBy[T any, K comparable](slice []T, f func(T) K) map[K]T {
	result := make(map[K]T, len(slice))
	for _, v := range slice {
		result[f(v)] = v
	}
	return result
}
