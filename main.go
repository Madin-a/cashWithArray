package main

import "fmt"

func main()  {
	purchases := [5]int{104, 102, 108, 105, 107}
	limitCashback :=30
	percent := 10
	sumPurchases := 0
	for _, value := range purchases {
		sumPurchases=sumPurchases+value
	}
	sumPersent := sumPurchases*percent/100
	if sumPersent<limitCashback {
		fmt.Println(sumPersent)
	}else{
		fmt.Println(limitCashback)
	}

}
