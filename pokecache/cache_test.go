package pokecache

import (
	"bytes"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cacheValues := []struct {
		key   string
		value []byte
	}{
		{
			key:   "url-1",
			value: []byte("value-1"),
		},
		{
			key:   "url-2",
			value: []byte("vealue-2"),
		},
	}

	cases := []struct {
		input    string
		expected []byte
		ok       bool
	}{
		{
			input:    "url-1",
			expected: []byte("value-1"),
			ok:       true,
		},
		{
			input:    "url-2",
			expected: []byte("vealue-2"),
			ok:       true,
		},
		{
			input:    "url-3",
			expected: []byte{},
			ok:       false,
		},
	}

	cache := Cache{Cache: make(map[string]CacheEntry)}

	for _, cacheValue := range cacheValues {
		cache.Add(cacheValue.key, cacheValue.value)
	}

	for _, c := range cases {
		cacheValue, ok := cache.Get(c.input)

		if !bytes.Equal(cacheValue, c.expected) {
			t.Errorf("Expecting %q as cache value for %s, got %q", string(c.expected), c.input, string(cacheValue))
			continue
		}

		if ok != c.ok {
			t.Errorf("Expecting %t as bool value for %s, got %t", c.ok, c.input, ok)
			continue
		}

	}
}
