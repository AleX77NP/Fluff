package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommanderCleanup(t *testing.T) {
	commander := NewCommander()

	commander.dir = "/"

	err := commander.Cleanup()

	assert.Nil(t, err)

}