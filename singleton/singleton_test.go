package singleton

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSingleton(t *testing.T) {
	name := "random name"
	s1 := NewSingleton()
	s2 := NewSingleton()
	s1.name = name
	assert.Equal(t, name, s2.name)
	assert.Equal(t, *s1, *s2)
}
