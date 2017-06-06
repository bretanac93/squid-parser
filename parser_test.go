package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	var e error
	var pl LogEntry
	pl, e = Parse("1480792667.325    388 10.2.37.199 TCP_MISS/200 804 POST http://clients1.google.com/ocsp juan.cabo HIER_DIRECT/172.217.2.142 application/ocsp-response")
	if assert.NoError(t, e) {
		assert.Equal(t, pl.User, "juan.cabo",
			"They should be equals")
	}
}

func TestTransformFile(t *testing.T) {
	logLine := "1480792667.325    388 10.2.37.199 TCP_MISS/200 804 POST http://clients1.google.com/ocsp juan.cabo HIER_DIRECT/172.217.2.142 application/ocsp-response\n1480792667.325    388 10.2.37.199 TCP_MISS/200 804 POST http://clients1.google.com/ocsp juan.cabo HIER_DIRECT/172.217.2.142 application/ocsp-response"
	var e error
	var tf []LogEntry
	tf, e = TransformFile(strings.NewReader(logLine))
	if assert.NoError(t, e) {
		expected := "juan.cabo"
		assert.Equal(t, tf[0].User, expected, "They should be equals")
	}
}
