package maps

func MapValues[K comparable, V any, R any](m map[K]V, f func(V) R) map[K]R {
	result := make(map[K]R, len(m))
	for k, v := range m {
		result[k] = f(v)
	}
	return result
}
