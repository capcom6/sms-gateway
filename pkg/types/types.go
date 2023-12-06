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
