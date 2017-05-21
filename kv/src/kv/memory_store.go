package main

type InMemoryStore struct {
    internalStore map[string]string
}

func NewInMemoryStore() *InMemoryStore {
    var s InMemoryStore
    s.internalStore = make(map[string]string)
    return &s
}

func (s *InMemoryStore) Get(key string) string {
    return s.internalStore[key]
}

func (s *InMemoryStore) Set(key string, value string) {
    s.internalStore[key] = value
}

func (s *InMemoryStore) GetAll() map[string]string {
    return s.internalStore
}

func (s *InMemoryStore) Sync() {
    // nothing to do
}
