package factory

import "fmt"

type (
	CarFactory struct{}

	ICar interface {
		Run()
	}

	Benz struct{}
	Bmw  struct{}

	CarType int
)

const (
	CarTypeBenz CarType = iota + 1
	CarTypeBmw
)

func NewCarFactory() *CarFactory {
	return &CarFactory{}
}

func (b Benz) Run() {
	fmt.Println("Benz run")
}

func (b Bmw) Run() {
	fmt.Println("Bmw run")
}

func (CarFactory) Build(tp CarType) ICar {
	switch tp {
	case CarTypeBenz:
		return Benz{}
	case CarTypeBmw:
		return Bmw{}
	}

	return nil
}
