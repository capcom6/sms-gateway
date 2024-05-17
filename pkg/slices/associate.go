package slices

// Associate generates a map by associating keys and values from the given slice using the provided key and value functions.
//
// slice: The input slice of type T.
// fk: The key function that maps elements of type T to keys of type K.
// fv: The value function that maps elements of type T to values of type V.
// Returns a map with keys of type K and values of type V.
func Associate[T any, K comparable, V any](slice []T, fk func(T) K, fv func(T) V) map[K]V {
	result := make(map[K]V, len(slice))
	for _, v := range slice {
		result[fk(v)] = fv(v)
	}
	return result
}
