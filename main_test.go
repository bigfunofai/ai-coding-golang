package main

import (
	"testing"
)

func TestLRUCache(t *testing.T) {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	if res := cache.Get(1); res != 1 {
		t.Errorf("expected 1, got %d", res)
	}
	cache.Put(3, 3)
	if res := cache.Get(2); res != -1 {
		t.Errorf("expected -1, got %d", res)
	}
	cache.Put(4, 4)
	if res := cache.Get(1); res != -1 {
		t.Errorf("expected -1, got %d", res)
	}
	if res := cache.Get(3); res != 3 {
		t.Errorf("expected 3, got %d", res)
	}
	if res := cache.Get(4); res != 4 {
		t.Errorf("expected 4, got %d", res)
	}
}