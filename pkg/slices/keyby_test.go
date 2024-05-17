package slices

import (
	"reflect"
	"testing"
)

func TestKeyBy(t *testing.T) {
	// Test when the input slice is empty
	t.Run("Empty Slice", func(t *testing.T) {
		got := KeyBy([]int{}, func(x int) int { return x })
		want := map[int]int{}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected %v, got %v", want, got)
		}
	})

	// Test when the mapping function returns duplicate keys
	t.Run("Duplicate Keys", func(t *testing.T) {
		got := KeyBy([]int{1, 2, 3}, func(x int) int { return 1 })
		want := map[int]int{1: 3}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected %v, got %v", want, got)
		}
	})

	// Test when the mapping function returns unique keys
	t.Run("Unique Keys", func(t *testing.T) {
		got := KeyBy([]int{1, 2, 3}, func(x int) int { return x })
		want := map[int]int{1: 1, 2: 2, 3: 3}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected %v, got %v", want, got)
		}
	})
}
