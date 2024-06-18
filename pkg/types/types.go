package types

func AsPointer[T any](v T) *T {
	return &v
}

func OrDefault[T any](v *T, def T) T {
	if v == nil {
		return def
	}
	return *v
}

// ZeroDefault returns the default value if the given value is zero, otherwise it returns the value.
//
// Parameters:
// - v: The value to check.
// - def: The default value to return if v is zero.
//
// Returns:
// - The default value if v is zero, otherwise the value.
func ZeroDefault[T comparable](v T, def T) T {
	zero := new(T)
	if v == *zero {
		return def
	}
	return v
}
