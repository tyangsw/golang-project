package shopping

import (
	"staging/shopping/db"
)

func PriceCheck(Id int) (float64, bool) {
	item := db.LoadItem(Id)

	if item == nil {
		return 0, false
	}

	return item.Price, true
}
