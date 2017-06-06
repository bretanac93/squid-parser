package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var storage = Storage{m: make(map[string]uint64)}

func TestInsertOrUpdate(t *testing.T) {
	c, e := storage.InsertOrUpdate("cesar.bretana", 300)
	if assert.NoError(t, e) {
		assert.Equal(t, c, true, "They should be equals")
	}
}

func TestGetQuota(t *testing.T) {
	q, e := storage.GetQuota("cesar.bretana")
	if assert.NoError(t, e) {
		assert.Equal(t, q, uint64(300), "They should be equals")
	}
}
