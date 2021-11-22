package mymutex

import "sync"

type Kit interface {
	Add(i float64)
	Has(i float64) bool
}

type MutexKit struct {
	sync.Mutex
	values map[float64]bool
}

type RWMutexKit struct {
	sync.RWMutex
	values map[float64]bool
}

func (s *MutexKit) Add(i float64) {
	s.Lock()
	defer s.Unlock()
	s.values[i] = true
}

func (s *MutexKit) Has(i float64) bool {
	s.Lock()
	defer s.Unlock()
	return s.values[i]
}

func NewMutexKit() Kit {
	return &MutexKit{values: make(map[float64]bool)}
}

func (s *RWMutexKit) Add(i float64) {
	s.Lock()
	defer s.Unlock()
	s.values[i] = true
}

func (s *RWMutexKit) Has(i float64) bool {
	s.RLock()
	defer s.RUnlock()
	return s.values[i]
}

func NewRWMutexKit() Kit {
	return &RWMutexKit{values: make(map[float64]bool)}
}
