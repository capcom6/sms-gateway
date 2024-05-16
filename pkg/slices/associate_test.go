package slices

import (
	"reflect"
	"testing"
)

func TestAssociate(t *testing.T) {
	// Test when the input slice is empty
	t.Run("Empty Slice", func(t *testing.T) {
		got := Associate([]int{}, func(x int) int { return x }, func(x int) int { return x })
		want := map[int]int{}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected %v, got %v", want, got)
		}
	})

	// Test when the fk function returns duplicate keys
	t.Run("Duplicate Keys", func(t *testing.T) {
		got := Associate([]int{1, 2, 3}, func(x int) int { return 1 }, func(x int) int { return x })
		want := map[int]int{1: 3}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected %v, got %v", want, got)
		}
	})

	// Test when the fv function returns duplicate values
	t.Run("Duplicate Values", func(t *testing.T) {
		got := Associate([]int{1, 2, 3}, func(x int) int { return x }, func(x int) int { return 1 })
		want := map[int]int{1: 1, 2: 1, 3: 1}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected %v, got %v", want, got)
		}
	})
}
