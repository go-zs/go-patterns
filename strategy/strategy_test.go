package strategy

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewChrome(t *testing.T) {
	c := NewChrome()
	require.NotNil(t, c)
}

func TestChrome(t *testing.T) {
	c := NewChrome()
	c.Find()

	c.SetSearcher(Baidu{})
	c.Find()

	c.SetSearcher(Sougou{})
	c.Find()

	c.SetSearcher(Google{})
	c.Find()
}
