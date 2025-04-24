package main

import (
	"container/list"
	"sync"
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

	lock sync.RWMutex
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
	// added read locks
	k.lock.RLock()
	elem, ok := k.cache[key]
	k.lock.RUnlock()

	if ok {
		// added write locks
		k.lock.Lock()
		k.pages.MoveToFront(elem)
		k.lock.Unlock()
		return elem.Value.(Node).Value
	}

	// Load from DB
	// not added locks, load(key) is slow (DB access). If we lock 
	// around it, we block all other goroutines (even readers).
	value := k.load(key)
	node := Node{Key: key, Value: value}

	k.lock.Lock()
	defer k.lock.Unlock()

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
