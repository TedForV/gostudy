package prototype

import (
	"errors"
)

type ShirtCloner interface {
	GetClone(s int) (ItemInfoGetter, error)
}

const (
	White = 1
	Black = 2
	Blue  = 3
)

func GetShirtsCloner() ShirtCloner {
	return nil
}

type ShirtsCache struct {
}

func (sc *ShirtsCache) GetClone(s int) (ItemInfoGetter, error) {
	switch s {
	case White:
		newItem := *whitePrototype
		return &newItem, nil
	case Black:
		newItem := *blackProtorype
		return &newItem, nil
	case Blue:
		newItem := *bluePrototype
		return &newItem, nil
	default:
		return nil, errors.New("Not implemented yet")
	}

}

type ItemInfoGetter interface {
	GetInfo() string
}

type ShirtColor byte

type Shirt struct {
	Price float32
	SKU   string
	Color ShirtColor
}

func (s *Shirt) GetInfo() string {
	return ""
}

//func (s *Shirt) GetShirtsCloner() ShirtCloner {
//	return nil
//}

var whitePrototype *Shirt = &Shirt{
	15.00,
	"empty",
	White,
}
var blackProtorype *Shirt = &Shirt{
	16.00,
	"empty",
	Black,
}
var bluePrototype *Shirt = &Shirt{
	17.00,
	"empty",
	Blue,
}

func (i *Shirt) GetPrice() float32 {
	return i.Price
}
