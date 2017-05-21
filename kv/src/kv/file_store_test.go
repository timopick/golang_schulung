package main

import (
    "github.com/stretchr/testify/assert"
    "testing"
    "strconv"
)

func Test_FileStore(t *testing.T) {
    fs := NewFileStore("/tmp/KvTest")
    defer fs.Close()

    for i := 0; i < 100; i++ {
        idx := strconv.Itoa(i)
        fs.Set("key-" + idx, "value-" + idx)
    }

    assert.Equal(t, fs.Get("key-0"), "value-0")
    assert.Equal(t, fs.Get("key-99"), "value-99")
}
