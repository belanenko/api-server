package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig_Init(t *testing.T) {
	c, err := Init()
	assert.Nil(t, err)
	assert.NotEqual(t, 0, c.App.Port)
}
