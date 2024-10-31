package cache_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/capcom6/sms-gateway/pkg/types/cache"
)

func TestCache_Set(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		value    any
		ttl      time.Duration
		wait     time.Duration
		expected any
	}{
		{
			name:     "Set string value w/o TTL",
			key:      "key1",
			value:    "value1",
			ttl:      0,
			wait:     0,
			expected: "value1",
		},
		{
			name:     "Set string value",
			key:      "key1",
			value:    "value1",
			ttl:      10 * time.Second,
			wait:     0,
			expected: "value1",
		},
		{
			name:     "Set integer value",
			key:      "key2",
			value:    123,
			ttl:      10 * time.Second,
			wait:     0,
			expected: 123,
		},
		{
			name:     "Set value with expired TTL",
			key:      "key3",
			value:    "value3",
			ttl:      1 * time.Second,
			wait:     2 * time.Second,
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache := cache.New[any](cache.Config{TTL: tt.ttl})
			if err := cache.Set(tt.key, tt.value); err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if tt.wait > 0 {
				time.Sleep(tt.wait)
			}

			result, err := cache.Get(tt.key)
			if tt.expected == nil {
				if err == nil {
					t.Errorf("Expected error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if result != tt.expected {
					t.Errorf("Expected %v, got %v", tt.expected, result)
				}
			}
		})
	}
}

func TestCache_Drain(t *testing.T) {
	tests := []struct {
		name     string
		setup    map[string]any
		ttl      time.Duration
		wait     time.Duration
		expected map[string]any
	}{
		{
			name: "Drain non-expired items",
			setup: map[string]any{
				"key1": "value1",
				"key2": 42,
				"key3": true,
			},
			ttl:  10 * time.Second,
			wait: 0,
			expected: map[string]any{
				"key1": "value1",
				"key2": 42,
				"key3": true,
			},
		},
		{
			name: "Drain with some expired items",
			setup: map[string]any{
				"key1": "value1",
				"key2": 42,
				"key3": true,
			},
			ttl:      1 * time.Second,
			wait:     2 * time.Second,
			expected: map[string]any{}, // All items should be expired
		},
		{
			name:     "Drain empty cache",
			setup:    map[string]any{},
			ttl:      10 * time.Second,
			wait:     0,
			expected: map[string]any{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache := cache.New[any](cache.Config{TTL: tt.ttl})

			// Setup cache
			for k, v := range tt.setup {
				err := cache.Set(k, v)
				if err != nil {
					t.Fatalf("Failed to set up cache: %v", err)
				}
			}

			// Wait if specified
			if tt.wait > 0 {
				time.Sleep(tt.wait)
			}

			// Drain cache
			result := cache.Drain()

			// Check result
			if len(result) != len(tt.expected) {
				t.Errorf("Expected %d items, got %d", len(tt.expected), len(result))
			}

			for k, v := range tt.expected {
				if rv, ok := result[k]; !ok || rv != v {
					t.Errorf("Expected %v for key %s, got %v", v, k, rv)
				}
			}

			// Verify cache is empty after drain
			if len(cache.Drain()) != 0 {
				t.Errorf("Cache should be empty after Drain")
			}
		})
	}
}

func TestCache_Cleanup(t *testing.T) {
	tests := []struct {
		name     string
		setup    map[string]any
		ttl      time.Duration
		wait     time.Duration
		expected map[string]any
	}{
		{
			name: "No expired items",
			setup: map[string]any{
				"key1": "value1",
				"key2": 42,
				"key3": true,
			},
			ttl:  10 * time.Second,
			wait: 0,
			expected: map[string]any{
				"key1": "value1",
				"key2": 42,
				"key3": true,
			},
		},
		{
			name: "Some expired items",
			setup: map[string]any{
				"key1": "value1",
				"key2": 42,
				"key3": true,
			},
			ttl:      1 * time.Second,
			wait:     2 * time.Second,
			expected: map[string]any{},
		},
		{
			name:     "Empty cache",
			setup:    map[string]any{},
			ttl:      10 * time.Second,
			wait:     0,
			expected: map[string]any{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache := cache.New[any](cache.Config{TTL: tt.ttl})

			// Setup cache
			for k, v := range tt.setup {
				err := cache.Set(k, v)
				if err != nil {
					t.Fatalf("Failed to set up cache: %v", err)
				}
			}

			// Wait if specified
			if tt.wait > 0 {
				time.Sleep(tt.wait)
			}

			// Run cleanup
			cache.Cleanup()

			// Check remaining items
			result := cache.Drain()

			if len(result) != len(tt.expected) {
				t.Errorf("Expected %d items after cleanup, got %d", len(tt.expected), len(result))
			}

			for k, v := range tt.expected {
				if rv, ok := result[k]; !ok || rv != v {
					t.Errorf("Expected %v for key %s after cleanup, got %v", v, k, rv)
				}
			}
		})
	}
}

///

func BenchmarkCache_Get(b *testing.B) {
	cache := cache.New[string](cache.Config{TTL: 1 * time.Hour})

	// Prepare the cache with some data
	testKey := "testKey"
	testValue := "testValue"
	err := cache.Set(testKey, testValue)
	if err != nil {
		b.Fatalf("Failed to set up cache: %v", err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		value, err := cache.Get(testKey)
		if err != nil {
			b.Fatalf("Unexpected error: %v", err)
		}
		if value != testValue {
			b.Fatalf("Unexpected value: got %v, want %v", value, testValue)
		}
	}
}

func BenchmarkCache_GetNonExistent(b *testing.B) {
	c := cache.New[string](cache.Config{TTL: 1 * time.Hour})

	nonExistentKey := "nonExistentKey"

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := c.Get(nonExistentKey)
		if err != cache.ErrKeyNotFound {
			b.Fatalf("Unexpected error: got %v, want %v", err, cache.ErrKeyNotFound)
		}
	}
}

func BenchmarkCache_GetConcurrent(b *testing.B) {
	cache := cache.New[string](cache.Config{TTL: 1 * time.Hour})

	// Prepare the cache with some data
	testKey := "testKey"
	testValue := "testValue"
	err := cache.Set(testKey, testValue)
	if err != nil {
		b.Fatalf("Failed to set up cache: %v", err)
	}

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			value, err := cache.Get(testKey)
			if err != nil {
				b.Fatalf("Unexpected error: %v", err)
			}
			if value != testValue {
				b.Fatalf("Unexpected value: got %v, want %v", value, testValue)
			}
		}
	})
}

func BenchmarkCache_Set(b *testing.B) {
	cache := cache.New[string](cache.Config{TTL: 1 * time.Hour})

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("key%d", i)
		value := fmt.Sprintf("value%d", i)
		err := cache.Set(key, value)
		if err != nil {
			b.Fatalf("Unexpected error: %v", err)
		}
	}
}

func BenchmarkCache_SetExisting(b *testing.B) {
	cache := cache.New[string](cache.Config{TTL: 1 * time.Hour})

	// Prepare the cache with initial data
	testKey := "testKey"
	initialValue := "initialValue"
	err := cache.Set(testKey, initialValue)
	if err != nil {
		b.Fatalf("Failed to set up cache: %v", err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		newValue := fmt.Sprintf("value%d", i)
		err := cache.Set(testKey, newValue)
		if err != nil {
			b.Fatalf("Unexpected error: %v", err)
		}
	}
}

func BenchmarkCache_SetConcurrent(b *testing.B) {
	cache := cache.New[string](cache.Config{TTL: 1 * time.Hour})

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			key := fmt.Sprintf("key%d", i)
			value := fmt.Sprintf("value%d", i)
			err := cache.Set(key, value)
			if err != nil {
				b.Fatalf("Unexpected error: %v", err)
			}
			i++
		}
	})
}
