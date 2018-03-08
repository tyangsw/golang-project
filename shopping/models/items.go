package models

type Item struct {
	Price float64
}

func NewItem() *Item {
	return &Item{
		Price: 0.0,
	}
}
