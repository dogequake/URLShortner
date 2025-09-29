package repository

import (
	"sync"
)

type URLRepository interface {
	Save(shortID, longURL string) error
	Find(shortID string) (string, bool)
}

type memoryRepo struct {
	data map[string]string
	mu   sync.RWMutex
}

func NewMemoryRepo() URLRepository {
	return &memoryRepo{
		data: make(map[string]string),
	}
}

func (r *memoryRepo) Save(shortID, longURL string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.data[shortID] = longURL
	return nil
}

func (r *memoryRepo) Find(shortID string) (string, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	url, ok := r.data[shortID]
	return url, ok
}
