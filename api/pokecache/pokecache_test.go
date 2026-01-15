package pokecache

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	tests := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
		{
			key: "https://example.com/path?offset=2&limit=2",
			val: []byte("testdatawithparameter"),
		},
		{
			key: "https://example.com/path?limit=2",
			val: []byte("moretestdatawithparameter"),
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			cache := NewCache(5 * time.Second)
			cache.Add(test.key, test.val)
			val, _ := cache.Get(test.key)
			if !reflect.DeepEqual(test.val, val) {
				t.Fatalf("expected: %v, got: %v", test.val, val)
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 8 * time.Millisecond
	const waitTime = baseTime + 2*time.Millisecond

	cache := NewCache(baseTime)
	cache.Add("https://example.com", []byte("reaptestdata"))
	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("no val found")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("Val found")
		return
	}
}
