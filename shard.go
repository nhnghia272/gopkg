package gopkg

import (
	"hash/fnv"
	"sync"
	"time"

	"github.com/dgraph-io/ristretto"
)

type CacheShard []*cacheObj

type cacheObj struct {
	sync.RWMutex
	items *ristretto.Cache
}

func NewCacheShard(shard uint64) CacheShard {
	m := make(CacheShard, shard)
	for i := 0; i < int(shard); i++ {
		items, _ := ristretto.NewCache(&ristretto.Config{NumCounters: 1_000_000, MaxCost: 1_000_000, BufferItems: 64})
		m[i] = &cacheObj{items: items}
	}
	return m
}

func (s CacheShard) getShard(key string) *cacheObj {
	h := fnv.New32a()
	h.Write([]byte(key))
	return s[uint(h.Sum32())%uint(len(s))]
}

func (s CacheShard) Set(key string, val any, exp time.Duration) {
	shard := s.getShard(key)
	shard.Lock()
	defer shard.Unlock()
	shard.items.SetWithTTL(key, val, 1, exp)
}

func (s CacheShard) Get(key string) (any, bool) {
	shard := s.getShard(key)
	shard.RLock()
	defer shard.RUnlock()
	return shard.items.Get(key)
}

func (s CacheShard) Del(key string) {
	shard := s.getShard(key)
	shard.Lock()
	defer shard.Unlock()
	shard.items.Del(key)
}

type MapShard []*mapObj

type mapObj struct {
	sync.RWMutex
	items map[string]any
}

func NewMapShard(shard uint64) MapShard {
	m := make(MapShard, shard)
	for i := 0; i < int(shard); i++ {
		m[i] = &mapObj{items: make(map[string]any)}
	}
	return m
}

func (s MapShard) getShard(key string) *mapObj {
	h := fnv.New32a()
	h.Write([]byte(key))
	return s[uint(h.Sum32())%uint(len(s))]
}

func (s MapShard) Set(key string, val any) {
	shard := s.getShard(key)
	shard.Lock()
	defer shard.Unlock()
	shard.items[key] = val
}

func (s MapShard) Get(key string) (any, bool) {
	shard := s.getShard(key)
	shard.RLock()
	defer shard.RUnlock()
	val, ok := shard.items[key]
	return val, ok
}

func (s MapShard) Del(key string) {
	shard := s.getShard(key)
	shard.Lock()
	defer shard.Unlock()
	delete(shard.items, key)
}

func (s MapShard) List() []any {
	items := make([]any, 0)
	for _, shard := range s {
		shard.RLock()
		defer shard.RUnlock()
		for _, item := range shard.items {
			items = append(items, item)
		}
	}
	return items
}
