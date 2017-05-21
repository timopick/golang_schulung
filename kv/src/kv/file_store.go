package main

import (
    "log"
    "os"
    "io/ioutil"
    "encoding/json"
    "time"
)

type FileStore struct {
    file *os.File
    store *InMemoryStore
    synced bool
}

func check(err error) {
    if err != nil {
        log.Fatal(err)
    }
}

func NewFileStore(path string) *FileStore {
    file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0755)
    check(err)

    var s FileStore
    s.file = file
    s.store = NewInMemoryStore()
    s.synced = true

    fileData, err := ioutil.ReadAll(s.file)
    check(err)

    if len(fileData) > 0 {
        err = json.Unmarshal(fileData, &s.store.internalStore)
        check(err)
    }

    return &s
}

func (f *FileStore) Get(key string) string {
    return f.store.Get(key)
}

func (f *FileStore) Set(key string, value string) {
    log.Printf("set %v, %v", key, value)
    f.store.Set(key, value)
    f.synced = false
}

func (f *FileStore) GetAll() map[string]string {
    return f.store.GetAll()
}

//TODO make thread safe
func (f *FileStore) Sync() {
    if f.synced {
        return
    }

    log.Println("Syncing store")
    startTime := time.Now()

    jsonData, err := json.Marshal(f.store.internalStore)
    check(err)

    n, err := f.file.WriteAt(jsonData, 0)
    check(err)
    err = f.file.Truncate(int64(len(jsonData)))
    check(err)

    f.synced = true
    log.Printf("Sync took %s, %d bytes written", time.Since(startTime), n)
}

func (f *FileStore) Close() {
    f.Sync()
    f.file.Close()
}
