package strategy

import "fmt"

type (
	ISearcher interface {
		Search() string
	}

	Baidu struct{}

	Google struct{}

	Sougou struct{}

	Chrome struct {
		searcher ISearcher
	}

	ChromeOption func(chrome *Chrome)
)

func (Google) Search() string {
	return "google"
}

func (Baidu) Search() string {
	return "baidu"
}

func (Sougou) Search() string {
	return "sougou"
}

var (
	defaultChromeSearcher = Google{}
)

func NewChrome(options ...ChromeOption) *Chrome {
	c := &Chrome{
		searcher: defaultChromeSearcher,
	}
	for _, o := range options {
		o(c)
	}
	return c
}

func (c *Chrome) SetSearcher(s ISearcher) {
	c.searcher = s
}

func (c *Chrome) Find() {
	fmt.Println(c.searcher.Search())
}
