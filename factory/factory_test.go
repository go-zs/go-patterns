package factory

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewCarFactory(t *testing.T) {
	f := NewCarFactory()
	require.NotNil(t, f)
}

func TestCarFactory_Build(t *testing.T) {
	f := NewCarFactory()
	c1 := f.Build(CarTypeBenz)
	require.NotNil(t, c1)
	c2 := f.Build(CarTypeBmw)
	require.NotNil(t, c2)
	c1.Run()
	c2.Run()
}
