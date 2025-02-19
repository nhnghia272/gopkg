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

type CacheShardOption struct {
	Shard uint
	Clean time.Duration
}

func NewCacheShard[E any](opts ...*CacheShardOption) CacheShard[E] {
	opt := &CacheShardOption{Shard: 5, Clean: time.Minute}

	if len(opts) > 0 {
		opt = opts[0]
	}
	if opt.Shard == 0 {
		opt.Shard = 5
	}
	if opt.Clean <= 0 {
		opt.Clean = time.Minute
	}

	var shards = make(CacheShard[E], opt.Shard)

	for i := 0; i < int(opt.Shard); i++ {
		shards[i] = &cache[E]{
			items:   make(map[string]E),
			expires: make(map[string]time.Time),
		}
	}

	Async().Go(func() {
		for {
			shards.clean()
			time.Sleep(opt.Clean)
		}
	})

	return shards
}

func (s CacheShard[E]) clean() {
	for _, shard := range s {
		shard.Lock()
		defer shard.Unlock()

		for key, exp := range shard.expires {
			if !exp.IsZero() && exp.Before(time.Now()) {
				delete(shard.items, key)
				delete(shard.expires, key)
			}
		}
	}
}

func (s CacheShard[E]) acquire(key string) *cache[E] {
	h := fnv.New32a()
	h.Write([]byte(key))
	return s[int(h.Sum32())%len(s)]
}

func (s CacheShard[E]) Get(key string) (E, error) {
	shard := s.acquire(key)

	shard.RLock()
	defer shard.RUnlock()

	val, ok := shard.items[key]
	if !ok {
		return *new(E), errors.New("key not found")
	}
	if exp := shard.expires[key]; !exp.IsZero() && exp.Before(time.Now()) {
		return *new(E), errors.New("key is expired")
	}

	return val, nil
}

func (s CacheShard[E]) Set(key string, val E, exp time.Duration) {
	shard := s.acquire(key)

	shard.Lock()
	defer shard.Unlock()

	shard.items[key] = val

	if exp >= 0 {
		shard.expires[key] = time.Now().Add(exp)
	}
}

func (s CacheShard[E]) Delete(key string) {
	shard := s.acquire(key)

	shard.Lock()
	defer shard.Unlock()

	delete(shard.items, key)
	delete(shard.expires, key)
}

func (s CacheShard[E]) Reset() {
	for _, shard := range s {
		shard.Lock()
		defer shard.Unlock()

		shard.items = make(map[string]E)
		shard.expires = make(map[string]time.Time)
	}
}

func (s CacheShard[E]) Keys() []string {
	keys := make([]string, 0)
	for _, shard := range s {
		shard.Lock()
		defer shard.Unlock()

		for key := range shard.items {
			if exp := shard.expires[key]; exp.IsZero() || exp.After(time.Now()) {
				keys = append(keys, key)
			}
		}
	}
	return keys
}

func (s CacheShard[E]) Values() []E {
	values := make([]E, 0)
	for _, shard := range s {
		shard.Lock()
		defer shard.Unlock()

		for key, value := range shard.items {
			if exp := shard.expires[key]; exp.IsZero() || exp.After(time.Now()) {
				values = append(values, value)
			}
		}
	}
	return values
}
