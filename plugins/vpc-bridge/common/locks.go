package common

import (
	"sync"
)

// NamedLockManager manages a set of named locks.
type NamedLockManager struct {
	namedLocks map[string]*sync.RWMutex // Map of named Read-Write locks.
	mutex      sync.Mutex               // Locks the access to the map namedLocks.
}

// NewNamedLockManager creates a new NamedLockManager instance.
func NewNamedLockManager() *NamedLockManager {
	return &NamedLockManager{
		mutex:      sync.Mutex{},
		namedLocks: make(map[string]*sync.RWMutex),
	}
}

// Lock acquires the named lock.
func (m *NamedLockManager) Lock(name string) {
	// Lock the mutex for the map.
	m.mutex.Lock()

	namedLock, ok := m.namedLocks[name]
	if !ok {
		namedLock = &sync.RWMutex{}
		m.namedLocks[name] = namedLock
	}

	// Release the mutex on the map.
	m.mutex.Unlock()

	// Lock the namedLock.
	namedLock.Lock()
}

// Unlock releases the named lock.
func (m *NamedLockManager) Unlock(name string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	namedLock, ok := m.namedLocks[name]
	if !ok {
		// Lock not found, do nothing
		// print the log
		return
	}

	namedLock.Unlock()
}
