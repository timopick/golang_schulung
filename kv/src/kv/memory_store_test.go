package main

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func Test_InMemoryStore(t *testing.T) {
    s := NewInMemoryStore()
    s.Set("key1", "value1")
    s.Set("key2", "value2")
    s.Set("key3", "value3")

    assert.Equal(t, s.Get("key1"), "value1")
    assert.Equal(t, s.Get("key2"), "value2")
    assert.Equal(t, s.Get("key3"), "value3")
}
