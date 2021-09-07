package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommanderRun(t *testing.T) {
	commander := NewCommander()

	commander.dir = "/"

	err := commander.Run("fluff-commander-test")

	assert.Nil(t, err)
}
