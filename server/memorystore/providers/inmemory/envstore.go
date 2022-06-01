package inmemory

import (
	"os"
	"sync"
)

// EnvStore struct to store the env variables
type EnvStore struct {
	mutex sync.Mutex
	store map[string]interface{}
}

// UpdateEnvStore to update the whole env store object
func (e *EnvStore) UpdateStore(store map[string]interface{}) {
	if os.Getenv("ENV") != "test" {
		e.mutex.Lock()
		defer e.mutex.Unlock()
	}
	// just override the keys + new keys

	for key, value := range store {
		e.store[key] = value
	}
}

// GetStore returns the env store
func (e *EnvStore) GetStore() map[string]interface{} {
	if os.Getenv("ENV") != "test" {
		e.mutex.Lock()
		defer e.mutex.Unlock()
	}

	return e.store
}

// Get returns the value of the key in evn store
func (e *EnvStore) Get(key string) interface{} {
	if os.Getenv("ENV") != "test" {
		e.mutex.Lock()
		defer e.mutex.Unlock()
	}
	return e.store[key]
}

// Set sets the value of the key in env store
func (e *EnvStore) Set(key string, value interface{}) {
	if os.Getenv("ENV") != "test" {
		e.mutex.Lock()
		defer e.mutex.Unlock()
	}
	e.store[key] = value
}
