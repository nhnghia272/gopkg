package gopkg

import (
	"errors"
	"hash/fnv"
	"sync"
	"time"
)

type CacheShard[E any] []*cache[E]

type cache[E any] struct {
	sync.RWMutex
	items   map[string]E
	expires map[string]time.Time
}

func (s *cache[E]) clean() {
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

type CacheConfig struct {
	Shard int64
	Clean time.Duration
}

func cacheConfigDefault(config ...CacheConfig) CacheConfig {
	if len(config) < 1 {
		return CacheConfig{Shard: 1, Clean: 30 * time.Second}
	}
	cfg := config[0]
	if cfg.Shard <= 0 {
		cfg.Shard = 1
	}
	if int64(cfg.Clean) <= 0 {
		cfg.Clean = 30 * time.Second
	}
	return cfg
}

func NewCacheShard[E any](config ...CacheConfig) CacheShard[E] {
	var (
		as     = Async()
		cfg    = cacheConfigDefault(config...)
		shards = make(CacheShard[E], cfg.Shard)
	)

	for i := 0; i < int(cfg.Shard); i++ {
		shards[i] = &cache[E]{
			items:   make(map[string]E),
			expires: make(map[string]time.Time),
		}
		as.Go(func() {
			for {
				shards[i].clean()
				time.Sleep(cfg.Clean)
			}
		})
	}

	return shards
}

func (s CacheShard[E]) acquire(key string) *cache[E] {
	h := fnv.New32a()
	h.Write([]byte(key))
	return s[uint(h.Sum32())%uint(len(s))]
}

func (s CacheShard[E]) Get(key string) (E, error) {
	shard := s.acquire(key)

	shard.RLock()
	defer shard.RUnlock()

	val, ok := shard.items[key]
	if !ok {
		return val, errors.New("key not found")
	}
	if exp := shard.expires[key]; exp.Before(time.Now()) {
		delete(shard.items, key)
		delete(shard.expires, key)
		return val, errors.New("key is expired")
	}

	return val, nil
}

func (s CacheShard[E]) Set(key string, val E, exp time.Duration) error {
	shard := s.acquire(key)

	shard.Lock()
	defer shard.Unlock()

	shard.items[key] = val
	shard.expires[key] = time.Now().Add(exp)

	return nil
}

func (s CacheShard[E]) Delete(key string) error {
	shard := s.acquire(key)

	shard.Lock()
	defer shard.Unlock()

	delete(shard.items, key)
	delete(shard.expires, key)

	return nil
}

func (s CacheShard[E]) Reset() error {
	for _, shard := range s {
		shard.Lock()
		defer shard.Unlock()

		shard.items = make(map[string]E)
		shard.expires = make(map[string]time.Time)
	}
	return nil
}

func (s CacheShard[E]) Keys() []string {
	keys := []string{}
	for _, shard := range s {
		shard.Lock()
		defer shard.Unlock()

		for key := range shard.items {
			if exp := shard.expires[key]; exp.After(time.Now()) {
				keys = append(keys, key)
			}
		}
	}
	return keys
}

func (s CacheShard[E]) Values() []E {
	values := []E{}
	for _, shard := range s {
		shard.Lock()
		defer shard.Unlock()

		for key, value := range shard.items {
			if exp := shard.expires[key]; exp.After(time.Now()) {
				values = append(values, value)
			}
		}
	}
	return values
}
