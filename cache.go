package main

import (
	"sync"
	"time"

	"github.com/faroedev/faroe"
)

type cacheStruct struct {
	records map[string]cacheRecordStruct
	m       *sync.Mutex
}

func newCache() *cacheStruct {
	storage := &cacheStruct{
		records: map[string]cacheRecordStruct{},
		m:       &sync.Mutex{},
	}

	return storage
}

func (cache *cacheStruct) Get(key string) ([]byte, error) {
	cache.m.Lock()
	defer cache.m.Unlock()

	record, ok := cache.records[key]
	if !ok {
		return nil, faroe.ErrCacheEntryNotFound
	}

	if time.Now().Compare(record.expiresAt) >= 0 {
		delete(cache.records, key)

		return nil, faroe.ErrCacheEntryNotFound
	}

	return record.value, nil
}

func (cache *cacheStruct) Set(key string, value []byte, ttl time.Duration) error {
	cache.m.Lock()
	defer cache.m.Unlock()

	record := cacheRecordStruct{
		value:     value,
		expiresAt: time.Now().Add(ttl),
	}
	cache.records[key] = record

	return nil
}

func (cache *cacheStruct) Delete(key string) error {
	cache.m.Lock()
	defer cache.m.Unlock()

	delete(cache.records, key)

	return nil
}

type cacheRecordStruct struct {
	value     []byte
	expiresAt time.Time
}
