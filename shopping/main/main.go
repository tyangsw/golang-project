package main

import (
	"fmt"
	"staging/shopping"
	"staging/shopping/models"
)

func main() {
	item := models.NewItem()
        item.Price = 23.65	
	fmt.Println(item.Price)
	fmt.Println(shopping.PriceCheck(1234))
}
