package cache

import (
	"testing"
)

func TestCacheSetAndGet(t *testing.T) {
	cache := New[string]()

	key := "myKey"
	value := "myValue"

	cache.Set(key, value)

	if cache.storage[key] != value {
		t.Errorf("Set failed: expected value %v, got %v", value, cache.storage[key])
	}
}

func TestCacheDrain(t *testing.T) {
	cache := New[string]()

	cache.Set("key1", "value1")
	cache.Set("key2", "value2")

	drained := cache.Drain()

	if len(drained) != 2 || drained["key1"] != "value1" || drained["key2"] != "value2" {
		t.Errorf("Drain failed: expected map[key1:value1 key2:value2], got %v", drained)
	}

	if len(cache.storage) != 0 {
		t.Errorf("Drain failed: expected empty cache storage, got %v", cache.storage)
	}
}
