package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetGreeting(t *testing.T) {
	greeting := getGreeting()
	assert.Equal(t, greeting, "Hello, security minicamp!")
}
