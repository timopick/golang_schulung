package main

type Store interface {
    Get(string) string
    Set(string,string)
    GetAll() map[string]string
    Sync()
}
