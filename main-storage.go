package main

import (
	"sync"
	"time"

	"github.com/faroedev/faroe"
)

type mainStorageStruct struct {
	m       *sync.Mutex
	records map[string]mainStorageRecordStruct
}

func newMainStorage() *mainStorageStruct {
	storage := &mainStorageStruct{
		m:       &sync.Mutex{},
		records: map[string]mainStorageRecordStruct{},
	}
	return storage
}

func (mainStorage *mainStorageStruct) Get(key string) ([]byte, int32, error) {
	mainStorage.m.Lock()
	defer mainStorage.m.Unlock()

	record, ok := mainStorage.records[key]
	if !ok {
		return nil, 0, faroe.ErrMainStorageEntryNotFound
	}

	return record.value, record.counter, nil
}

func (mainStorage *mainStorageStruct) Set(key string, value []byte, _ time.Time) error {
	mainStorage.m.Lock()
	defer mainStorage.m.Unlock()

	mainStorage.records[key] = mainStorageRecordStruct{
		key:     key,
		value:   value,
		counter: 0,
	}

	return nil
}

func (mainStorage *mainStorageStruct) Update(key string, value []byte, _ time.Time, counter int32) error {
	mainStorage.m.Lock()
	defer mainStorage.m.Unlock()

	record, ok := mainStorage.records[key]
	if !ok {
		return faroe.ErrMainStorageEntryNotFound
	}
	if record.counter != counter {
		return faroe.ErrMainStorageEntryNotFound
	}
	record.value = value
	record.counter++
	mainStorage.records[key] = record

	return nil
}

func (mainStorage *mainStorageStruct) Delete(key string) error {
	mainStorage.m.Lock()
	defer mainStorage.m.Unlock()

	_, ok := mainStorage.records[key]
	if !ok {
		return faroe.ErrMainStorageEntryNotFound
	}
	delete(mainStorage.records, key)

	return nil
}

type mainStorageRecordStruct struct {
	key     string
	value   []byte
	counter int32
}
