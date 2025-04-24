package main

import (
	"container/list"
	"testing"
)

const CacheSize = 100

// This struct holds one key-value entry
type Node struct {
	Key   string
	Value string
}

// This is the main cache structure
type KeyStoreCache struct {
	cache map[string]*list.Element // for fast key lookup
	pages list.List                // to maintain order for LRU
	load  func(string) string      // function to load from DB
}

// This struct simulates a loader which has DB
type Loader struct {
	DB *MockDB
}

// This method is used to load value from database
func (l *Loader) Load(key string) string {
	val, err := l.DB.Get(key)
	if err != nil {
		panic(err)
	}
	return val
}

// This function is used to create a new cache
func New(loader *Loader) *KeyStoreCache {
	return &KeyStoreCache{
		cache: make(map[string]*list.Element),
		load:  loader.Load,
	}
}

// This function is used to get value for a key from cache
func (k *KeyStoreCache) Get(key string) string {
	elem, ok := k.cache[key]

	if ok {
		k.pages.MoveToFront(elem)
		return elem.Value.(Node).Value
	}

	value := k.load(key)
	node := Node{Key: key, Value: value}

	// Evict least recently used if full
	if len(k.cache) >= CacheSize {
		last := k.pages.Back()
		if last != nil {
			delete(k.cache, last.Value.(Node).Key)
			k.pages.Remove(last)
		}
	}

	// Add new item to front
	k.pages.PushFront(node)
	k.cache[key] = k.pages.Front()

	return value
}

// This is the main runner for mock testing
func run(t *testing.T) (*KeyStoreCache, *MockDB) {
	loader := &Loader{
		DB: GetMockDB(),
	}
	cache := New(loader)
	RunMockServer(cache, t)
	return cache, loader.DB
}

func main() {
	run(nil)
}
