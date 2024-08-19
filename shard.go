package gopkg

import (
	"hash/fnv"
	"sync"
	"time"
)

type CacheShard []*cache

type cache struct {
	sync.RWMutex
	items   map[string]any
	expires map[string]time.Time
}

func NewCacheShard(shard uint64) CacheShard {
	m := make(CacheShard, shard)
	for i := 0; i < int(shard); i++ {
		m[i] = &cache{
			items:   make(map[string]any),
			expires: make(map[string]time.Time),
		}
		Async().Go(func() {
			for {
				m[i].clean()
				time.Sleep(15 * time.Second)
			}
		})
	}
	return m
}

func (s *cache) clean() {
	s.Lock()
	defer s.Unlock()
	keys := []string{}
	for key, exp := range s.expires {
		if exp.Before(time.Now()) {
			keys = append(keys, key)
		}
	}
	for _, key := range keys {
		delete(s.items, key)
		delete(s.expires, key)
	}
}

func (s CacheShard) getShard(key string) *cache {
	h := fnv.New32a()
	h.Write([]byte(key))
	return s[uint(h.Sum32())%uint(len(s))]
}

func (s CacheShard) Set(key string, val any, exp time.Duration) {
	shard := s.getShard(key)
	shard.Lock()
	defer shard.Unlock()
	shard.items[key] = val
	shard.expires[key] = time.Now().Add(exp)
}

func (s CacheShard) Get(key string) (any, bool) {
	shard := s.getShard(key)
	shard.RLock()
	defer shard.RUnlock()

	val, ok := shard.items[key]
	if !ok {
		return nil, false
	}
	if exp := shard.expires[key]; exp.Before(time.Now()) {
		delete(shard.items, key)
		delete(shard.expires, key)
		return nil, false
	}

	return val, ok
}

func (s CacheShard) Del(key string) {
	shard := s.getShard(key)
	shard.Lock()
	defer shard.Unlock()
	delete(shard.items, key)
	delete(shard.expires, key)
}
