package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var storage = Storage{m: make(map[string]Stat)}

func TestInsertOrUpdate(t *testing.T) {
	c, e := storage.InsertOrUpdate("cesar.bretana", 300)
	if assert.NoError(t, e) {
		assert.Equal(t, c, true, "They should be equals")
	}
}

func TestGetUsage(t *testing.T) {
	q, e := storage.GetUsage("cesar.bretana")
	t.Log(q)
	if assert.NoError(t, e) {
		assert.Equal(t, q, uint64(300), "They should be equals")
	}
}
