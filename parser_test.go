package squidparser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParse(t *testing.T) {
	actual := Parse("1480792667.325    388 10.2.37.199 TCP_MISS/200 804 POST http://clients1.google.com/ocsp juan.cabo HIER_DIRECT/172.217.2.142 application/ocsp-response").User
	expected := "juan.cabo"
	assert := assert.New(t)
	assert.Equal(actual, expected, "They should be equals")
}
