package main

import (
	"sync"
	"time"

	"github.com/faroedev/faroe"
)

type rateLimitStorageStruct struct {
	records map[string]rateLimitStorageRecordStruct
	m       *sync.Mutex
}

func newRateLimitStorage() *rateLimitStorageStruct {
	records := map[string]rateLimitStorageRecordStruct{}
	storage := &rateLimitStorageStruct{
		records: records,
		m:       &sync.Mutex{},
	}

	return storage
}

func (rateLimitStorage *rateLimitStorageStruct) Get(key string) ([]byte, string, int32, error) {
	rateLimitStorage.m.Lock()
	defer rateLimitStorage.m.Unlock()

	record, ok := rateLimitStorage.records[key]
	if !ok {
		return nil, "", 0, faroe.ErrRateLimitStorageEntryNotFound
	}

	return record.value, record.id, record.counter, nil
}

func (rateLimitStorage *rateLimitStorageStruct) Add(key string, value []byte, entryId string, expiresAt time.Time) error {
	rateLimitStorage.m.Lock()
	defer rateLimitStorage.m.Unlock()

	_, ok := rateLimitStorage.records[key]
	if ok {
		return faroe.ErrRateLimitStorageEntryAlreadyExists
	}

	rateLimitStorage.records[key] = rateLimitStorageRecordStruct{
		value:   value,
		id:      entryId,
		counter: 0,
	}

	return nil
}

func (rateLimitStorage *rateLimitStorageStruct) Update(key string, value []byte, expiresAt time.Time, entryId string, counter int32) error {
	rateLimitStorage.m.Lock()
	defer rateLimitStorage.m.Unlock()

	record, ok := rateLimitStorage.records[key]
	if !ok {
		return faroe.ErrRateLimitStorageEntryNotFound
	}
	if record.id != entryId || record.counter != counter {
		return faroe.ErrRateLimitStorageEntryNotFound
	}

	record.value = value
	record.counter++

	rateLimitStorage.records[key] = record

	return nil
}

func (rateLimitStorage *rateLimitStorageStruct) Delete(key string, entryId string, counter int32) error {
	rateLimitStorage.m.Lock()
	defer rateLimitStorage.m.Unlock()

	record, ok := rateLimitStorage.records[key]
	if !ok {
		return faroe.ErrRateLimitStorageEntryNotFound
	}
	if record.id != entryId || record.counter != counter {
		return faroe.ErrRateLimitStorageEntryNotFound
	}

	delete(rateLimitStorage.records, key)

	return nil
}

type rateLimitStorageRecordStruct struct {
	value   []byte
	id      string
	counter int32
}
